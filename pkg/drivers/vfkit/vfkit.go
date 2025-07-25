//go:build darwin

/*
Copyright 2024 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vfkit

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/mcnutils"
	"github.com/docker/machine/libmachine/ssh"
	"github.com/docker/machine/libmachine/state"
	"github.com/pkg/errors"

	"k8s.io/klog/v2"
	pkgdrivers "k8s.io/minikube/pkg/drivers"
	"k8s.io/minikube/pkg/drivers/vmnet"
	"k8s.io/minikube/pkg/minikube/detect"
	"k8s.io/minikube/pkg/minikube/exit"
	"k8s.io/minikube/pkg/minikube/firewall"
	"k8s.io/minikube/pkg/minikube/out"
	"k8s.io/minikube/pkg/minikube/process"
	"k8s.io/minikube/pkg/minikube/reason"
	"k8s.io/minikube/pkg/minikube/style"
)

const (
	isoFilename    = "boot2docker.iso"
	pidFileName    = "vfkit.pid"
	sockFilename   = "vfkit.sock"
	logFileName    = "vfkit.log"
	serialFileName = "serial.log"
	defaultSSHUser = "docker"
)

// Driver is the machine driver for vfkit (Virtualization.framework)
type Driver struct {
	*drivers.BaseDriver
	*pkgdrivers.CommonDriver
	Boot2DockerURL string
	DiskSize       int
	CPU            int
	Memory         int
	ExtraDisks     int
	Network        string        // "", "nat", "vmnet-shared"
	MACAddress     string        // For network=nat, network=""
	VmnetHelper    *vmnet.Helper // For network=vmnet-shared
}

func NewDriver(hostName, storePath string) drivers.Driver {
	return &Driver{
		BaseDriver: &drivers.BaseDriver{
			SSHUser:     defaultSSHUser,
			MachineName: hostName,
			StorePath:   storePath,
		},
		CommonDriver: &pkgdrivers.CommonDriver{},
	}
}

func (d *Driver) PreCreateCheck() error {
	return nil
}

func (d *Driver) GetMachineName() string {
	return d.MachineName
}

func (d *Driver) DriverName() string {
	return "vfkit"
}

func (d *Driver) GetSSHHostname() (string, error) {
	return d.IPAddress, nil
}

func (d *Driver) GetSSHKeyPath() string {
	return d.ResolveStorePath("id_rsa")
}

func (d *Driver) GetSSHPort() (int, error) {
	if d.SSHPort == 0 {
		d.SSHPort = 22
	}
	return d.SSHPort, nil
}

func (d *Driver) GetSSHUsername() string {
	if d.SSHUser == "" {
		d.SSHUser = defaultSSHUser
	}

	return d.SSHUser
}

func (d *Driver) GetURL() (string, error) {
	if _, err := os.Stat(d.pidfilePath()); err != nil {
		return "", nil
	}
	ip, err := d.GetIP()
	if err != nil {
		log.Warnf("Failed to get IP: %v", err)
		return "", err
	}
	if ip == "" {
		return "", nil
	}
	return fmt.Sprintf("tcp://%s:2376", ip), nil
}

func (d *Driver) GetIP() (string, error) {
	return d.IPAddress, nil
}

func (d *Driver) getVfkitState() (state.State, error) {
	pidfile := d.pidfilePath()
	pid, err := process.ReadPidfile(pidfile)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return state.Error, err
		}
		return state.Stopped, nil
	}
	exists, err := process.Exists(pid, "vfkit")
	if err != nil {
		return state.Error, err
	}
	if !exists {
		// No process, stale pidfile.
		if err := os.Remove(pidfile); err != nil {
			log.Debugf("failed to remove %q: %s", pidfile, err)
		}
		return state.Stopped, nil
	}
	return state.Running, nil
}

func (d *Driver) getVmnetHelperState() (state.State, error) {
	if d.VmnetHelper == nil {
		return state.Stopped, nil
	}
	return d.VmnetHelper.GetState()
}

// GetState returns driver state. Since vfkit driver may use 2 processes
// (vmnet-helper, vfkit), this returns combined state of both processes.
func (d *Driver) GetState() (state.State, error) {
	if vfkitState, err := d.getVfkitState(); err != nil {
		return state.Error, err
	} else if vfkitState == state.Running {
		return state.Running, nil
	}
	return d.getVmnetHelperState()
}

func (d *Driver) Create() error {
	var err error
	if d.SSHPort, err = d.GetSSHPort(); err != nil {
		return err
	}
	b2dutils := mcnutils.NewB2dUtils(d.StorePath)
	if err := b2dutils.CopyIsoToMachineDir(d.Boot2DockerURL, d.MachineName); err != nil {
		return err
	}

	if err := d.extractKernel(); err != nil {
		return err
	}

	log.Info("Creating SSH key...")
	if err := ssh.GenerateSSHKey(d.sshKeyPath()); err != nil {
		return err
	}

	log.Info("Creating Disk image...")
	if err := d.generateDiskImage(d.DiskSize); err != nil {
		return err
	}

	if d.ExtraDisks > 0 {
		log.Info("Creating extra disk images...")
		for i := 0; i < d.ExtraDisks; i++ {
			path := pkgdrivers.ExtraDiskPath(d.BaseDriver, i)
			if err := pkgdrivers.CreateRawDisk(path, d.DiskSize); err != nil {
				return err
			}
		}
	}

	log.Info("Starting vfkit VM...")
	return d.Start()
}

func (d *Driver) extractKernel() error {
	log.Info("Extracting bzimage and initrd...")
	isoPath := d.ResolveStorePath(isoFilename)
	if err := pkgdrivers.ExtractFile(isoPath, "/boot/bzimage", d.kernelPath()); err != nil {
		return err
	}
	return pkgdrivers.ExtractFile(isoPath, "/boot/initrd", d.initrdPath())
}

func (d *Driver) Start() error {
	var socketPath string

	if d.VmnetHelper != nil {
		socketPath = d.VmnetHelper.SocketPath()
		if err := d.VmnetHelper.Start(socketPath); err != nil {
			return err
		}

		d.MACAddress = d.VmnetHelper.GetMACAddress()
	}

	if err := d.startVfkit(socketPath); err != nil {
		return err
	}

	if err := d.setupIP(d.MACAddress); err != nil {
		return err
	}

	log.Infof("Waiting for VM to start (ssh -p %d docker@%s)...", d.SSHPort, d.IPAddress)

	return WaitForTCPWithDelay(fmt.Sprintf("%s:%d", d.IPAddress, d.SSHPort), time.Second)
}

// startVfkit starts the vfkit child process. If socketPath is not empty, vfkit
// is connected to the vmnet network via the socket instead of "nat" network.
func (d *Driver) startVfkit(socketPath string) error {
	machineDir := filepath.Join(d.StorePath, "machines", d.GetMachineName())

	var startCmd []string

	startCmd = append(startCmd,
		"--memory", fmt.Sprintf("%d", d.Memory),
		"--cpus", fmt.Sprintf("%d", d.CPU),
		"--restful-uri", fmt.Sprintf("unix://%s", d.sockfilePath()),
		"--log-level", "debug")

	// On arm64 console= is required get boot messages in serial.log. On x86_64
	// serial log is always empty.
	var cmdline string
	switch runtime.GOARCH {
	case "arm64":
		cmdline = "console=hvc0"
	case "amd64":
		cmdline = "console=ttyS0"
	}

	// TODO: Switch to --bootloader efi when x86_64 iso changed to EFI.
	startCmd = append(startCmd,
		"--bootloader", fmt.Sprintf("linux,kernel=%s,initrd=%s,cmdline=\"%s\"",
			d.kernelPath(), d.initrdPath(), cmdline))

	if socketPath != "" {
		// The guest will be able to access other guests in the vmnet network.
		startCmd = append(startCmd,
			"--device", fmt.Sprintf("virtio-net,unixSocketPath=%s,mac=%s", socketPath, d.MACAddress))
	} else {
		// The guest will not be able to access other guests.
		startCmd = append(startCmd,
			"--device", fmt.Sprintf("virtio-net,nat,mac=%s", d.MACAddress))
	}

	startCmd = append(startCmd,
		"--device", "virtio-rng")

	var isoPath = filepath.Join(machineDir, isoFilename)
	startCmd = append(startCmd,
		"--device", fmt.Sprintf("virtio-blk,path=%s", isoPath))

	startCmd = append(startCmd,
		"--device", fmt.Sprintf("virtio-blk,path=%s", d.diskPath()))

	for i := 0; i < d.ExtraDisks; i++ {
		startCmd = append(startCmd,
			"--device", fmt.Sprintf("virtio-blk,path=%s", pkgdrivers.ExtraDiskPath(d.BaseDriver, i)))
	}

	serialPath := d.ResolveStorePath(serialFileName)
	startCmd = append(startCmd,
		"--device", fmt.Sprintf("virtio-serial,logFilePath=%s", serialPath))

	log.Debugf("executing: vfkit %s", strings.Join(startCmd, " "))
	os.Remove(d.sockfilePath())
	cmd := exec.Command("vfkit", startCmd...)

	// Create vfkit in a new process group, so minikube caller can use killpg
	// to terminate the entire process group without harming the vfkit process.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	logfile, err := d.openLogfile()
	if err != nil {
		return fmt.Errorf("failed to open vfkit logfile: %w", err)
	}
	defer logfile.Close()
	cmd.Stderr = logfile

	if err := cmd.Start(); err != nil {
		return err
	}
	return process.WritePidfile(d.pidfilePath(), cmd.Process.Pid)
}

func (d *Driver) setupIP(mac string) error {
	var err error
	getIP := func() error {
		d.IPAddress, err = pkgdrivers.GetIPAddressByMACAddress(mac)
		if err != nil {
			return errors.Wrap(err, "failed to get IP address")
		}
		return nil
	}
	// Implement a retry loop because IP address isn't added to dhcp leases file immediately
	multiplier := 1
	if detect.NestedVM() {
		multiplier = 3 // will help with running in Free github action Macos VMs (takes 160+ retries on average)
	}
	for i := 0; i < 60*multiplier; i++ {
		log.Debugf("Attempt %d", i)
		err = getIP()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err == nil {
		log.Debugf("IP: %s", d.IPAddress)
		return nil
	}
	if !isBootpdError(err) {
		return errors.Wrap(err, "IP address never found in dhcp leases file")
	}
	if unblockErr := firewall.UnblockBootpd(); unblockErr != nil {
		klog.Errorf("failed unblocking bootpd from firewall: %v", unblockErr)
		exit.Error(reason.IfBootpdFirewall, "ip not found", err)
	}
	out.Styled(style.Restarting, "Successfully unblocked bootpd process from firewall, retrying")
	return fmt.Errorf("ip not found: %v", err)
}

func isBootpdError(err error) bool {
	return strings.Contains(err.Error(), "could not find an IP address")
}

func (d *Driver) openLogfile() (*os.File, error) {
	logfile := d.ResolveStorePath(logFileName)
	return os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o600)
}

func (d *Driver) stopVfkit() error {
	if err := d.SetVFKitState("Stop"); err != nil {
		// vfkit may be already stopped, shutting down, or not listening.
		// We don't fallback to "HardStop" since it typically fails due to
		// https://github.com/crc-org/vfkit/issues/277.
		log.Debugf("Failed to set vfkit state to 'Stop': %s", err)
		pidfile := d.pidfilePath()
		pid, err := process.ReadPidfile(pidfile)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
			// No pidfile.
			return nil
		}
		if err := process.Terminate(pid, "vfkit"); err != nil {
			if err != os.ErrProcessDone {
				return err
			}
			// No process, stale pidfile.
			if err := os.Remove(pidfile); err != nil {
				log.Debugf("failed to remove %q: %s", pidfile, err)
			}
			return nil
		}
	}
	return nil
}

func (d *Driver) stopVmnetHelper() error {
	if d.VmnetHelper == nil {
		return nil
	}
	return d.VmnetHelper.Stop()
}

func (d *Driver) Stop() error {
	if err := d.stopVfkit(); err != nil {
		return err
	}
	return d.stopVmnetHelper()
}

func (d *Driver) Remove() error {
	s, err := d.GetState()
	if err != nil {
		return errors.Wrap(err, "get state")
	}
	if s == state.Running {
		if err := d.Kill(); err != nil {
			return errors.Wrap(err, "kill")
		}
	}
	return nil
}

func (d *Driver) Restart() error {
	s, err := d.GetState()
	if err != nil {
		return err
	}

	if s == state.Running {
		if err := d.Stop(); err != nil {
			return err
		}
	}
	return d.Start()
}

func (d *Driver) killVfkit() error {
	if err := d.SetVFKitState("HardStop"); err != nil {
		// Typically fails with EOF due to https://github.com/crc-org/vfkit/issues/277.
		log.Debugf("Failed to set vfkit state to 'HardStop': %s", err)
		pidfile := d.pidfilePath()
		pid, err := process.ReadPidfile(pidfile)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
			// No pidfile.
			return nil
		}
		if err := process.Kill(pid, "vfkit"); err != nil {
			if err != os.ErrProcessDone {
				return err
			}
			// No process, stale pidfile.
			if err := os.Remove(pidfile); err != nil {
				log.Debugf("failed to remove %q: %s", pidfile, err)
			}
			return nil
		}
	}
	return nil
}

func (d *Driver) killVmnetHelper() error {
	if d.VmnetHelper == nil {
		return nil
	}
	return d.VmnetHelper.Kill()
}

func (d *Driver) Kill() error {
	if err := d.killVfkit(); err != nil {
		return err
	}
	return d.killVmnetHelper()
}

func (d *Driver) StartDocker() error {
	return fmt.Errorf("hosts without a driver cannot start docker")
}

func (d *Driver) StopDocker() error {
	return fmt.Errorf("hosts without a driver cannot stop docker")
}

func (d *Driver) GetDockerConfigDir() string {
	return ""
}

func (d *Driver) Upgrade() error {
	return fmt.Errorf("hosts without a driver cannot be upgraded")
}

func (d *Driver) sshKeyPath() string {
	machineDir := filepath.Join(d.StorePath, "machines", d.GetMachineName())
	return filepath.Join(machineDir, "id_rsa")
}

func (d *Driver) publicSSHKeyPath() string {
	return d.sshKeyPath() + ".pub"
}

func (d *Driver) kernelPath() string {
	return d.ResolveStorePath("bzimage")
}

func (d *Driver) initrdPath() string {
	return d.ResolveStorePath("initrd")
}

func (d *Driver) diskPath() string {
	machineDir := filepath.Join(d.StorePath, "machines", d.GetMachineName())
	return filepath.Join(machineDir, "disk.img")
}

func (d *Driver) sockfilePath() string {
	machineDir := filepath.Join(d.StorePath, "machines", d.GetMachineName())
	return filepath.Join(machineDir, sockFilename)
}

func (d *Driver) pidfilePath() string {
	machineDir := filepath.Join(d.StorePath, "machines", d.GetMachineName())
	return filepath.Join(machineDir, pidFileName)
}

// Make a boot2docker VM disk image.
func (d *Driver) generateDiskImage(size int) error {
	log.Debugf("Creating %d MB hard disk image...", size)

	magicString := "boot2docker, please format-me"

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	// magicString first so the automount script knows to format the disk
	file := &tar.Header{Name: magicString, Size: int64(len(magicString))}
	if err := tw.WriteHeader(file); err != nil {
		return err
	}
	if _, err := tw.Write([]byte(magicString)); err != nil {
		return err
	}
	// .ssh/key.pub => authorized_keys
	file = &tar.Header{Name: ".ssh", Typeflag: tar.TypeDir, Mode: 0700}
	if err := tw.WriteHeader(file); err != nil {
		return err
	}
	pubKey, err := os.ReadFile(d.publicSSHKeyPath())
	if err != nil {
		return err
	}
	file = &tar.Header{Name: ".ssh/authorized_keys", Size: int64(len(pubKey)), Mode: 0644}
	if err := tw.WriteHeader(file); err != nil {
		return err
	}
	if _, err := tw.Write(pubKey); err != nil {
		return err
	}
	file = &tar.Header{Name: ".ssh/authorized_keys2", Size: int64(len(pubKey)), Mode: 0644}
	if err := tw.WriteHeader(file); err != nil {
		return err
	}
	if _, err := tw.Write(pubKey); err != nil {
		return err
	}
	if err := tw.Close(); err != nil {
		return err
	}
	rawFile := d.diskPath()
	if err := os.WriteFile(rawFile, buf.Bytes(), 0644); err != nil {
		return nil
	}
	if err := os.Truncate(rawFile, int64(size)*int64(1024*1024)); err != nil {
		return nil
	}
	log.Debugf("DONE writing to %s and %s", rawFile, d.diskPath())
	return nil
}

func httpUnixClient(path string) http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", path)
			},
		},
	}
}

type VMState struct {
	State string `json:"state"`
}

func (d *Driver) GetVFKitState() (string, error) {
	httpc := httpUnixClient(d.sockfilePath())
	var vmstate VMState
	response, err := httpc.Get("http://_/vm/state")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&vmstate)
	if err != nil {
		return "", err
	}
	log.Debugf("get state: %+v", vmstate)
	return vmstate.State, nil
}

// SetVFKitState sets the state of the vfkit VM, (s is the state)
func (d *Driver) SetVFKitState(s string) error {
	httpc := httpUnixClient(d.sockfilePath())
	var vmstate VMState
	vmstate.State = s
	data, err := json.Marshal(&vmstate)
	if err != nil {
		return err
	}
	_, err = httpc.Post("http://_/vm/state", "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	log.Infof("Set vfkit state: %+v", vmstate)
	return nil
}

func WaitForTCPWithDelay(addr string, duration time.Duration) error {
	for {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			continue
		}
		defer conn.Close()
		if _, err := conn.Read(make([]byte, 1)); err != nil && err != io.EOF {
			time.Sleep(duration)
			continue
		}
		break
	}
	return nil
}
