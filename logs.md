=== RUN   TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads
no_kubernetes_test.go:91: Checking cache directory: /home/jenkins/minikube-integration/21139-381342/.minikube/cache/linux/amd64/v0.0.0
no_kubernetes_test.go:100: Cache directory exists but is empty
no_kubernetes_test.go:102: Cache directory /home/jenkins/minikube-integration/21139-381342/.minikube/cache/linux/amd64/v0.0.0 should not exist when using --no-kubernetes
helpers_test.go:222: -----------------------post-mortem--------------------------------
helpers_test.go:223: ======>  post-mortem[TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads]: network settings <======
helpers_test.go:230: HOST ENV snapshots: PROXY env: HTTP_PROXY="<empty>" HTTPS_PROXY="<empty>" NO_PROXY="<empty>"
helpers_test.go:238: ======>  post-mortem[TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads]: docker inspect <======
helpers_test.go:239: (dbg) Run:  docker inspect NoKubernetes-115222
helpers_test.go:243: (dbg) docker inspect NoKubernetes-115222:
-- stdout --
	[
	    {
	        "Id": "b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17",
	        "Created": "2025-10-02T12:43:58.931222211Z",
	        "Path": "/usr/local/bin/entrypoint",
	        "Args": [
	            "/sbin/init"
	        ],
	        "State": {
	            "Status": "running",
	            "Running": true,
	            "Paused": false,
	            "Restarting": false,
	            "OOMKilled": false,
	            "Dead": false,
	            "Pid": 586336,
	            "ExitCode": 0,
	            "Error": "",
	            "StartedAt": "2025-10-02T12:43:58.960923912Z",
	            "FinishedAt": "0001-01-01T00:00:00Z"
	        },
	        "Image": "sha256:c6b5532e987b5b4f5fc9cb0336e378ed49c0542bad8cbfc564b71e977a6269de",
	        "ResolvConfPath": "/var/lib/docker/containers/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17/resolv.conf",
	        "HostnamePath": "/var/lib/docker/containers/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17/hostname",
	        "HostsPath": "/var/lib/docker/containers/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17/hosts",
	        "LogPath": "/var/lib/docker/containers/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17-json.log",
	        "Name": "/NoKubernetes-115222",
	        "RestartCount": 0,
	        "Driver": "overlay2",
	        "Platform": "linux",
	        "MountLabel": "",
	        "ProcessLabel": "",
	        "AppArmorProfile": "unconfined",
	        "ExecIDs": null,
	        "HostConfig": {
	            "Binds": [
	                "/lib/modules:/lib/modules:ro",
	                "NoKubernetes-115222:/var"
	            ],
	            "ContainerIDFile": "",
	            "LogConfig": {
	                "Type": "json-file",
	                "Config": {
	                    "max-size": "100m"
	                }
	            },
	            "NetworkMode": "NoKubernetes-115222",
	            "PortBindings": {
	                "22/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ],
	                "2376/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ],
	                "32443/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ],
	                "5000/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ],
	                "8443/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ]
	            },
	            "RestartPolicy": {
	                "Name": "no",
	                "MaximumRetryCount": 0
	            },
	            "AutoRemove": false,
	            "VolumeDriver": "",
	            "VolumesFrom": null,
	            "ConsoleSize": [
	                0,
	                0
	            ],
	            "CapAdd": null,
	            "CapDrop": null,
	            "CgroupnsMode": "private",
	            "Dns": [],
	            "DnsOptions": [],
	            "DnsSearch": [],
	            "ExtraHosts": null,
	            "GroupAdd": null,
	            "IpcMode": "private",
	            "Cgroup": "",
	            "Links": null,
	            "OomScoreAdj": 0,
	            "PidMode": "",
	            "Privileged": true,
	            "PublishAllPorts": false,
	            "ReadonlyRootfs": false,
	            "SecurityOpt": [
	                "seccomp=unconfined",
	                "apparmor=unconfined",
	                "label=disable"
	            ],
	            "Tmpfs": {
	                "/run": "",
	                "/tmp": ""
	            },
	            "UTSMode": "",
	            "UsernsMode": "",
	            "ShmSize": 67108864,
	            "Runtime": "runc",
	            "Isolation": "",
	            "CpuShares": 0,
	            "Memory": 3221225472,
	            "NanoCpus": 0,
	            "CgroupParent": "",
	            "BlkioWeight": 0,
	            "BlkioWeightDevice": [],
	            "BlkioDeviceReadBps": [],
	            "BlkioDeviceWriteBps": [],
	            "BlkioDeviceReadIOps": [],
	            "BlkioDeviceWriteIOps": [],
	            "CpuPeriod": 0,
	            "CpuQuota": 0,
	            "CpuRealtimePeriod": 0,
	            "CpuRealtimeRuntime": 0,
	            "CpusetCpus": "",
	            "CpusetMems": "",
	            "Devices": [],
	            "DeviceCgroupRules": null,
	            "DeviceRequests": null,
	            "MemoryReservation": 0,
	            "MemorySwap": 6442450944,
	            "MemorySwappiness": null,
	            "OomKillDisable": null,
	            "PidsLimit": null,
	            "Ulimits": [],
	            "CpuCount": 0,
	            "CpuPercent": 0,
	            "IOMaximumIOps": 0,
	            "IOMaximumBandwidth": 0,
	            "MaskedPaths": null,
	            "ReadonlyPaths": null
	        },
	        "GraphDriver": {
	            "Data": {
	                "ID": "b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17",
	                "LowerDir": "/var/lib/docker/overlay2/7caa147e4324acc509479494c329388ece54f4832250947aeb81d55f419ff34b-init/diff:/var/lib/docker/overlay2/74e23e2d414e71bc60aae2442f772d94c45dcc7da38ffe98fa74cb259c3e7865/diff",
	                "MergedDir": "/var/lib/docker/overlay2/7caa147e4324acc509479494c329388ece54f4832250947aeb81d55f419ff34b/merged",
	                "UpperDir": "/var/lib/docker/overlay2/7caa147e4324acc509479494c329388ece54f4832250947aeb81d55f419ff34b/diff",
	                "WorkDir": "/var/lib/docker/overlay2/7caa147e4324acc509479494c329388ece54f4832250947aeb81d55f419ff34b/work"
	            },
	            "Name": "overlay2"
	        },
	        "Mounts": [
	            {
	                "Type": "volume",
	                "Name": "NoKubernetes-115222",
	                "Source": "/var/lib/docker/volumes/NoKubernetes-115222/_data",
	                "Destination": "/var",
	                "Driver": "local",
	                "Mode": "z",
	                "RW": true,
	                "Propagation": ""
	            },
	            {
	                "Type": "bind",
	                "Source": "/lib/modules",
	                "Destination": "/lib/modules",
	                "Mode": "ro",
	                "RW": false,
	                "Propagation": "rprivate"
	            }
	        ],
	        "Config": {
	            "Hostname": "NoKubernetes-115222",
	            "Domainname": "",
	            "User": "",
	            "AttachStdin": false,
	            "AttachStdout": false,
	            "AttachStderr": false,
	            "ExposedPorts": {
	                "22/tcp": {},
	                "2376/tcp": {},
	                "32443/tcp": {},
	                "5000/tcp": {},
	                "8443/tcp": {}
	            },
	            "Tty": true,
	            "OpenStdin": false,
	            "StdinOnce": false,
	            "Env": [
	                "container=docker",
	                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
	            ],
	            "Cmd": null,
	            "Image": "gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1",
	            "Volumes": null,
	            "WorkingDir": "/",
	            "Entrypoint": [
	                "/usr/local/bin/entrypoint",
	                "/sbin/init"
	            ],
	            "OnBuild": null,
	            "Labels": {
	                "created_by.minikube.sigs.k8s.io": "true",
	                "mode.minikube.sigs.k8s.io": "NoKubernetes-115222",
	                "name.minikube.sigs.k8s.io": "NoKubernetes-115222",
	                "role.minikube.sigs.k8s.io": ""
	            },
	            "StopSignal": "SIGRTMIN+3"
	        },
	        "NetworkSettings": {
	            "Bridge": "",
	            "SandboxID": "684712f731d0b43f5cd680d6bec133d72160dd1584ed185f8bc18deca1caec9b",
	            "SandboxKey": "/var/run/docker/netns/684712f731d0",
	            "Ports": {
	                "22/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33390"
	                    }
	                ],
	                "2376/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33391"
	                    }
	                ],
	                "32443/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33394"
	                    }
	                ],
	                "5000/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33392"
	                    }
	                ],
	                "8443/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33393"
	                    }
	                ]
	            },
	            "HairpinMode": false,
	            "LinkLocalIPv6Address": "",
	            "LinkLocalIPv6PrefixLen": 0,
	            "SecondaryIPAddresses": null,
	            "SecondaryIPv6Addresses": null,
	            "EndpointID": "",
	            "Gateway": "",
	            "GlobalIPv6Address": "",
	            "GlobalIPv6PrefixLen": 0,
	            "IPAddress": "",
	            "IPPrefixLen": 0,
	            "IPv6Gateway": "",
	            "MacAddress": "",
	            "Networks": {
	                "NoKubernetes-115222": {
	                    "IPAMConfig": {
	                        "IPv4Address": "192.168.85.2"
	                    },
	                    "Links": null,
	                    "Aliases": null,
	                    "MacAddress": "06:12:47:d1:15:94",
	                    "DriverOpts": null,
	                    "GwPriority": 0,
	                    "NetworkID": "b83a7484a0439ab1227e41307c48bfff54849f6fa122889143471cb7e6e2a017",
	                    "EndpointID": "ca65baae7c15a7c93d2ff87d7e4421a583cf63977e817f6bde2e82559b78c05f",
	                    "Gateway": "192.168.85.1",
	                    "IPAddress": "192.168.85.2",
	                    "IPPrefixLen": 24,
	                    "IPv6Gateway": "",
	                    "GlobalIPv6Address": "",
	                    "GlobalIPv6PrefixLen": 0,
	                    "DNSNames": [
	                        "NoKubernetes-115222",
	                        "b4bf0efb9b94"
	                    ]
	                }
	            }
	        }
	    }
	]
-- /stdout --
helpers_test.go:247: (dbg) Run:  out/minikube-linux-amd64 status --format={{.Host}} -p NoKubernetes-115222 -n NoKubernetes-115222
helpers_test.go:247: (dbg) Non-zero exit: out/minikube-linux-amd64 status --format={{.Host}} -p NoKubernetes-115222 -n NoKubernetes-115222: exit status 6 (344.060193ms)
-- stdout --
	Running
	WARNING: Your kubectl is pointing to stale minikube-vm.
	To fix the kubectl context, run `minikube update-context`
-- /stdout --
** stderr ** 
	E1002 12:44:02.192955  588550 status.go:458] kubeconfig endpoint: get endpoint: "NoKubernetes-115222" does not appear in /home/jenkins/minikube-integration/21139-381342/kubeconfig
** /stderr **
helpers_test.go:247: status error: exit status 6 (may be ok)
helpers_test.go:252: <<< TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads FAILED: start of post-mortem logs <<<
helpers_test.go:253: ======>  post-mortem[TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads]: minikube logs <======
helpers_test.go:255: (dbg) Run:  out/minikube-linux-amd64 -p NoKubernetes-115222 logs -n 25
helpers_test.go:260: TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads logs: 
-- stdout --
	
	==> Audit <==
	┌─────────┬───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┬─────────────────────────────┬─────────┬─────────┬─────────────────────┬─────────────────────┐
	│ COMMAND │                                                             ARGS                                                              │           PROFILE           │  USER   │ VERSION │     START TIME      │      END TIME       │
	├─────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼─────────────────────────────┼─────────┼─────────┼─────────────────────┼─────────────────────┤
	│ stop    │ -p scheduled-stop-180674 --schedule 5m                                                                                        │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 5m                                                                                        │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --cancel-scheduled                                                                                   │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │ 02 Oct 25 12:41 UTC │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │ 02 Oct 25 12:42 UTC │
	│ delete  │ -p scheduled-stop-180674                                                                                                      │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:42 UTC │
	│ start   │ -p insufficient-storage-013758 --memory=3072 --output=json --wait=true --driver=docker  --container-runtime=containerd        │ insufficient-storage-013758 │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │                     │
	│ delete  │ -p insufficient-storage-013758                                                                                                │ insufficient-storage-013758 │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:42 UTC │
	│ start   │ -p offline-containerd-106797 --alsologtostderr -v=1 --memory=3072 --wait=true --driver=docker  --container-runtime=containerd │ offline-containerd-106797   │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p NoKubernetes-115222 --no-kubernetes --kubernetes-version=v1.28.0 --driver=docker  --container-runtime=containerd           │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │                     │
	│ start   │ -p NoKubernetes-115222 --memory=3072 --alsologtostderr -v=5 --driver=docker  --container-runtime=containerd                   │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p running-upgrade-583950 --memory=3072 --vm-driver=docker  --container-runtime=containerd                                    │ running-upgrade-583950      │ jenkins │ v1.32.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p stopped-upgrade-250350 --memory=3072 --vm-driver=docker  --container-runtime=containerd                                    │ stopped-upgrade-250350      │ jenkins │ v1.32.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p NoKubernetes-115222 --no-kubernetes --memory=3072 --alsologtostderr -v=5 --driver=docker  --container-runtime=containerd   │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:43 UTC │
	│ delete  │ -p offline-containerd-106797                                                                                                  │ offline-containerd-106797   │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p running-upgrade-583950 --memory=3072 --alsologtostderr -v=1 --driver=docker  --container-runtime=containerd                │ running-upgrade-583950      │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │                     │
	│ stop    │ stopped-upgrade-250350 stop                                                                                                   │ stopped-upgrade-250350      │ jenkins │ v1.32.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p missing-upgrade-404552 --memory=3072 --driver=docker  --container-runtime=containerd                                       │ missing-upgrade-404552      │ jenkins │ v1.32.0 │ 02 Oct 25 12:43 UTC │                     │
	│ start   │ -p stopped-upgrade-250350 --memory=3072 --alsologtostderr -v=1 --driver=docker  --container-runtime=containerd                │ stopped-upgrade-250350      │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │                     │
	│ delete  │ -p NoKubernetes-115222                                                                                                        │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p NoKubernetes-115222 --no-kubernetes --memory=3072 --alsologtostderr -v=5 --driver=docker  --container-runtime=containerd   │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:44 UTC │
	└─────────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┴─────────────────────────────┴─────────┴─────────┴─────────────────────┴─────────────────────┘
	
	
	==> Last Start <==
	Log file created at: 2025/10/02 12:43:58
	Running on machine: ubuntu-20-agent-6
	Binary: Built with gc go1.24.6 for linux/amd64
	Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg
	I1002 12:43:58.143632  585741 out.go:360] Setting OutFile to fd 1 ...
	I1002 12:43:58.143942  585741 out.go:408] TERM=,COLORTERM=, which probably does not support color
	I1002 12:43:58.143954  585741 out.go:374] Setting ErrFile to fd 2...
	I1002 12:43:58.143958  585741 out.go:408] TERM=,COLORTERM=, which probably does not support color
	I1002 12:43:58.144172  585741 root.go:338] Updating PATH: /home/jenkins/minikube-integration/21139-381342/.minikube/bin
	I1002 12:43:58.144654  585741 out.go:368] Setting JSON to false
	I1002 12:43:58.146105  585741 start.go:130] hostinfo: {"hostname":"ubuntu-20-agent-6","uptime":8776,"bootTime":1759400262,"procs":281,"os":"linux","platform":"ubuntu","platformFamily":"debian","platformVersion":"22.04","kernelVersion":"6.8.0-1041-gcp","kernelArch":"x86_64","virtualizationSystem":"kvm","virtualizationRole":"guest","hostId":"591c9f12-2938-3743-e2bf-c56a050d43d1"}
	I1002 12:43:58.146235  585741 start.go:140] virtualization: kvm guest
	I1002 12:43:58.148027  585741 out.go:179] * [NoKubernetes-115222] minikube v1.37.0 on Ubuntu 22.04 (kvm/amd64)
	I1002 12:43:58.149166  585741 out.go:179]   - MINIKUBE_LOCATION=21139
	I1002 12:43:58.149202  585741 notify.go:220] Checking for updates...
	I1002 12:43:58.151111  585741 out.go:179]   - MINIKUBE_SUPPRESS_DOCKER_PERFORMANCE=true
	I1002 12:43:58.152058  585741 out.go:179]   - KUBECONFIG=/home/jenkins/minikube-integration/21139-381342/kubeconfig
	I1002 12:43:58.153034  585741 out.go:179]   - MINIKUBE_HOME=/home/jenkins/minikube-integration/21139-381342/.minikube
	I1002 12:43:58.157294  585741 out.go:179]   - MINIKUBE_BIN=out/minikube-linux-amd64
	I1002 12:43:58.158404  585741 out.go:179]   - MINIKUBE_FORCE_SYSTEMD=
	I1002 12:43:56.113021  581986 kic_runner.go:191] docker (temp): /home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa.pub --> /home/docker/.ssh/authorized_keys (381 bytes)
	I1002 12:43:56.140789  581986 cli_runner.go:164] Run: docker container inspect missing-upgrade-404552 --format={{.State.Status}}
	I1002 12:43:56.158961  581986 kic_runner.go:93] Run: chown docker:docker /home/docker/.ssh/authorized_keys
	I1002 12:43:56.158975  581986 kic_runner.go:114] Args: [docker exec --privileged missing-upgrade-404552 chown docker:docker /home/docker/.ssh/authorized_keys]
	I1002 12:43:56.208006  581986 cli_runner.go:164] Run: docker container inspect missing-upgrade-404552 --format={{.State.Status}}
	I1002 12:43:56.228078  581986 machine.go:88] provisioning docker machine ...
	I1002 12:43:56.228119  581986 ubuntu.go:169] provisioning hostname "missing-upgrade-404552"
	I1002 12:43:56.228181  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:56.247049  581986 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:56.248203  581986 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x808a40] 0x80b720 <nil>  [] 0s} 127.0.0.1 33385 <nil> <nil>}
	I1002 12:43:56.248220  581986 main.go:141] libmachine: About to run SSH command:
	sudo hostname missing-upgrade-404552 && echo "missing-upgrade-404552" | sudo tee /etc/hostname
	I1002 12:43:56.379515  581986 main.go:141] libmachine: SSH cmd err, output: <nil>: missing-upgrade-404552
	
	I1002 12:43:56.379591  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:56.398347  581986 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:56.398900  581986 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x808a40] 0x80b720 <nil>  [] 0s} 127.0.0.1 33385 <nil> <nil>}
	I1002 12:43:56.398922  581986 main.go:141] libmachine: About to run SSH command:
	
			if ! grep -xq '.*\smissing-upgrade-404552' /etc/hosts; then
				if grep -xq '127.0.1.1\s.*' /etc/hosts; then
					sudo sed -i 's/^127.0.1.1\s.*/127.0.1.1 missing-upgrade-404552/g' /etc/hosts;
				else 
					echo '127.0.1.1 missing-upgrade-404552' | sudo tee -a /etc/hosts; 
				fi
			fi
	I1002 12:43:56.515082  581986 main.go:141] libmachine: SSH cmd err, output: <nil>: 
	I1002 12:43:56.515103  581986 ubuntu.go:175] set auth options {CertDir:/home/jenkins/minikube-integration/21139-381342/.minikube CaCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem CaPrivateKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem CaCertRemotePath:/etc/docker/ca.pem ServerCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ServerKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem ClientKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem ServerCertRemotePath:/etc/docker/server.pem ServerKeyRemotePath:/etc/docker/server-key.pem ClientCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem ServerCertSANs:[] StorePath:/home/jenkins/minikube-integration/21139-381342/.minikube}
	I1002 12:43:56.515141  581986 ubuntu.go:177] setting up certificates
	I1002 12:43:56.515151  581986 provision.go:83] configureAuth start
	I1002 12:43:56.515204  581986 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" missing-upgrade-404552
	I1002 12:43:56.534704  581986 provision.go:138] copyHostCerts
	I1002 12:43:56.534762  581986 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem, removing ...
	I1002 12:43:56.534772  581986 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem
	I1002 12:43:56.534864  581986 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem (1082 bytes)
	I1002 12:43:56.535019  581986 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem, removing ...
	I1002 12:43:56.535027  581986 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem
	I1002 12:43:56.535078  581986 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem (1123 bytes)
	I1002 12:43:56.535178  581986 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem, removing ...
	I1002 12:43:56.535198  581986 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem
	I1002 12:43:56.535249  581986 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem (1675 bytes)
	I1002 12:43:56.535380  581986 provision.go:112] generating server cert: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ca-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem private-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem org=jenkins.missing-upgrade-404552 san=[192.168.76.2 127.0.0.1 localhost 127.0.0.1 minikube missing-upgrade-404552]
	I1002 12:43:56.697791  581986 provision.go:172] copyRemoteCerts
	I1002 12:43:56.697863  581986 ssh_runner.go:195] Run: sudo mkdir -p /etc/docker /etc/docker /etc/docker
	I1002 12:43:56.697903  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:56.715683  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:56.800884  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem --> /etc/docker/server-key.pem (1679 bytes)
	I1002 12:43:56.826192  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /etc/docker/ca.pem (1082 bytes)
	I1002 12:43:56.849217  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem --> /etc/docker/server.pem (1241 bytes)
	I1002 12:43:56.871626  581986 provision.go:86] duration metric: configureAuth took 356.465021ms
	I1002 12:43:56.871642  581986 ubuntu.go:193] setting minikube options for container-runtime
	I1002 12:43:56.871791  581986 config.go:182] Loaded profile config "missing-upgrade-404552": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:56.871796  581986 machine.go:91] provisioned docker machine in 643.706565ms
	I1002 12:43:56.871801  581986 client.go:171] LocalClient.Create took 5.519165598s
	I1002 12:43:56.871817  581986 start.go:167] duration metric: libmachine.API.Create for "missing-upgrade-404552" took 5.51922352s
	I1002 12:43:56.871822  581986 start.go:300] post-start starting for "missing-upgrade-404552" (driver="docker")
	I1002 12:43:56.871849  581986 start.go:329] creating required directories: [/etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs]
	I1002 12:43:56.871902  581986 ssh_runner.go:195] Run: sudo mkdir -p /etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs
	I1002 12:43:56.871947  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:56.890200  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:56.982746  581986 ssh_runner.go:195] Run: cat /etc/os-release
	I1002 12:43:56.985888  581986 main.go:141] libmachine: Couldn't set key VERSION_CODENAME, no corresponding struct field found
	I1002 12:43:56.985921  581986 main.go:141] libmachine: Couldn't set key PRIVACY_POLICY_URL, no corresponding struct field found
	I1002 12:43:56.985938  581986 main.go:141] libmachine: Couldn't set key UBUNTU_CODENAME, no corresponding struct field found
	I1002 12:43:56.985944  581986 info.go:137] Remote host: Ubuntu 22.04.3 LTS
	I1002 12:43:56.985953  581986 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/addons for local assets ...
	I1002 12:43:56.986014  581986 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/files for local assets ...
	I1002 12:43:56.986100  581986 filesync.go:149] local asset: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem -> 3849552.pem in /etc/ssl/certs
	I1002 12:43:56.986211  581986 ssh_runner.go:195] Run: sudo mkdir -p /etc/ssl/certs
	I1002 12:43:56.994580  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem --> /etc/ssl/certs/3849552.pem (1708 bytes)
	I1002 12:43:57.020319  581986 start.go:303] post-start completed in 148.471394ms
	I1002 12:43:57.020659  581986 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" missing-upgrade-404552
	I1002 12:43:57.038303  581986 profile.go:148] Saving config to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/config.json ...
	I1002 12:43:57.038500  581986 ssh_runner.go:195] Run: sh -c "df -h /var | awk 'NR==2{print $5}'"
	I1002 12:43:57.038529  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:57.055262  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:57.135385  581986 ssh_runner.go:195] Run: sh -c "df -BG /var | awk 'NR==2{print $4}'"
	I1002 12:43:57.139586  581986 start.go:128] duration metric: createHost completed in 5.789433984s
	I1002 12:43:57.139600  581986 start.go:83] releasing machines lock for "missing-upgrade-404552", held for 5.789554894s
	I1002 12:43:57.139657  581986 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" missing-upgrade-404552
	I1002 12:43:57.156682  581986 ssh_runner.go:195] Run: cat /version.json
	I1002 12:43:57.156713  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:57.156773  581986 ssh_runner.go:195] Run: curl -sS -m 2 https://registry.k8s.io/
	I1002 12:43:57.156843  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:57.175099  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:57.175452  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:57.348369  581986 ssh_runner.go:195] Run: systemctl --version
	I1002 12:43:57.353082  581986 ssh_runner.go:195] Run: sh -c "stat /etc/cni/net.d/*loopback.conf*"
	I1002 12:43:57.357358  581986 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f -name *loopback.conf* -not -name *.mk_disabled -exec sh -c "grep -q loopback {} && ( grep -q name {} || sudo sed -i '/"type": "loopback"/i \ \ \ \ "name": "loopback",' {} ) && sudo sed -i 's|"cniVersion": ".*"|"cniVersion": "1.0.0"|g' {}" ;
	I1002 12:43:57.384724  581986 cni.go:230] loopback cni configuration patched: "/etc/cni/net.d/*loopback.conf*" found
	I1002 12:43:57.384773  581986 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f ( ( -name *bridge* -or -name *podman* ) -and -not -name *.mk_disabled ) -printf "%p, " -exec sh -c "sudo mv {} {}.mk_disabled" ;
	I1002 12:43:57.409586  581986 cni.go:262] disabled [/etc/cni/net.d/87-podman-bridge.conflist, /etc/cni/net.d/100-crio-bridge.conf] bridge cni config(s)
	I1002 12:43:57.409603  581986 start.go:472] detecting cgroup driver to use...
	I1002 12:43:57.409629  581986 detect.go:199] detected "systemd" cgroup driver on host os
	I1002 12:43:57.409668  581986 ssh_runner.go:195] Run: sudo systemctl stop -f crio
	I1002 12:43:57.421371  581986 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service crio
	I1002 12:43:57.431690  581986 docker.go:203] disabling cri-docker service (if available) ...
	I1002 12:43:57.431744  581986 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.socket
	I1002 12:43:57.444305  581986 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.service
	I1002 12:43:57.457523  581986 ssh_runner.go:195] Run: sudo systemctl disable cri-docker.socket
	I1002 12:43:57.524054  581986 ssh_runner.go:195] Run: sudo systemctl mask cri-docker.service
	I1002 12:43:57.598654  581986 docker.go:219] disabling docker service ...
	I1002 12:43:57.598703  581986 ssh_runner.go:195] Run: sudo systemctl stop -f docker.socket
	I1002 12:43:57.616246  581986 ssh_runner.go:195] Run: sudo systemctl stop -f docker.service
	I1002 12:43:57.627447  581986 ssh_runner.go:195] Run: sudo systemctl disable docker.socket
	I1002 12:43:57.704488  581986 ssh_runner.go:195] Run: sudo systemctl mask docker.service
	I1002 12:43:57.785554  581986 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service docker
	I1002 12:43:57.797381  581986 ssh_runner.go:195] Run: /bin/bash -c "sudo mkdir -p /etc && printf %s "runtime-endpoint: unix:///run/containerd/containerd.sock
	" | sudo tee /etc/crictl.yaml"
	I1002 12:43:57.813760  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)sandbox_image = .*$|\1sandbox_image = "registry.k8s.io/pause:3.9"|' /etc/containerd/config.toml"
	I1002 12:43:57.829274  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)restrict_oom_score_adj = .*$|\1restrict_oom_score_adj = false|' /etc/containerd/config.toml"
	I1002 12:43:57.839084  581986 containerd.go:145] configuring containerd to use "systemd" as cgroup driver...
	I1002 12:43:57.839137  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)SystemdCgroup = .*$|\1SystemdCgroup = true|g' /etc/containerd/config.toml"
	I1002 12:43:57.848695  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runtime.v1.linux"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:43:57.858251  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i '/systemd_cgroup/d' /etc/containerd/config.toml"
	I1002 12:43:57.868364  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runc.v1"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:43:57.878878  581986 ssh_runner.go:195] Run: sh -c "sudo rm -rf /etc/cni/net.mk"
	I1002 12:43:57.888799  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)conf_dir = .*$|\1conf_dir = "/etc/cni/net.d"|g' /etc/containerd/config.toml"
	I1002 12:43:57.898544  581986 ssh_runner.go:195] Run: sudo sysctl net.bridge.bridge-nf-call-iptables
	I1002 12:43:57.929002  581986 ssh_runner.go:195] Run: sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
	I1002 12:43:57.938069  581986 ssh_runner.go:195] Run: sudo systemctl daemon-reload
	I1002 12:43:58.002820  581986 ssh_runner.go:195] Run: sudo systemctl restart containerd
	I1002 12:43:58.106242  581986 start.go:519] Will wait 60s for socket path /run/containerd/containerd.sock
	I1002 12:43:58.106289  581986 ssh_runner.go:195] Run: stat /run/containerd/containerd.sock
	I1002 12:43:58.110402  581986 start.go:540] Will wait 60s for crictl version
	I1002 12:43:58.110447  581986 ssh_runner.go:195] Run: which crictl
	I1002 12:43:58.114209  581986 ssh_runner.go:195] Run: sudo /usr/bin/crictl version
	I1002 12:43:58.149277  581986 start.go:556] Version:  0.1.0
	RuntimeName:  containerd
	RuntimeVersion:  1.6.24
	RuntimeApiVersion:  v1
	I1002 12:43:58.149349  581986 ssh_runner.go:195] Run: containerd --version
	I1002 12:43:58.174699  581986 ssh_runner.go:195] Run: containerd --version
	I1002 12:43:58.203536  581986 out.go:177] * Preparing Kubernetes v1.28.3 on containerd 1.6.24 ...
	I1002 12:43:58.160098  585741 config.go:182] Loaded profile config "missing-upgrade-404552": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:58.160250  585741 config.go:182] Loaded profile config "running-upgrade-583950": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:58.160372  585741 config.go:182] Loaded profile config "stopped-upgrade-250350": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:58.160403  585741 start.go:1898] No Kubernetes flag is set, setting Kubernetes version to v0.0.0
	I1002 12:43:58.160483  585741 driver.go:421] Setting default libvirt URI to qemu:///system
	I1002 12:43:58.184848  585741 docker.go:123] docker version: linux-28.4.0:Docker Engine - Community
	I1002 12:43:58.184952  585741 cli_runner.go:164] Run: docker system info --format "{{json .}}"
	I1002 12:43:58.239813  585741 info.go:266] docker info: {ID:TS6T:UINC:MIYS:RZPA:KS6T:4JQK:7JHN:D6RA:LDP2:MHAE:G32M:C5NQ Containers:3 ContainersRunning:3 ContainersPaused:0 ContainersStopped:0 Images:4 Driver:overlay2 DriverStatus:[[Backing Filesystem extfs] [Supports d_type true] [Using metacopy false] [Native Overlay Diff true] [userxattr false]] SystemStatus:<nil> Plugins:{Volume:[local] Network:[bridge host ipvlan macvlan null overlay] Authorization:<nil> Log:[awslogs fluentd gcplogs gelf journald json-file local splunk syslog]} MemoryLimit:true SwapLimit:true KernelMemory:false KernelMemoryTCP:false CPUCfsPeriod:true CPUCfsQuota:true CPUShares:true CPUSet:true PidsLimit:true IPv4Forwarding:true BridgeNfIptables:false BridgeNfIP6Tables:false Debug:false NFd:68 OomKillDisable:false NGoroutines:77 SystemTime:2025-10-02 12:43:58.229497835 +0000 UTC LoggingDriver:json-file CgroupDriver:cgroupfs NEventsListener:0 KernelVersion:6.8.0-1041-gcp OperatingSystem:Ubuntu 22.04.5 LTS OSType:linux Architecture:
x86_64 IndexServerAddress:https://index.docker.io/v1/ RegistryConfig:{AllowNondistributableArtifactsCIDRs:[] AllowNondistributableArtifactsHostnames:[] InsecureRegistryCIDRs:[::1/128 127.0.0.0/8] IndexConfigs:{DockerIo:{Name:docker.io Mirrors:[] Secure:true Official:true}} Mirrors:[]} NCPU:8 MemTotal:33652174848 GenericResources:<nil> DockerRootDir:/var/lib/docker HTTPProxy: HTTPSProxy: NoProxy: Name:ubuntu-20-agent-6 Labels:[] ExperimentalBuild:false ServerVersion:28.4.0 ClusterStore: ClusterAdvertise: Runtimes:{Runc:{Path:runc}} DefaultRuntime:runc Swarm:{NodeID: NodeAddr: LocalNodeState:inactive ControlAvailable:false Error: RemoteManagers:<nil>} LiveRestoreEnabled:false Isolation: InitBinary:docker-init ContainerdCommit:{ID:b98a3aace656320842a23f4a392a33f46af97866 Expected:} RuncCommit:{ID:v1.3.0-0-g4ca628d1 Expected:} InitCommit:{ID:de40ad0 Expected:} SecurityOptions:[name=apparmor name=seccomp,profile=builtin name=cgroupns] ProductLicense: Warnings:<nil> ServerErrors:[] ClientInfo:{Debug:false Plugins:[
map[Name:buildx Path:/usr/libexec/docker/cli-plugins/docker-buildx SchemaVersion:0.1.0 ShortDescription:Docker Buildx Vendor:Docker Inc. Version:v0.29.0] map[Name:compose Path:/usr/libexec/docker/cli-plugins/docker-compose SchemaVersion:0.1.0 ShortDescription:Docker Compose Vendor:Docker Inc. Version:v2.39.4] map[Name:model Path:/usr/libexec/docker/cli-plugins/docker-model SchemaVersion:0.1.0 ShortDescription:Docker Model Runner Vendor:Docker Inc. Version:v0.1.40] map[Name:scan Path:/usr/libexec/docker/cli-plugins/docker-scan SchemaVersion:0.1.0 ShortDescription:Docker Scan Vendor:Docker Inc. Version:v0.23.0]] Warnings:<nil>}}
	I1002 12:43:58.239954  585741 docker.go:318] overlay module found
	I1002 12:43:58.242138  585741 out.go:179] * Using the docker driver based on user configuration
	I1002 12:43:58.243209  585741 start.go:304] selected driver: docker
	I1002 12:43:58.243231  585741 start.go:924] validating driver "docker" against <nil>
	I1002 12:43:58.243242  585741 start.go:935] status for docker: {Installed:true Healthy:true Running:false NeedsImprovement:false Error:<nil> Reason: Fix: Doc: Version:}
	I1002 12:43:58.243815  585741 cli_runner.go:164] Run: docker system info --format "{{json .}}"
	I1002 12:43:58.299577  585741 info.go:266] docker info: {ID:TS6T:UINC:MIYS:RZPA:KS6T:4JQK:7JHN:D6RA:LDP2:MHAE:G32M:C5NQ Containers:3 ContainersRunning:3 ContainersPaused:0 ContainersStopped:0 Images:4 Driver:overlay2 DriverStatus:[[Backing Filesystem extfs] [Supports d_type true] [Using metacopy false] [Native Overlay Diff true] [userxattr false]] SystemStatus:<nil> Plugins:{Volume:[local] Network:[bridge host ipvlan macvlan null overlay] Authorization:<nil> Log:[awslogs fluentd gcplogs gelf journald json-file local splunk syslog]} MemoryLimit:true SwapLimit:true KernelMemory:false KernelMemoryTCP:false CPUCfsPeriod:true CPUCfsQuota:true CPUShares:true CPUSet:true PidsLimit:true IPv4Forwarding:true BridgeNfIptables:false BridgeNfIP6Tables:false Debug:false NFd:68 OomKillDisable:false NGoroutines:77 SystemTime:2025-10-02 12:43:58.288929094 +0000 UTC LoggingDriver:json-file CgroupDriver:cgroupfs NEventsListener:0 KernelVersion:6.8.0-1041-gcp OperatingSystem:Ubuntu 22.04.5 LTS OSType:linux Architecture:
x86_64 IndexServerAddress:https://index.docker.io/v1/ RegistryConfig:{AllowNondistributableArtifactsCIDRs:[] AllowNondistributableArtifactsHostnames:[] InsecureRegistryCIDRs:[::1/128 127.0.0.0/8] IndexConfigs:{DockerIo:{Name:docker.io Mirrors:[] Secure:true Official:true}} Mirrors:[]} NCPU:8 MemTotal:33652174848 GenericResources:<nil> DockerRootDir:/var/lib/docker HTTPProxy: HTTPSProxy: NoProxy: Name:ubuntu-20-agent-6 Labels:[] ExperimentalBuild:false ServerVersion:28.4.0 ClusterStore: ClusterAdvertise: Runtimes:{Runc:{Path:runc}} DefaultRuntime:runc Swarm:{NodeID: NodeAddr: LocalNodeState:inactive ControlAvailable:false Error: RemoteManagers:<nil>} LiveRestoreEnabled:false Isolation: InitBinary:docker-init ContainerdCommit:{ID:b98a3aace656320842a23f4a392a33f46af97866 Expected:} RuncCommit:{ID:v1.3.0-0-g4ca628d1 Expected:} InitCommit:{ID:de40ad0 Expected:} SecurityOptions:[name=apparmor name=seccomp,profile=builtin name=cgroupns] ProductLicense: Warnings:<nil> ServerErrors:[] ClientInfo:{Debug:false Plugins:[
map[Name:buildx Path:/usr/libexec/docker/cli-plugins/docker-buildx SchemaVersion:0.1.0 ShortDescription:Docker Buildx Vendor:Docker Inc. Version:v0.29.0] map[Name:compose Path:/usr/libexec/docker/cli-plugins/docker-compose SchemaVersion:0.1.0 ShortDescription:Docker Compose Vendor:Docker Inc. Version:v2.39.4] map[Name:model Path:/usr/libexec/docker/cli-plugins/docker-model SchemaVersion:0.1.0 ShortDescription:Docker Model Runner Vendor:Docker Inc. Version:v0.1.40] map[Name:scan Path:/usr/libexec/docker/cli-plugins/docker-scan SchemaVersion:0.1.0 ShortDescription:Docker Scan Vendor:Docker Inc. Version:v0.23.0]] Warnings:<nil>}}
	I1002 12:43:58.299663  585741 start.go:1898] No Kubernetes flag is set, setting Kubernetes version to v0.0.0
	I1002 12:43:58.299741  585741 start_flags.go:327] no existing cluster config was found, will generate one from the flags 
	I1002 12:43:58.300012  585741 start_flags.go:974] Wait components to verify : map[apiserver:true system_pods:true]
	I1002 12:43:58.301478  585741 out.go:179] * Using Docker driver with root privileges
	I1002 12:43:58.302563  585741 cni.go:84] Creating CNI manager for ""
	I1002 12:43:58.302631  585741 cni.go:143] "docker" driver + "containerd" runtime found, recommending kindnet
	I1002 12:43:58.302645  585741 start_flags.go:336] Found "CNI" CNI - setting NetworkPlugin=cni
	I1002 12:43:58.302670  585741 start.go:1898] No Kubernetes flag is set, setting Kubernetes version to v0.0.0
	I1002 12:43:58.302712  585741 start.go:348] cluster config:
	{Name:NoKubernetes-115222 KeepContext:false EmbedCerts:false MinikubeISO: KicBaseImage:gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 Memory:3072 CPUs:2 DiskSize:20000 Driver:docker HyperkitVpnKitSock: HyperkitVSockPorts:[] DockerEnv:[] ContainerVolumeMounts:[] InsecureRegistry:[] RegistryMirror:[] HostOnlyCIDR:192.168.59.1/24 HypervVirtualSwitch: HypervUseExternalSwitch:false HypervExternalAdapter: KVMNetwork:default KVMQemuURI:qemu:///system KVMGPU:false KVMHidden:false KVMNUMACount:1 APIServerPort:8443 DockerOpt:[] DisableDriverMounts:false NFSShare:[] NFSSharesRoot:/nfsshares UUID: NoVTXCheck:false DNSProxy:false HostDNSResolver:true HostOnlyNicType:virtio NatNicType:virtio SSHIPAddress: SSHUser:root SSHKey: SSHPort:22 KubernetesConfig:{KubernetesVersion:v0.0.0 ClusterName:NoKubernetes-115222 Namespace:default APIServerHAVIP: APIServerName:minikubeCA APIServerNames:[] APIServerIPs:[] DNSDomain:cluster.local ContainerRuntime:containerd C
RISocket: NetworkPlugin:cni FeatureGates: ServiceCIDR:10.96.0.0/12 ImageRepository: LoadBalancerStartIP: LoadBalancerEndIP: CustomIngressCert: RegistryAliases: ExtraOptions:[] ShouldLoadCachedImages:true EnableDefaultCNI:false CNI:} Nodes:[{Name: IP: Port:8443 KubernetesVersion:v0.0.0 ContainerRuntime:containerd ControlPlane:true Worker:true}] Addons:map[] CustomAddonImages:map[] CustomAddonRegistries:map[] VerifyComponents:map[apiserver:true system_pods:true] StartHostTimeout:6m0s ScheduledStop:<nil> ExposedPorts:[] ListenAddress: Network: Subnet: MultiNodeRequested:false ExtraDisks:0 CertExpiration:26280h0m0s MountString: Mount9PVersion:9p2000.L MountGID:docker MountIP: MountMSize:262144 MountOptions:[] MountPort:0 MountType:9p MountUID:docker BinaryMirror: DisableOptimizations:false DisableMetrics:false DisableCoreDNSLog:false CustomQemuFirmwarePath: SocketVMnetClientPath: SocketVMnetPath: StaticIP: SSHAuthSock: SSHAgentPID:0 GPUs: AutoPauseInterval:1m0s}
	I1002 12:43:58.303779  585741 out.go:179] * Starting minikube without Kubernetes in cluster NoKubernetes-115222
	I1002 12:43:58.304702  585741 cache.go:133] Beginning downloading kic base image for docker with containerd
	I1002 12:43:58.308050  585741 out.go:179] * Pulling base image v0.0.48 ...
	I1002 12:43:58.309193  585741 cache.go:58] Skipping Kubernetes image caching due to --no-kubernetes flag
	I1002 12:43:58.309306  585741 image.go:81] Checking for gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 in local docker daemon
	I1002 12:43:58.309348  585741 profile.go:143] Saving config to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/NoKubernetes-115222/config.json ...
	I1002 12:43:58.309398  585741 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/NoKubernetes-115222/config.json: {Name:mkf16e835ee733faaa5453f2629f21d135e155cb Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.331669  585741 image.go:100] Found gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 in local docker daemon, skipping pull
	I1002 12:43:58.331703  585741 cache.go:157] gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 exists in daemon, skipping load
	I1002 12:43:58.331726  585741 cache.go:242] Successfully downloaded all kic artifacts
	I1002 12:43:58.331756  585741 start.go:360] acquireMachinesLock for NoKubernetes-115222: {Name:mkcdc227f947561006dc8621b5597787add0762d Clock:{} Delay:500ms Timeout:10m0s Cancel:<nil>}
	I1002 12:43:58.331820  585741 start.go:364] duration metric: took 43.427µs to acquireMachinesLock for "NoKubernetes-115222"
	I1002 12:43:58.331854  585741 start.go:93] Provisioning new machine with config: &{Name:NoKubernetes-115222 KeepContext:false EmbedCerts:false MinikubeISO: KicBaseImage:gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 Memory:3072 CPUs:2 DiskSize:20000 Driver:docker HyperkitVpnKitSock: HyperkitVSockPorts:[] DockerEnv:[] ContainerVolumeMounts:[] InsecureRegistry:[] RegistryMirror:[] HostOnlyCIDR:192.168.59.1/24 HypervVirtualSwitch: HypervUseExternalSwitch:false HypervExternalAdapter: KVMNetwork:default KVMQemuURI:qemu:///system KVMGPU:false KVMHidden:false KVMNUMACount:1 APIServerPort:8443 DockerOpt:[] DisableDriverMounts:false NFSShare:[] NFSSharesRoot:/nfsshares UUID: NoVTXCheck:false DNSProxy:false HostDNSResolver:true HostOnlyNicType:virtio NatNicType:virtio SSHIPAddress: SSHUser:root SSHKey: SSHPort:22 KubernetesConfig:{KubernetesVersion:v0.0.0 ClusterName:NoKubernetes-115222 Namespace:default APIServerHAVIP: APIServerName:minikubeCA APISe
rverNames:[] APIServerIPs:[] DNSDomain:cluster.local ContainerRuntime:containerd CRISocket: NetworkPlugin:cni FeatureGates: ServiceCIDR:10.96.0.0/12 ImageRepository: LoadBalancerStartIP: LoadBalancerEndIP: CustomIngressCert: RegistryAliases: ExtraOptions:[] ShouldLoadCachedImages:true EnableDefaultCNI:false CNI:} Nodes:[{Name: IP: Port:8443 KubernetesVersion:v0.0.0 ContainerRuntime:containerd ControlPlane:true Worker:true}] Addons:map[] CustomAddonImages:map[] CustomAddonRegistries:map[] VerifyComponents:map[apiserver:true system_pods:true] StartHostTimeout:6m0s ScheduledStop:<nil> ExposedPorts:[] ListenAddress: Network: Subnet: MultiNodeRequested:false ExtraDisks:0 CertExpiration:26280h0m0s MountString: Mount9PVersion:9p2000.L MountGID:docker MountIP: MountMSize:262144 MountOptions:[] MountPort:0 MountType:9p MountUID:docker BinaryMirror: DisableOptimizations:false DisableMetrics:false DisableCoreDNSLog:false CustomQemuFirmwarePath: SocketVMnetClientPath: SocketVMnetPath: StaticIP: SSHAuthSock: SSHAgentPID:0
GPUs: AutoPauseInterval:1m0s} &{Name: IP: Port:8443 KubernetesVersion:v0.0.0 ContainerRuntime:containerd ControlPlane:true Worker:true}
	I1002 12:43:58.331941  585741 start.go:125] createHost starting for "" (driver="docker")
	I1002 12:43:55.053661  583227 out.go:252] * Restarting existing docker container for "stopped-upgrade-250350" ...
	I1002 12:43:55.053739  583227 cli_runner.go:164] Run: docker start stopped-upgrade-250350
	I1002 12:43:55.339095  583227 cli_runner.go:164] Run: docker container inspect stopped-upgrade-250350 --format={{.State.Status}}
	I1002 12:43:55.363276  583227 kic.go:430] container "stopped-upgrade-250350" state is running.
	I1002 12:43:55.363776  583227 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" stopped-upgrade-250350
	I1002 12:43:55.386979  583227 profile.go:143] Saving config to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/stopped-upgrade-250350/config.json ...
	I1002 12:43:55.387263  583227 machine.go:93] provisionDockerMachine start ...
	I1002 12:43:55.387348  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:55.411893  583227 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:55.412281  583227 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33380 <nil> <nil>}
	I1002 12:43:55.412318  583227 main.go:141] libmachine: About to run SSH command:
	hostname
	I1002 12:43:55.413003  583227 main.go:141] libmachine: Error dialing TCP: ssh: handshake failed: read tcp 127.0.0.1:43828->127.0.0.1:33380: read: connection reset by peer
	I1002 12:43:58.531506  583227 main.go:141] libmachine: SSH cmd err, output: <nil>: stopped-upgrade-250350
	
	I1002 12:43:58.531541  583227 ubuntu.go:182] provisioning hostname "stopped-upgrade-250350"
	I1002 12:43:58.531614  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:58.551286  583227 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:58.551592  583227 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33380 <nil> <nil>}
	I1002 12:43:58.551620  583227 main.go:141] libmachine: About to run SSH command:
	sudo hostname stopped-upgrade-250350 && echo "stopped-upgrade-250350" | sudo tee /etc/hostname
	I1002 12:43:58.680783  583227 main.go:141] libmachine: SSH cmd err, output: <nil>: stopped-upgrade-250350
	
	I1002 12:43:58.680891  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:58.700423  583227 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:58.700708  583227 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33380 <nil> <nil>}
	I1002 12:43:58.700734  583227 main.go:141] libmachine: About to run SSH command:
	
			if ! grep -xq '.*\sstopped-upgrade-250350' /etc/hosts; then
				if grep -xq '127.0.1.1\s.*' /etc/hosts; then
					sudo sed -i 's/^127.0.1.1\s.*/127.0.1.1 stopped-upgrade-250350/g' /etc/hosts;
				else 
					echo '127.0.1.1 stopped-upgrade-250350' | sudo tee -a /etc/hosts; 
				fi
			fi
	I1002 12:43:58.824367  583227 main.go:141] libmachine: SSH cmd err, output: <nil>: 
	I1002 12:43:58.824399  583227 ubuntu.go:188] set auth options {CertDir:/home/jenkins/minikube-integration/21139-381342/.minikube CaCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem CaPrivateKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem CaCertRemotePath:/etc/docker/ca.pem ServerCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ServerKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem ClientKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem ServerCertRemotePath:/etc/docker/server.pem ServerKeyRemotePath:/etc/docker/server-key.pem ClientCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem ServerCertSANs:[] StorePath:/home/jenkins/minikube-integration/21139-381342/.minikube}
	I1002 12:43:58.824423  583227 ubuntu.go:190] setting up certificates
	I1002 12:43:58.824433  583227 provision.go:84] configureAuth start
	I1002 12:43:58.824510  583227 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" stopped-upgrade-250350
	I1002 12:43:58.844605  583227 provision.go:143] copyHostCerts
	I1002 12:43:58.844675  583227 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem, removing ...
	I1002 12:43:58.844700  583227 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem
	I1002 12:43:58.844760  583227 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem (1082 bytes)
	I1002 12:43:58.844913  583227 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem, removing ...
	I1002 12:43:58.844927  583227 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem
	I1002 12:43:58.844973  583227 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem (1123 bytes)
	I1002 12:43:58.845069  583227 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem, removing ...
	I1002 12:43:58.845079  583227 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem
	I1002 12:43:58.845118  583227 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem (1675 bytes)
	I1002 12:43:58.845208  583227 provision.go:117] generating server cert: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ca-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem private-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem org=jenkins.stopped-upgrade-250350 san=[127.0.0.1 192.168.103.2 localhost minikube stopped-upgrade-250350]
	I1002 12:43:58.204603  581986 cli_runner.go:164] Run: docker network inspect missing-upgrade-404552 --format "{"Name": "{{.Name}}","Driver": "{{.Driver}}","Subnet": "{{range .IPAM.Config}}{{.Subnet}}{{end}}","Gateway": "{{range .IPAM.Config}}{{.Gateway}}{{end}}","MTU": {{if (index .Options "com.docker.network.driver.mtu")}}{{(index .Options "com.docker.network.driver.mtu")}}{{else}}0{{end}}, "ContainerIPs": [{{range $k,$v := .Containers }}"{{$v.IPv4Address}}",{{end}}]}"
	I1002 12:43:58.223594  581986 ssh_runner.go:195] Run: grep 192.168.76.1	host.minikube.internal$ /etc/hosts
	I1002 12:43:58.227716  581986 ssh_runner.go:195] Run: /bin/bash -c "{ grep -v $'\thost.minikube.internal$' "/etc/hosts"; echo "192.168.76.1	host.minikube.internal"; } > /tmp/h.$$; sudo cp /tmp/h.$$ "/etc/hosts""
	I1002 12:43:58.240322  581986 preload.go:132] Checking if preload exists for k8s version v1.28.3 and runtime containerd
	I1002 12:43:58.240665  581986 ssh_runner.go:195] Run: sudo crictl images --output json
	I1002 12:43:58.277665  581986 containerd.go:604] all images are preloaded for containerd runtime.
	I1002 12:43:58.277685  581986 containerd.go:518] Images already preloaded, skipping extraction
	I1002 12:43:58.277745  581986 ssh_runner.go:195] Run: sudo crictl images --output json
	I1002 12:43:58.313507  581986 containerd.go:604] all images are preloaded for containerd runtime.
	I1002 12:43:58.313522  581986 cache_images.go:84] Images are preloaded, skipping loading
	I1002 12:43:58.313581  581986 ssh_runner.go:195] Run: sudo crictl info
	I1002 12:43:58.352177  581986 cni.go:84] Creating CNI manager for ""
	I1002 12:43:58.352192  581986 cni.go:143] "docker" driver + "containerd" runtime found, recommending kindnet
	I1002 12:43:58.352221  581986 kubeadm.go:87] Using pod CIDR: 10.244.0.0/16
	I1002 12:43:58.352243  581986 kubeadm.go:176] kubeadm options: {CertDir:/var/lib/minikube/certs ServiceCIDR:10.96.0.0/12 PodSubnet:10.244.0.0/16 AdvertiseAddress:192.168.76.2 APIServerPort:8443 KubernetesVersion:v1.28.3 EtcdDataDir:/var/lib/minikube/etcd EtcdExtraArgs:map[] ClusterName:missing-upgrade-404552 NodeName:missing-upgrade-404552 DNSDomain:cluster.local CRISocket:/run/containerd/containerd.sock ImageRepository: ComponentOptions:[{Component:apiServer ExtraArgs:map[enable-admission-plugins:NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota] Pairs:map[certSANs:["127.0.0.1", "localhost", "192.168.76.2"]]} {Component:controllerManager ExtraArgs:map[allocate-node-cidrs:true leader-elect:false] Pairs:map[]} {Component:scheduler ExtraArgs:map[leader-elect:false] Pairs:map[]}] FeatureArgs:map[] NodeIP:192.168.76.2 CgroupDriver:systemd ClientCAFile:/var/lib/minikube/certs/ca.crt S
taticPodPath:/etc/kubernetes/manifests ControlPlaneAddress:control-plane.minikube.internal KubeProxyOptions:map[] ResolvConfSearchRegression:false KubeletConfigOpts:map[hairpinMode:hairpin-veth runtimeRequestTimeout:15m] PrependCriSocketUnix:true}
	I1002 12:43:58.352404  581986 kubeadm.go:181] kubeadm config:
	apiVersion: kubeadm.k8s.io/v1beta3
	kind: InitConfiguration
	localAPIEndpoint:
	  advertiseAddress: 192.168.76.2
	  bindPort: 8443
	bootstrapTokens:
	  - groups:
	      - system:bootstrappers:kubeadm:default-node-token
	    ttl: 24h0m0s
	    usages:
	      - signing
	      - authentication
	nodeRegistration:
	  criSocket: unix:///run/containerd/containerd.sock
	  name: "missing-upgrade-404552"
	  kubeletExtraArgs:
	    node-ip: 192.168.76.2
	  taints: []
	---
	apiVersion: kubeadm.k8s.io/v1beta3
	kind: ClusterConfiguration
	apiServer:
	  certSANs: ["127.0.0.1", "localhost", "192.168.76.2"]
	  extraArgs:
	    enable-admission-plugins: "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota"
	controllerManager:
	  extraArgs:
	    allocate-node-cidrs: "true"
	    leader-elect: "false"
	scheduler:
	  extraArgs:
	    leader-elect: "false"
	certificatesDir: /var/lib/minikube/certs
	clusterName: mk
	controlPlaneEndpoint: control-plane.minikube.internal:8443
	etcd:
	  local:
	    dataDir: /var/lib/minikube/etcd
	    extraArgs:
	      proxy-refresh-interval: "70000"
	kubernetesVersion: v1.28.3
	networking:
	  dnsDomain: cluster.local
	  podSubnet: "10.244.0.0/16"
	  serviceSubnet: 10.96.0.0/12
	---
	apiVersion: kubelet.config.k8s.io/v1beta1
	kind: KubeletConfiguration
	authentication:
	  x509:
	    clientCAFile: /var/lib/minikube/certs/ca.crt
	cgroupDriver: systemd
	hairpinMode: hairpin-veth
	runtimeRequestTimeout: 15m
	clusterDomain: "cluster.local"
	# disable disk resource management by default
	imageGCHighThresholdPercent: 100
	evictionHard:
	  nodefs.available: "0%"
	  nodefs.inodesFree: "0%"
	  imagefs.available: "0%"
	failSwapOn: false
	staticPodPath: /etc/kubernetes/manifests
	---
	apiVersion: kubeproxy.config.k8s.io/v1alpha1
	kind: KubeProxyConfiguration
	clusterCIDR: "10.244.0.0/16"
	metricsBindAddress: 0.0.0.0:10249
	conntrack:
	  maxPerCore: 0
	# Skip setting "net.netfilter.nf_conntrack_tcp_timeout_established"
	  tcpEstablishedTimeout: 0s
	# Skip setting "net.netfilter.nf_conntrack_tcp_timeout_close"
	  tcpCloseWaitTimeout: 0s
	
	I1002 12:43:58.352466  581986 kubeadm.go:976] kubelet [Unit]
	Wants=containerd.service
	
	[Service]
	ExecStart=
	ExecStart=/var/lib/minikube/binaries/v1.28.3/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --config=/var/lib/kubelet/config.yaml --container-runtime-endpoint=unix:///run/containerd/containerd.sock --hostname-override=missing-upgrade-404552 --kubeconfig=/etc/kubernetes/kubelet.conf --node-ip=192.168.76.2
	
	[Install]
	 config:
	{KubernetesVersion:v1.28.3 ClusterName:missing-upgrade-404552 Namespace:default APIServerName:minikubeCA APIServerNames:[] APIServerIPs:[] DNSDomain:cluster.local ContainerRuntime:containerd CRISocket: NetworkPlugin:cni FeatureGates: ServiceCIDR:10.96.0.0/12 ImageRepository: LoadBalancerStartIP: LoadBalancerEndIP: CustomIngressCert: RegistryAliases: ExtraOptions:[] ShouldLoadCachedImages:true EnableDefaultCNI:false CNI: NodeIP: NodePort:8443 NodeName:}
	I1002 12:43:58.352507  581986 ssh_runner.go:195] Run: sudo ls /var/lib/minikube/binaries/v1.28.3
	I1002 12:43:58.361810  581986 binaries.go:44] Found k8s binaries, skipping transfer
	I1002 12:43:58.361897  581986 ssh_runner.go:195] Run: sudo mkdir -p /etc/systemd/system/kubelet.service.d /lib/systemd/system /var/tmp/minikube
	I1002 12:43:58.371174  581986 ssh_runner.go:362] scp memory --> /etc/systemd/system/kubelet.service.d/10-kubeadm.conf (394 bytes)
	I1002 12:43:58.389160  581986 ssh_runner.go:362] scp memory --> /lib/systemd/system/kubelet.service (352 bytes)
	I1002 12:43:58.411246  581986 ssh_runner.go:362] scp memory --> /var/tmp/minikube/kubeadm.yaml.new (2110 bytes)
	I1002 12:43:58.431430  581986 ssh_runner.go:195] Run: grep 192.168.76.2	control-plane.minikube.internal$ /etc/hosts
	I1002 12:43:58.435174  581986 ssh_runner.go:195] Run: /bin/bash -c "{ grep -v $'\tcontrol-plane.minikube.internal$' "/etc/hosts"; echo "192.168.76.2	control-plane.minikube.internal"; } > /tmp/h.$$; sudo cp /tmp/h.$$ "/etc/hosts""
	I1002 12:43:58.446345  581986 certs.go:56] Setting up /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552 for IP: 192.168.76.2
	I1002 12:43:58.446370  581986 certs.go:190] acquiring lock for shared ca certs: {Name:mk0b4dee9533eb55e655fb5dad1d990d151f0d2d Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.446509  581986 certs.go:199] skipping minikubeCA CA generation: /home/jenkins/minikube-integration/21139-381342/.minikube/ca.key
	I1002 12:43:58.446557  581986 certs.go:199] skipping proxyClientCA CA generation: /home/jenkins/minikube-integration/21139-381342/.minikube/proxy-client-ca.key
	I1002 12:43:58.446607  581986 certs.go:319] generating minikube-user signed cert: /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.key
	I1002 12:43:58.446615  581986 crypto.go:68] Generating cert /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.crt with IP's: []
	I1002 12:43:58.560294  581986 crypto.go:156] Writing cert to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.crt ...
	I1002 12:43:58.560321  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.crt: {Name:mk4b5f44166e00d4bcde8b8ec27b665d690df8bf Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.560489  581986 crypto.go:164] Writing key to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.key ...
	I1002 12:43:58.560505  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.key: {Name:mkb898b47df603bf78782059d8a04e6f09d8784d Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.560627  581986 certs.go:319] generating minikube signed cert: /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key.31bdca25
	I1002 12:43:58.560638  581986 crypto.go:68] Generating cert /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt.31bdca25 with IP's: [192.168.76.2 10.96.0.1 127.0.0.1 10.0.0.1]
	I1002 12:43:58.650268  581986 crypto.go:156] Writing cert to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt.31bdca25 ...
	I1002 12:43:58.650284  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt.31bdca25: {Name:mk433035fa83eca243b87bc5424c9d7571a6132d Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.650416  581986 crypto.go:164] Writing key to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key.31bdca25 ...
	I1002 12:43:58.650425  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key.31bdca25: {Name:mk57c1c8dfb5d302d0dbc40c6e11db8c54cb244e Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.650490  581986 certs.go:337] copying /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt.31bdca25 -> /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt
	I1002 12:43:58.650569  581986 certs.go:341] copying /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key.31bdca25 -> /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key
	I1002 12:43:58.650622  581986 certs.go:319] generating aggregator signed cert: /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.key
	I1002 12:43:58.650636  581986 crypto.go:68] Generating cert /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.crt with IP's: []
	I1002 12:43:58.956632  581986 crypto.go:156] Writing cert to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.crt ...
	I1002 12:43:58.956649  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.crt: {Name:mk2e1384cf3a50706208d42e4648a292f00d76c3 Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.956803  581986 crypto.go:164] Writing key to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.key ...
	I1002 12:43:58.956814  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.key: {Name:mka6a0a0cf42a7c2aae8a0a46864b4a18f560a3f Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.957036  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/384955.pem (1338 bytes)
	W1002 12:43:58.957070  581986 certs.go:433] ignoring /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/384955_empty.pem, impossibly tiny 0 bytes
	I1002 12:43:58.957078  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem (1675 bytes)
	I1002 12:43:58.957104  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem (1082 bytes)
	I1002 12:43:58.957124  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem (1123 bytes)
	I1002 12:43:58.957144  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem (1675 bytes)
	I1002 12:43:58.957179  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem (1708 bytes)
	I1002 12:43:58.957893  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt --> /var/lib/minikube/certs/apiserver.crt (1399 bytes)
	I1002 12:43:58.987145  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key --> /var/lib/minikube/certs/apiserver.key (1679 bytes)
	I1002 12:43:59.011640  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.crt --> /var/lib/minikube/certs/proxy-client.crt (1147 bytes)
	I1002 12:43:59.043029  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.key --> /var/lib/minikube/certs/proxy-client.key (1679 bytes)
	I1002 12:43:59.071607  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/ca.crt --> /var/lib/minikube/certs/ca.crt (1111 bytes)
	I1002 12:43:59.095951  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/ca.key --> /var/lib/minikube/certs/ca.key (1679 bytes)
	I1002 12:43:59.120454  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/proxy-client-ca.crt --> /var/lib/minikube/certs/proxy-client-ca.crt (1119 bytes)
	I1002 12:43:59.144500  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/proxy-client-ca.key --> /var/lib/minikube/certs/proxy-client-ca.key (1679 bytes)
	I1002 12:43:59.171360  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem --> /usr/share/ca-certificates/3849552.pem (1708 bytes)
	I1002 12:43:59.201181  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/ca.crt --> /usr/share/ca-certificates/minikubeCA.pem (1111 bytes)
	I1002 12:43:59.228983  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/certs/384955.pem --> /usr/share/ca-certificates/384955.pem (1338 bytes)
	I1002 12:43:59.255173  581986 ssh_runner.go:362] scp memory --> /var/lib/minikube/kubeconfig (738 bytes)
	I1002 12:43:59.274224  581986 ssh_runner.go:195] Run: openssl version
	I1002 12:43:59.281002  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -s /usr/share/ca-certificates/384955.pem && ln -fs /usr/share/ca-certificates/384955.pem /etc/ssl/certs/384955.pem"
	I1002 12:43:59.291291  581986 ssh_runner.go:195] Run: ls -la /usr/share/ca-certificates/384955.pem
	I1002 12:43:59.294937  581986 certs.go:480] hashing: -rw-r--r-- 1 root root 1338 Oct  2 12:20 /usr/share/ca-certificates/384955.pem
	I1002 12:43:59.294983  581986 ssh_runner.go:195] Run: openssl x509 -hash -noout -in /usr/share/ca-certificates/384955.pem
	I1002 12:43:59.301811  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -L /etc/ssl/certs/51391683.0 || ln -fs /etc/ssl/certs/384955.pem /etc/ssl/certs/51391683.0"
	I1002 12:43:59.315896  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -s /usr/share/ca-certificates/3849552.pem && ln -fs /usr/share/ca-certificates/3849552.pem /etc/ssl/certs/3849552.pem"
	I1002 12:43:59.326590  581986 ssh_runner.go:195] Run: ls -la /usr/share/ca-certificates/3849552.pem
	I1002 12:43:59.330767  581986 certs.go:480] hashing: -rw-r--r-- 1 root root 1708 Oct  2 12:20 /usr/share/ca-certificates/3849552.pem
	I1002 12:43:59.330814  581986 ssh_runner.go:195] Run: openssl x509 -hash -noout -in /usr/share/ca-certificates/3849552.pem
	I1002 12:43:59.339545  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -L /etc/ssl/certs/3ec20f2e.0 || ln -fs /etc/ssl/certs/3849552.pem /etc/ssl/certs/3ec20f2e.0"
	I1002 12:43:59.350676  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -s /usr/share/ca-certificates/minikubeCA.pem && ln -fs /usr/share/ca-certificates/minikubeCA.pem /etc/ssl/certs/minikubeCA.pem"
	I1002 12:43:59.361202  581986 ssh_runner.go:195] Run: ls -la /usr/share/ca-certificates/minikubeCA.pem
	I1002 12:43:59.365896  581986 certs.go:480] hashing: -rw-r--r-- 1 root root 1111 Oct  2 12:15 /usr/share/ca-certificates/minikubeCA.pem
	I1002 12:43:59.365958  581986 ssh_runner.go:195] Run: openssl x509 -hash -noout -in /usr/share/ca-certificates/minikubeCA.pem
	I1002 12:43:59.375584  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -L /etc/ssl/certs/b5213941.0 || ln -fs /etc/ssl/certs/minikubeCA.pem /etc/ssl/certs/b5213941.0"
	I1002 12:43:59.388988  581986 ssh_runner.go:195] Run: ls /var/lib/minikube/certs/etcd
	I1002 12:43:59.393371  581986 certs.go:353] certs directory doesn't exist, likely first start: ls /var/lib/minikube/certs/etcd: Process exited with status 2
	stdout:
	
	stderr:
	ls: cannot access '/var/lib/minikube/certs/etcd': No such file or directory
	I1002 12:43:59.393422  581986 kubeadm.go:404] StartCluster: {Name:missing-upgrade-404552 KeepContext:false EmbedCerts:false MinikubeISO: KicBaseImage:gcr.io/k8s-minikube/kicbase:v0.0.42@sha256:d35ac07dfda971cabee05e0deca8aeac772f885a5348e1a0c0b0a36db20fcfc0 Memory:3072 CPUs:2 DiskSize:20000 VMDriver: Driver:docker HyperkitVpnKitSock: HyperkitVSockPorts:[] DockerEnv:[] ContainerVolumeMounts:[] InsecureRegistry:[] RegistryMirror:[] HostOnlyCIDR:192.168.59.1/24 HypervVirtualSwitch: HypervUseExternalSwitch:false HypervExternalAdapter: KVMNetwork:default KVMQemuURI:qemu:///system KVMGPU:false KVMHidden:false KVMNUMACount:1 APIServerPort:0 DockerOpt:[] DisableDriverMounts:false NFSShare:[] NFSSharesRoot:/nfsshares UUID: NoVTXCheck:false DNSProxy:false HostDNSResolver:true HostOnlyNicType:virtio NatNicType:virtio SSHIPAddress: SSHUser:root SSHKey: SSHPort:22 KubernetesConfig:{KubernetesVersion:v1.28.3 ClusterName:missing-upgrade-404552 Namespace:default APIServerName:minikubeCA APIServerNames:[] APIServerIP
s:[] DNSDomain:cluster.local ContainerRuntime:containerd CRISocket: NetworkPlugin:cni FeatureGates: ServiceCIDR:10.96.0.0/12 ImageRepository: LoadBalancerStartIP: LoadBalancerEndIP: CustomIngressCert: RegistryAliases: ExtraOptions:[] ShouldLoadCachedImages:true EnableDefaultCNI:false CNI: NodeIP: NodePort:8443 NodeName:} Nodes:[{Name: IP:192.168.76.2 Port:8443 KubernetesVersion:v1.28.3 ContainerRuntime:containerd ControlPlane:true Worker:true}] Addons:map[] CustomAddonImages:map[] CustomAddonRegistries:map[] VerifyComponents:map[apiserver:true system_pods:true] StartHostTimeout:6m0s ScheduledStop:<nil> ExposedPorts:[] ListenAddress: Network: Subnet: MultiNodeRequested:false ExtraDisks:0 CertExpiration:26280h0m0s Mount:false MountString:/home/jenkins:/minikube-host Mount9PVersion:9p2000.L MountGID:docker MountIP: MountMSize:262144 MountOptions:[] MountPort:0 MountType:9p MountUID:docker BinaryMirror: DisableOptimizations:false DisableMetrics:false CustomQemuFirmwarePath: SocketVMnetClientPath: SocketVMnetPath:
StaticIP: SSHAuthSock: SSHAgentPID:0 AutoPauseInterval:1m0s GPUs:}
	I1002 12:43:59.393503  581986 cri.go:54] listing CRI containers in root /run/containerd/runc/k8s.io: {State:paused Name: Namespaces:[kube-system]}
	I1002 12:43:59.393542  581986 ssh_runner.go:195] Run: sudo -s eval "crictl ps -a --quiet --label io.kubernetes.pod.namespace=kube-system"
	I1002 12:43:59.451021  581986 cri.go:89] found id: ""
	I1002 12:43:59.451105  581986 ssh_runner.go:195] Run: sudo ls /var/lib/kubelet/kubeadm-flags.env /var/lib/kubelet/config.yaml /var/lib/minikube/etcd
	I1002 12:43:59.465493  581986 ssh_runner.go:195] Run: sudo cp /var/tmp/minikube/kubeadm.yaml.new /var/tmp/minikube/kubeadm.yaml
	I1002 12:43:59.476153  581986 kubeadm.go:226] ignoring SystemVerification for kubeadm because of docker driver
	I1002 12:43:59.476209  581986 ssh_runner.go:195] Run: sudo ls -la /etc/kubernetes/admin.conf /etc/kubernetes/kubelet.conf /etc/kubernetes/controller-manager.conf /etc/kubernetes/scheduler.conf
	I1002 12:43:59.486438  581986 kubeadm.go:152] config check failed, skipping stale config cleanup: sudo ls -la /etc/kubernetes/admin.conf /etc/kubernetes/kubelet.conf /etc/kubernetes/controller-manager.conf /etc/kubernetes/scheduler.conf: Process exited with status 2
	stdout:
	
	stderr:
	ls: cannot access '/etc/kubernetes/admin.conf': No such file or directory
	ls: cannot access '/etc/kubernetes/kubelet.conf': No such file or directory
	ls: cannot access '/etc/kubernetes/controller-manager.conf': No such file or directory
	ls: cannot access '/etc/kubernetes/scheduler.conf': No such file or directory
	I1002 12:43:59.486490  581986 ssh_runner.go:286] Start: /bin/bash -c "sudo env PATH="/var/lib/minikube/binaries/v1.28.3:$PATH" kubeadm init --config /var/tmp/minikube/kubeadm.yaml  --ignore-preflight-errors=DirAvailable--etc-kubernetes-manifests,DirAvailable--var-lib-minikube,DirAvailable--var-lib-minikube-etcd,FileAvailable--etc-kubernetes-manifests-kube-scheduler.yaml,FileAvailable--etc-kubernetes-manifests-kube-apiserver.yaml,FileAvailable--etc-kubernetes-manifests-kube-controller-manager.yaml,FileAvailable--etc-kubernetes-manifests-etcd.yaml,Port-10250,Swap,NumCPU,Mem,SystemVerification,FileContent--proc-sys-net-bridge-bridge-nf-call-iptables"
	I1002 12:43:59.551520  581986 kubeadm.go:322] [init] Using Kubernetes version: v1.28.3
	I1002 12:43:59.551643  581986 kubeadm.go:322] [preflight] Running pre-flight checks
	I1002 12:43:59.591707  581986 kubeadm.go:322] [preflight] The system verification failed. Printing the output from the verification:
	I1002 12:43:59.591787  581986 kubeadm.go:322] [0;37mKERNEL_VERSION[0m: [0;32m6.8.0-1041-gcp[0m
	I1002 12:43:59.591870  581986 kubeadm.go:322] [0;37mOS[0m: [0;32mLinux[0m
	I1002 12:43:59.591928  581986 kubeadm.go:322] [0;37mCGROUPS_CPU[0m: [0;32menabled[0m
	I1002 12:43:59.591995  581986 kubeadm.go:322] [0;37mCGROUPS_CPUSET[0m: [0;32menabled[0m
	I1002 12:43:59.592058  581986 kubeadm.go:322] [0;37mCGROUPS_DEVICES[0m: [0;32menabled[0m
	I1002 12:43:59.592121  581986 kubeadm.go:322] [0;37mCGROUPS_FREEZER[0m: [0;32menabled[0m
	I1002 12:43:59.592183  581986 kubeadm.go:322] [0;37mCGROUPS_MEMORY[0m: [0;32menabled[0m
	I1002 12:43:59.592259  581986 kubeadm.go:322] [0;37mCGROUPS_PIDS[0m: [0;32menabled[0m
	I1002 12:43:59.592332  581986 kubeadm.go:322] [0;37mCGROUPS_HUGETLB[0m: [0;32menabled[0m
	I1002 12:43:59.592396  581986 kubeadm.go:322] [0;37mCGROUPS_IO[0m: [0;32menabled[0m
	I1002 12:43:59.679277  581986 kubeadm.go:322] [preflight] Pulling images required for setting up a Kubernetes cluster
	I1002 12:43:59.679443  581986 kubeadm.go:322] [preflight] This might take a minute or two, depending on the speed of your internet connection
	I1002 12:43:59.679569  581986 kubeadm.go:322] [preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
	I1002 12:43:59.903212  581986 kubeadm.go:322] [certs] Using certificateDir folder "/var/lib/minikube/certs"
	I1002 12:43:59.136103  583227 provision.go:177] copyRemoteCerts
	I1002 12:43:59.136189  583227 ssh_runner.go:195] Run: sudo mkdir -p /etc/docker /etc/docker /etc/docker
	I1002 12:43:59.136246  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.156567  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	I1002 12:43:59.246326  583227 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /etc/docker/ca.pem (1082 bytes)
	I1002 12:43:59.274495  583227 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem --> /etc/docker/server.pem (1233 bytes)
	I1002 12:43:59.300439  583227 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem --> /etc/docker/server-key.pem (1679 bytes)
	I1002 12:43:59.330525  583227 provision.go:87] duration metric: took 506.064105ms to configureAuth
	I1002 12:43:59.330574  583227 ubuntu.go:206] setting minikube options for container-runtime
	I1002 12:43:59.330764  583227 config.go:182] Loaded profile config "stopped-upgrade-250350": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:59.330785  583227 machine.go:96] duration metric: took 3.943503201s to provisionDockerMachine
	I1002 12:43:59.330798  583227 start.go:293] postStartSetup for "stopped-upgrade-250350" (driver="docker")
	I1002 12:43:59.330812  583227 start.go:322] creating required directories: [/etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs]
	I1002 12:43:59.330941  583227 ssh_runner.go:195] Run: sudo mkdir -p /etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs
	I1002 12:43:59.331092  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.351859  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	I1002 12:43:59.448934  583227 ssh_runner.go:195] Run: cat /etc/os-release
	I1002 12:43:59.456107  583227 main.go:141] libmachine: Couldn't set key VERSION_CODENAME, no corresponding struct field found
	I1002 12:43:59.456147  583227 main.go:141] libmachine: Couldn't set key PRIVACY_POLICY_URL, no corresponding struct field found
	I1002 12:43:59.456232  583227 main.go:141] libmachine: Couldn't set key UBUNTU_CODENAME, no corresponding struct field found
	I1002 12:43:59.456282  583227 info.go:137] Remote host: Ubuntu 22.04.3 LTS
	I1002 12:43:59.456298  583227 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/addons for local assets ...
	I1002 12:43:59.456371  583227 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/files for local assets ...
	I1002 12:43:59.456465  583227 filesync.go:149] local asset: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem -> 3849552.pem in /etc/ssl/certs
	I1002 12:43:59.456602  583227 ssh_runner.go:195] Run: sudo mkdir -p /etc/ssl/certs
	I1002 12:43:59.469387  583227 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem --> /etc/ssl/certs/3849552.pem (1708 bytes)
	I1002 12:43:59.498329  583227 start.go:296] duration metric: took 167.514063ms for postStartSetup
	I1002 12:43:59.498398  583227 ssh_runner.go:195] Run: sh -c "df -h /var | awk 'NR==2{print $5}'"
	I1002 12:43:59.498445  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.524891  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	I1002 12:43:59.614708  583227 ssh_runner.go:195] Run: sh -c "df -BG /var | awk 'NR==2{print $4}'"
	I1002 12:43:59.620747  583227 fix.go:56] duration metric: took 4.606941327s for fixHost
	I1002 12:43:59.620768  583227 start.go:83] releasing machines lock for "stopped-upgrade-250350", held for 4.606988432s
	I1002 12:43:59.620860  583227 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" stopped-upgrade-250350
	I1002 12:43:59.645041  583227 ssh_runner.go:195] Run: cat /version.json
	I1002 12:43:59.645101  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.645135  583227 ssh_runner.go:195] Run: curl -sS -m 2 https://registry.k8s.io/
	I1002 12:43:59.645202  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.668811  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	I1002 12:43:59.668811  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	W1002 12:43:59.878768  583227 out.go:285] ! Image was not built for the current minikube version. To resolve this you can delete and recreate your minikube cluster using the latest images. Expected minikube version: v1.32.0 -> Actual minikube version: v1.37.0
	I1002 12:43:59.878919  583227 ssh_runner.go:195] Run: systemctl --version
	I1002 12:43:59.883939  583227 ssh_runner.go:195] Run: sh -c "stat /etc/cni/net.d/*loopback.conf*"
	I1002 12:43:59.888632  583227 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f -name *loopback.conf* -not -name *.mk_disabled -exec sh -c "grep -q loopback {} && ( grep -q name {} || sudo sed -i '/"type": "loopback"/i \ \ \ \ "name": "loopback",' {} ) && sudo sed -i 's|"cniVersion": ".*"|"cniVersion": "1.0.0"|g' {}" ;
	I1002 12:43:59.909328  583227 cni.go:230] loopback cni configuration patched: "/etc/cni/net.d/*loopback.conf*" found
	I1002 12:43:59.909394  583227 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f ( ( -name *bridge* -or -name *podman* ) -and -not -name *.mk_disabled ) -printf "%p, " -exec sh -c "sudo mv {} {}.mk_disabled" ;
	I1002 12:43:59.918213  583227 cni.go:259] no active bridge cni configs found in "/etc/cni/net.d" - nothing to disable
	I1002 12:43:59.918236  583227 start.go:495] detecting cgroup driver to use...
	I1002 12:43:59.918270  583227 detect.go:190] detected "systemd" cgroup driver on host os
	I1002 12:43:59.918309  583227 ssh_runner.go:195] Run: sudo systemctl stop -f crio
	I1002 12:43:59.931237  583227 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service crio
	I1002 12:43:59.942383  583227 docker.go:218] disabling cri-docker service (if available) ...
	I1002 12:43:59.942437  583227 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.socket
	I1002 12:43:59.954435  583227 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.service
	I1002 12:43:59.965821  583227 ssh_runner.go:195] Run: sudo systemctl disable cri-docker.socket
	I1002 12:44:00.032939  583227 ssh_runner.go:195] Run: sudo systemctl mask cri-docker.service
	I1002 12:44:00.097207  583227 docker.go:234] disabling docker service ...
	I1002 12:44:00.097264  583227 ssh_runner.go:195] Run: sudo systemctl stop -f docker.socket
	I1002 12:44:00.109382  583227 ssh_runner.go:195] Run: sudo systemctl stop -f docker.service
	I1002 12:44:00.120575  583227 ssh_runner.go:195] Run: sudo systemctl disable docker.socket
	I1002 12:44:00.187758  583227 ssh_runner.go:195] Run: sudo systemctl mask docker.service
	I1002 12:44:00.253215  583227 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service docker
	I1002 12:44:00.264420  583227 ssh_runner.go:195] Run: /bin/bash -c "sudo mkdir -p /etc && printf %s "runtime-endpoint: unix:///run/containerd/containerd.sock
	" | sudo tee /etc/crictl.yaml"
	I1002 12:44:00.282072  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)sandbox_image = .*$|\1sandbox_image = "registry.k8s.io/pause:3.9"|' /etc/containerd/config.toml"
	I1002 12:44:00.291907  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)restrict_oom_score_adj = .*$|\1restrict_oom_score_adj = false|' /etc/containerd/config.toml"
	I1002 12:44:00.301493  583227 containerd.go:146] configuring containerd to use "systemd" as cgroup driver...
	I1002 12:44:00.301553  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)SystemdCgroup = .*$|\1SystemdCgroup = true|g' /etc/containerd/config.toml"
	I1002 12:44:00.311281  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runtime.v1.linux"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:44:00.321045  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i '/systemd_cgroup/d' /etc/containerd/config.toml"
	I1002 12:44:00.330693  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runc.v1"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:44:00.340571  583227 ssh_runner.go:195] Run: sh -c "sudo rm -rf /etc/cni/net.mk"
	I1002 12:44:00.349706  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)conf_dir = .*$|\1conf_dir = "/etc/cni/net.d"|g' /etc/containerd/config.toml"
	I1002 12:44:00.359285  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i '/^ *enable_unprivileged_ports = .*/d' /etc/containerd/config.toml"
	I1002 12:44:00.369202  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)\[plugins."io.containerd.grpc.v1.cri"\]|&\n\1  enable_unprivileged_ports = true|' /etc/containerd/config.toml"
	I1002 12:44:00.379878  583227 ssh_runner.go:195] Run: sudo sysctl net.bridge.bridge-nf-call-iptables
	I1002 12:44:00.388529  583227 ssh_runner.go:195] Run: sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
	I1002 12:44:00.396849  583227 ssh_runner.go:195] Run: sudo systemctl daemon-reload
	I1002 12:44:00.464051  583227 ssh_runner.go:195] Run: sudo systemctl restart containerd
	I1002 12:44:00.576008  583227 start.go:542] Will wait 60s for socket path /run/containerd/containerd.sock
	I1002 12:44:00.576088  583227 ssh_runner.go:195] Run: stat /run/containerd/containerd.sock
	I1002 12:44:00.580326  583227 start.go:563] Will wait 60s for crictl version
	I1002 12:44:00.580379  583227 ssh_runner.go:195] Run: which crictl
	I1002 12:44:00.583808  583227 ssh_runner.go:195] Run: sudo /usr/bin/crictl version
	I1002 12:44:00.618127  583227 start.go:579] Version:  0.1.0
	RuntimeName:  containerd
	RuntimeVersion:  1.6.24
	RuntimeApiVersion:  v1
	I1002 12:44:00.618181  583227 ssh_runner.go:195] Run: containerd --version
	I1002 12:44:00.642732  583227 ssh_runner.go:195] Run: containerd --version
	I1002 12:44:00.670321  583227 out.go:179] * Preparing Kubernetes v1.28.3 on containerd 1.6.24 ...
	I1002 12:43:59.904975  581986 out.go:204]   - Generating certificates and keys ...
	I1002 12:43:59.905096  581986 kubeadm.go:322] [certs] Using existing ca certificate authority
	I1002 12:43:59.905187  581986 kubeadm.go:322] [certs] Using existing apiserver certificate and key on disk
	I1002 12:44:00.230450  581986 kubeadm.go:322] [certs] Generating "apiserver-kubelet-client" certificate and key
	I1002 12:44:00.342809  581986 kubeadm.go:322] [certs] Generating "front-proxy-ca" certificate and key
	I1002 12:44:00.427885  581986 kubeadm.go:322] [certs] Generating "front-proxy-client" certificate and key
	I1002 12:44:00.668257  581986 kubeadm.go:322] [certs] Generating "etcd/ca" certificate and key
	I1002 12:44:00.911902  581986 kubeadm.go:322] [certs] Generating "etcd/server" certificate and key
	I1002 12:44:00.912096  581986 kubeadm.go:322] [certs] etcd/server serving cert is signed for DNS names [localhost missing-upgrade-404552] and IPs [192.168.76.2 127.0.0.1 ::1]
	I1002 12:43:58.333473  585741 out.go:252] * Creating docker container (CPUs=2, Memory=3072MB) ...
	I1002 12:43:58.333729  585741 start.go:159] libmachine.API.Create for "NoKubernetes-115222" (driver="docker")
	I1002 12:43:58.333767  585741 client.go:168] LocalClient.Create starting
	I1002 12:43:58.333864  585741 main.go:141] libmachine: Reading certificate data from /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem
	I1002 12:43:58.333920  585741 main.go:141] libmachine: Decoding PEM data...
	I1002 12:43:58.333951  585741 main.go:141] libmachine: Parsing certificate...
	I1002 12:43:58.334022  585741 main.go:141] libmachine: Reading certificate data from /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem
	I1002 12:43:58.334051  585741 main.go:141] libmachine: Decoding PEM data...
	I1002 12:43:58.334068  585741 main.go:141] libmachine: Parsing certificate...
	I1002 12:43:58.334563  585741 cli_runner.go:164] Run: docker network inspect NoKubernetes-115222 --format "{"Name": "{{.Name}}","Driver": "{{.Driver}}","Subnet": "{{range .IPAM.Config}}{{.Subnet}}{{end}}","Gateway": "{{range .IPAM.Config}}{{.Gateway}}{{end}}","MTU": {{if (index .Options "com.docker.network.driver.mtu")}}{{(index .Options "com.docker.network.driver.mtu")}}{{else}}0{{end}}, "ContainerIPs": [{{range $k,$v := .Containers }}"{{$v.IPv4Address}}",{{end}}]}"
	W1002 12:43:58.352366  585741 cli_runner.go:211] docker network inspect NoKubernetes-115222 --format "{"Name": "{{.Name}}","Driver": "{{.Driver}}","Subnet": "{{range .IPAM.Config}}{{.Subnet}}{{end}}","Gateway": "{{range .IPAM.Config}}{{.Gateway}}{{end}}","MTU": {{if (index .Options "com.docker.network.driver.mtu")}}{{(index .Options "com.docker.network.driver.mtu")}}{{else}}0{{end}}, "ContainerIPs": [{{range $k,$v := .Containers }}"{{$v.IPv4Address}}",{{end}}]}" returned with exit code 1
	I1002 12:43:58.352435  585741 network_create.go:284] running [docker network inspect NoKubernetes-115222] to gather additional debugging logs...
	I1002 12:43:58.352469  585741 cli_runner.go:164] Run: docker network inspect NoKubernetes-115222
	W1002 12:43:58.368779  585741 cli_runner.go:211] docker network inspect NoKubernetes-115222 returned with exit code 1
	I1002 12:43:58.368806  585741 network_create.go:287] error running [docker network inspect NoKubernetes-115222]: docker network inspect NoKubernetes-115222: exit status 1
	stdout:
	[]
	
	stderr:
	Error response from daemon: network NoKubernetes-115222 not found
	I1002 12:43:58.368821  585741 network_create.go:289] output of [docker network inspect NoKubernetes-115222]: -- stdout --
	[]
	
	-- /stdout --
	** stderr ** 
	Error response from daemon: network NoKubernetes-115222 not found
	
	** /stderr **
	I1002 12:43:58.368926  585741 cli_runner.go:164] Run: docker network inspect bridge --format "{"Name": "{{.Name}}","Driver": "{{.Driver}}","Subnet": "{{range .IPAM.Config}}{{.Subnet}}{{end}}","Gateway": "{{range .IPAM.Config}}{{.Gateway}}{{end}}","MTU": {{if (index .Options "com.docker.network.driver.mtu")}}{{(index .Options "com.docker.network.driver.mtu")}}{{else}}0{{end}}, "ContainerIPs": [{{range $k,$v := .Containers }}"{{$v.IPv4Address}}",{{end}}]}"
	I1002 12:43:58.386690  585741 network.go:211] skipping subnet 192.168.49.0/24 that is taken: &{IP:192.168.49.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.49.0/24 Gateway:192.168.49.1 ClientMin:192.168.49.2 ClientMax:192.168.49.254 Broadcast:192.168.49.255 IsPrivate:true Interface:{IfaceName:br-c8fe0fd21bd1 IfaceIPv4:192.168.49.1 IfaceMTU:1500 IfaceMAC:46:4d:f9:f4:cc:90} reservation:<nil>}
	I1002 12:43:58.387729  585741 network.go:211] skipping subnet 192.168.58.0/24 that is taken: &{IP:192.168.58.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.58.0/24 Gateway:192.168.58.1 ClientMin:192.168.58.2 ClientMax:192.168.58.254 Broadcast:192.168.58.255 IsPrivate:true Interface:{IfaceName:br-5cdc9a6b29c5 IfaceIPv4:192.168.58.1 IfaceMTU:1500 IfaceMAC:62:2f:0d:cf:b5:61} reservation:<nil>}
	I1002 12:43:58.388265  585741 network.go:211] skipping subnet 192.168.67.0/24 that is taken: &{IP:192.168.67.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.67.0/24 Gateway:192.168.67.1 ClientMin:192.168.67.2 ClientMax:192.168.67.254 Broadcast:192.168.67.255 IsPrivate:true Interface:{IfaceName:br-cb914d436a66 IfaceIPv4:192.168.67.1 IfaceMTU:1500 IfaceMAC:9a:69:ee:9d:31:3d} reservation:<nil>}
	I1002 12:43:58.388874  585741 network.go:211] skipping subnet 192.168.76.0/24 that is taken: &{IP:192.168.76.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.76.0/24 Gateway:192.168.76.1 ClientMin:192.168.76.2 ClientMax:192.168.76.254 Broadcast:192.168.76.255 IsPrivate:true Interface:{IfaceName:br-f4a55d02f8d3 IfaceIPv4:192.168.76.1 IfaceMTU:1500 IfaceMAC:9a:b9:5b:5d:66:d3} reservation:<nil>}
	I1002 12:43:58.389804  585741 network.go:206] using free private subnet 192.168.85.0/24: &{IP:192.168.85.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.85.0/24 Gateway:192.168.85.1 ClientMin:192.168.85.2 ClientMax:192.168.85.254 Broadcast:192.168.85.255 IsPrivate:true Interface:{IfaceName: IfaceIPv4: IfaceMTU:0 IfaceMAC:} reservation:0xc001e9dda0}
	I1002 12:43:58.389845  585741 network_create.go:124] attempt to create docker network NoKubernetes-115222 192.168.85.0/24 with gateway 192.168.85.1 and MTU of 1500 ...
	I1002 12:43:58.389896  585741 cli_runner.go:164] Run: docker network create --driver=bridge --subnet=192.168.85.0/24 --gateway=192.168.85.1 -o --ip-masq -o --icc -o com.docker.network.driver.mtu=1500 --label=created_by.minikube.sigs.k8s.io=true --label=name.minikube.sigs.k8s.io=NoKubernetes-115222 NoKubernetes-115222
	I1002 12:43:58.449771  585741 network_create.go:108] docker network NoKubernetes-115222 192.168.85.0/24 created
	I1002 12:43:58.449818  585741 kic.go:121] calculated static IP "192.168.85.2" for the "NoKubernetes-115222" container
	I1002 12:43:58.449907  585741 cli_runner.go:164] Run: docker ps -a --format {{.Names}}
	I1002 12:43:58.470236  585741 cli_runner.go:164] Run: docker volume create NoKubernetes-115222 --label name.minikube.sigs.k8s.io=NoKubernetes-115222 --label created_by.minikube.sigs.k8s.io=true
	I1002 12:43:58.487719  585741 oci.go:103] Successfully created a docker volume NoKubernetes-115222
	I1002 12:43:58.487792  585741 cli_runner.go:164] Run: docker run --rm --name NoKubernetes-115222-preload-sidecar --label created_by.minikube.sigs.k8s.io=true --label name.minikube.sigs.k8s.io=NoKubernetes-115222 --entrypoint /usr/bin/test -v NoKubernetes-115222:/var gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 -d /var/lib
	I1002 12:43:58.855992  585741 oci.go:107] Successfully prepared a docker volume NoKubernetes-115222
	I1002 12:43:58.856028  585741 preload.go:178] Skipping preload logic due to --no-kubernetes flag
	W1002 12:43:58.856092  585741 cgroups_linux.go:77] Your kernel does not support swap limit capabilities or the cgroup is not mounted.
	W1002 12:43:58.856120  585741 oci.go:252] Your kernel does not support CPU cfs period/quota or the cgroup is not mounted.
	I1002 12:43:58.856155  585741 cli_runner.go:164] Run: docker info --format "'{{json .SecurityOptions}}'"
	I1002 12:43:58.914050  585741 cli_runner.go:164] Run: docker run -d -t --privileged --security-opt seccomp=unconfined --tmpfs /tmp --tmpfs /run -v /lib/modules:/lib/modules:ro --hostname NoKubernetes-115222 --name NoKubernetes-115222 --label created_by.minikube.sigs.k8s.io=true --label name.minikube.sigs.k8s.io=NoKubernetes-115222 --label role.minikube.sigs.k8s.io= --label mode.minikube.sigs.k8s.io=NoKubernetes-115222 --network NoKubernetes-115222 --ip 192.168.85.2 --volume NoKubernetes-115222:/var --security-opt apparmor=unconfined --memory=3072mb -e container=docker --expose 8443 --publish=127.0.0.1::8443 --publish=127.0.0.1::22 --publish=127.0.0.1::2376 --publish=127.0.0.1::5000 --publish=127.0.0.1::32443 gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1
	I1002 12:43:59.170328  585741 cli_runner.go:164] Run: docker container inspect NoKubernetes-115222 --format={{.State.Running}}
	I1002 12:43:59.190310  585741 cli_runner.go:164] Run: docker container inspect NoKubernetes-115222 --format={{.State.Status}}
	I1002 12:43:59.210643  585741 cli_runner.go:164] Run: docker exec NoKubernetes-115222 stat /var/lib/dpkg/alternatives/iptables
	I1002 12:43:59.256165  585741 oci.go:144] the created container "NoKubernetes-115222" has a running status.
	I1002 12:43:59.256202  585741 kic.go:225] Creating ssh key for kic: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa...
	I1002 12:43:59.313979  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa.pub -> /home/docker/.ssh/authorized_keys
	I1002 12:43:59.314036  585741 kic_runner.go:191] docker (temp): /home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa.pub --> /home/docker/.ssh/authorized_keys (381 bytes)
	I1002 12:43:59.343661  585741 cli_runner.go:164] Run: docker container inspect NoKubernetes-115222 --format={{.State.Status}}
	I1002 12:43:59.364124  585741 kic_runner.go:93] Run: chown docker:docker /home/docker/.ssh/authorized_keys
	I1002 12:43:59.364145  585741 kic_runner.go:114] Args: [docker exec --privileged NoKubernetes-115222 chown docker:docker /home/docker/.ssh/authorized_keys]
	I1002 12:43:59.411233  585741 cli_runner.go:164] Run: docker container inspect NoKubernetes-115222 --format={{.State.Status}}
	I1002 12:43:59.437345  585741 machine.go:93] provisionDockerMachine start ...
	I1002 12:43:59.437476  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:43:59.465028  585741 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:59.465377  585741 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33390 <nil> <nil>}
	I1002 12:43:59.465401  585741 main.go:141] libmachine: About to run SSH command:
	hostname
	I1002 12:43:59.613979  585741 main.go:141] libmachine: SSH cmd err, output: <nil>: NoKubernetes-115222
	
	I1002 12:43:59.614086  585741 ubuntu.go:182] provisioning hostname "NoKubernetes-115222"
	I1002 12:43:59.614195  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:43:59.637770  585741 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:59.638098  585741 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33390 <nil> <nil>}
	I1002 12:43:59.638114  585741 main.go:141] libmachine: About to run SSH command:
	sudo hostname NoKubernetes-115222 && echo "NoKubernetes-115222" | sudo tee /etc/hostname
	I1002 12:43:59.798996  585741 main.go:141] libmachine: SSH cmd err, output: <nil>: NoKubernetes-115222
	
	I1002 12:43:59.799084  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:43:59.820859  585741 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:59.821181  585741 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33390 <nil> <nil>}
	I1002 12:43:59.821206  585741 main.go:141] libmachine: About to run SSH command:
	
			if ! grep -xq '.*\sNoKubernetes-115222' /etc/hosts; then
				if grep -xq '127.0.1.1\s.*' /etc/hosts; then
					sudo sed -i 's/^127.0.1.1\s.*/127.0.1.1 NoKubernetes-115222/g' /etc/hosts;
				else 
					echo '127.0.1.1 NoKubernetes-115222' | sudo tee -a /etc/hosts; 
				fi
			fi
	I1002 12:43:59.959737  585741 main.go:141] libmachine: SSH cmd err, output: <nil>: 
	I1002 12:43:59.959763  585741 ubuntu.go:188] set auth options {CertDir:/home/jenkins/minikube-integration/21139-381342/.minikube CaCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem CaPrivateKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem CaCertRemotePath:/etc/docker/ca.pem ServerCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ServerKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem ClientKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem ServerCertRemotePath:/etc/docker/server.pem ServerKeyRemotePath:/etc/docker/server-key.pem ClientCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem ServerCertSANs:[] StorePath:/home/jenkins/minikube-integration/21139-381342/.minikube}
	I1002 12:43:59.959785  585741 ubuntu.go:190] setting up certificates
	I1002 12:43:59.959797  585741 provision.go:84] configureAuth start
	I1002 12:43:59.959863  585741 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" NoKubernetes-115222
	I1002 12:43:59.977886  585741 provision.go:143] copyHostCerts
	I1002 12:43:59.977920  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem -> /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem
	I1002 12:43:59.977955  585741 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem, removing ...
	I1002 12:43:59.977972  585741 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem
	I1002 12:43:59.978051  585741 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem (1082 bytes)
	I1002 12:43:59.978144  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem -> /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem
	I1002 12:43:59.978171  585741 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem, removing ...
	I1002 12:43:59.978181  585741 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem
	I1002 12:43:59.978225  585741 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem (1123 bytes)
	I1002 12:43:59.978295  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem -> /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem
	I1002 12:43:59.978321  585741 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem, removing ...
	I1002 12:43:59.978330  585741 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem
	I1002 12:43:59.978366  585741 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem (1675 bytes)
	I1002 12:43:59.978453  585741 provision.go:117] generating server cert: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ca-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem private-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem org=jenkins.NoKubernetes-115222 san=[127.0.0.1 192.168.85.2 NoKubernetes-115222 localhost minikube]
	I1002 12:44:00.173620  585741 provision.go:177] copyRemoteCerts
	I1002 12:44:00.173691  585741 ssh_runner.go:195] Run: sudo mkdir -p /etc/docker /etc/docker /etc/docker
	I1002 12:44:00.173738  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.192246  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.288330  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem -> /etc/docker/ca.pem
	I1002 12:44:00.288391  585741 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /etc/docker/ca.pem (1082 bytes)
	I1002 12:44:00.313987  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem -> /etc/docker/server.pem
	I1002 12:44:00.314044  585741 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem --> /etc/docker/server.pem (1224 bytes)
	I1002 12:44:00.338589  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem -> /etc/docker/server-key.pem
	I1002 12:44:00.338650  585741 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem --> /etc/docker/server-key.pem (1675 bytes)
	I1002 12:44:00.363237  585741 provision.go:87] duration metric: took 403.427992ms to configureAuth
	I1002 12:44:00.363265  585741 ubuntu.go:206] setting minikube options for container-runtime
	I1002 12:44:00.363490  585741 config.go:182] Loaded profile config "NoKubernetes-115222": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v0.0.0
	I1002 12:44:00.363508  585741 machine.go:96] duration metric: took 926.124193ms to provisionDockerMachine
	I1002 12:44:00.363522  585741 client.go:171] duration metric: took 2.029739003s to LocalClient.Create
	I1002 12:44:00.363546  585741 start.go:167] duration metric: took 2.029818501s to libmachine.API.Create "NoKubernetes-115222"
	I1002 12:44:00.363559  585741 start.go:293] postStartSetup for "NoKubernetes-115222" (driver="docker")
	I1002 12:44:00.363574  585741 start.go:322] creating required directories: [/etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs]
	I1002 12:44:00.363623  585741 ssh_runner.go:195] Run: sudo mkdir -p /etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs
	I1002 12:44:00.363670  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.382605  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.481623  585741 ssh_runner.go:195] Run: cat /etc/os-release
	I1002 12:44:00.484862  585741 main.go:141] libmachine: Couldn't set key VERSION_CODENAME, no corresponding struct field found
	I1002 12:44:00.484897  585741 main.go:141] libmachine: Couldn't set key PRIVACY_POLICY_URL, no corresponding struct field found
	I1002 12:44:00.484908  585741 main.go:141] libmachine: Couldn't set key UBUNTU_CODENAME, no corresponding struct field found
	I1002 12:44:00.484919  585741 info.go:137] Remote host: Ubuntu 22.04.5 LTS
	I1002 12:44:00.484930  585741 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/addons for local assets ...
	I1002 12:44:00.484995  585741 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/files for local assets ...
	I1002 12:44:00.485089  585741 filesync.go:149] local asset: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem -> 3849552.pem in /etc/ssl/certs
	I1002 12:44:00.485102  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem -> /etc/ssl/certs/3849552.pem
	I1002 12:44:00.485223  585741 ssh_runner.go:195] Run: sudo mkdir -p /etc/ssl/certs
	I1002 12:44:00.493785  585741 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem --> /etc/ssl/certs/3849552.pem (1708 bytes)
	I1002 12:44:00.520356  585741 start.go:296] duration metric: took 156.779587ms for postStartSetup
	I1002 12:44:00.520741  585741 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" NoKubernetes-115222
	I1002 12:44:00.540900  585741 profile.go:143] Saving config to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/NoKubernetes-115222/config.json ...
	I1002 12:44:00.541121  585741 ssh_runner.go:195] Run: sh -c "df -h /var | awk 'NR==2{print $5}'"
	I1002 12:44:00.541161  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.558900  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.653633  585741 ssh_runner.go:195] Run: sh -c "df -BG /var | awk 'NR==2{print $4}'"
	I1002 12:44:00.658171  585741 start.go:128] duration metric: took 2.32621501s to createHost
	I1002 12:44:00.658198  585741 start.go:83] releasing machines lock for "NoKubernetes-115222", held for 2.326350832s
	I1002 12:44:00.658263  585741 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" NoKubernetes-115222
	I1002 12:44:00.676947  585741 ssh_runner.go:195] Run: cat /version.json
	I1002 12:44:00.676960  585741 ssh_runner.go:195] Run: curl -sS -m 2 https://registry.k8s.io/
	I1002 12:44:00.676998  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.677032  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.696896  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.697291  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.875633  585741 ssh_runner.go:195] Run: systemctl --version
	I1002 12:44:00.880342  585741 ssh_runner.go:195] Run: sh -c "stat /etc/cni/net.d/*loopback.conf*"
	I1002 12:44:00.885039  585741 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f -name *loopback.conf* -not -name *.mk_disabled -exec sh -c "grep -q loopback {} && ( grep -q name {} || sudo sed -i '/"type": "loopback"/i \ \ \ \ "name": "loopback",' {} ) && sudo sed -i 's|"cniVersion": ".*"|"cniVersion": "1.0.0"|g' {}" ;
	I1002 12:44:00.913715  585741 cni.go:230] loopback cni configuration patched: "/etc/cni/net.d/*loopback.conf*" found
	I1002 12:44:00.913787  585741 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f ( ( -name *bridge* -or -name *podman* ) -and -not -name *.mk_disabled ) -printf "%p, " -exec sh -c "sudo mv {} {}.mk_disabled" ;
	I1002 12:44:00.943890  585741 cni.go:262] disabled [/etc/cni/net.d/87-podman-bridge.conflist, /etc/cni/net.d/100-crio-bridge.conf] bridge cni config(s)
	I1002 12:44:00.943918  585741 start.go:495] detecting cgroup driver to use...
	I1002 12:44:00.943951  585741 detect.go:190] detected "systemd" cgroup driver on host os
	I1002 12:44:00.943999  585741 ssh_runner.go:195] Run: sudo systemctl stop -f crio
	I1002 12:44:00.956343  585741 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service crio
	I1002 12:44:00.967470  585741 docker.go:218] disabling cri-docker service (if available) ...
	I1002 12:44:00.967519  585741 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.socket
	I1002 12:44:00.980623  585741 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.service
	I1002 12:44:00.994982  585741 ssh_runner.go:195] Run: sudo systemctl disable cri-docker.socket
	I1002 12:44:01.077417  585741 ssh_runner.go:195] Run: sudo systemctl mask cri-docker.service
	I1002 12:44:01.161042  585741 docker.go:234] disabling docker service ...
	I1002 12:44:01.161112  585741 ssh_runner.go:195] Run: sudo systemctl stop -f docker.socket
	I1002 12:44:01.179656  585741 ssh_runner.go:195] Run: sudo systemctl stop -f docker.service
	I1002 12:44:01.191648  585741 ssh_runner.go:195] Run: sudo systemctl disable docker.socket
	I1002 12:44:01.265615  585741 ssh_runner.go:195] Run: sudo systemctl mask docker.service
	I1002 12:44:01.334379  585741 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service docker
	I1002 12:44:01.346062  585741 ssh_runner.go:195] Run: /bin/bash -c "sudo mkdir -p /etc && printf %s "runtime-endpoint: unix:///run/containerd/containerd.sock
	" | sudo tee /etc/crictl.yaml"
	I1002 12:44:01.362809  585741 binary.go:59] Skipping Kubernetes binary download due to --no-kubernetes flag
	I1002 12:44:01.362921  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)sandbox_image = .*$|\1sandbox_image = "registry.k8s.io/pause:3.9"|' /etc/containerd/config.toml"
	I1002 12:44:01.374360  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)restrict_oom_score_adj = .*$|\1restrict_oom_score_adj = false|' /etc/containerd/config.toml"
	I1002 12:44:01.384353  585741 containerd.go:146] configuring containerd to use "systemd" as cgroup driver...
	I1002 12:44:01.384409  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)SystemdCgroup = .*$|\1SystemdCgroup = true|g' /etc/containerd/config.toml"
	I1002 12:44:01.394457  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runtime.v1.linux"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:44:01.404615  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i '/systemd_cgroup/d' /etc/containerd/config.toml"
	I1002 12:44:01.415284  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runc.v1"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:44:01.425249  585741 ssh_runner.go:195] Run: sh -c "sudo rm -rf /etc/cni/net.mk"
	I1002 12:44:01.434805  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)conf_dir = .*$|\1conf_dir = "/etc/cni/net.d"|g' /etc/containerd/config.toml"
	I1002 12:44:01.445128  585741 ssh_runner.go:195] Run: sudo sysctl net.bridge.bridge-nf-call-iptables
	I1002 12:44:01.453731  585741 ssh_runner.go:195] Run: sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
	I1002 12:44:01.462601  585741 ssh_runner.go:195] Run: sudo systemctl daemon-reload
	I1002 12:44:01.560410  585741 ssh_runner.go:195] Run: sudo systemctl restart containerd
	I1002 12:44:01.669135  585741 start.go:542] Will wait 60s for socket path /run/containerd/containerd.sock
	I1002 12:44:01.669211  585741 ssh_runner.go:195] Run: stat /run/containerd/containerd.sock
	I1002 12:44:01.675844  585741 start.go:563] Will wait 60s for crictl version
	I1002 12:44:01.675914  585741 ssh_runner.go:195] Run: which crictl
	I1002 12:44:01.680923  585741 ssh_runner.go:195] Run: sudo /usr/bin/crictl version
	I1002 12:44:01.737546  585741 start.go:579] Version:  0.1.0
	RuntimeName:  containerd
	RuntimeVersion:  1.7.27
	RuntimeApiVersion:  v1
	I1002 12:44:01.737615  585741 ssh_runner.go:195] Run: containerd --version
	I1002 12:44:01.776697  585741 ssh_runner.go:195] Run: containerd --version
	I1002 12:44:01.812218  585741 out.go:179] * Preparing containerd 1.7.27 ...
	I1002 12:44:01.813955  585741 ssh_runner.go:195] Run: rm -f paused
	I1002 12:44:01.820058  585741 out.go:179] * Done! minikube is ready without Kubernetes!
	I1002 12:44:01.823078  585741 out.go:203] ╭──────────────────────────────────────────────────────────╮
	│                                                          │
	│          * Things to try without Kubernetes ...          │
	│                                                          │
	│    - "minikube ssh" to SSH into minikube's node.         │
	│    - "minikube image" to build images without docker.    │
	│                                                          │
	╰──────────────────────────────────────────────────────────╯
	
	
	==> container status <==
	CONTAINER           IMAGE               CREATED             STATE               NAME                ATTEMPT             POD ID              POD
	
	
	==> containerd <==
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663216733Z" level=info msg="loading plugin \"io.containerd.grpc.v1.version\"..." type=io.containerd.grpc.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663279467Z" level=info msg="loading plugin \"io.containerd.internal.v1.restart\"..." type=io.containerd.internal.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663417335Z" level=info msg="loading plugin \"io.containerd.tracing.processor.v1.otlp\"..." type=io.containerd.tracing.processor.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663498028Z" level=info msg="skip loading plugin \"io.containerd.tracing.processor.v1.otlp\"..." error="skip plugin: tracing endpoint not configured" type=io.containerd.tracing.processor.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663558306Z" level=info msg="loading plugin \"io.containerd.internal.v1.tracing\"..." type=io.containerd.internal.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663654570Z" level=info msg="skip loading plugin \"io.containerd.internal.v1.tracing\"..." error="skip plugin: tracing endpoint not configured" type=io.containerd.internal.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663715445Z" level=info msg="loading plugin \"io.containerd.grpc.v1.healthcheck\"..." type=io.containerd.grpc.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663812713Z" level=info msg="loading plugin \"io.containerd.nri.v1.nri\"..." type=io.containerd.nri.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663908312Z" level=info msg="NRI interface is disabled by configuration."
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663975835Z" level=info msg="loading plugin \"io.containerd.grpc.v1.cri\"..." type=io.containerd.grpc.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.664550919Z" level=info msg="Start cri plugin with config {PluginConfig:{ContainerdConfig:{Snapshotter:overlayfs DefaultRuntimeName:runc DefaultRuntime:{Type: Path: Engine: PodAnnotations:[] ContainerAnnotations:[] Root: Options:map[] PrivilegedWithoutHostDevices:false PrivilegedWithoutHostDevicesAllDevicesAllowed:false BaseRuntimeSpec: NetworkPluginConfDir: NetworkPluginMaxConfNum:0 Snapshotter: SandboxMode:} UntrustedWorkloadRuntime:{Type: Path: Engine: PodAnnotations:[] ContainerAnnotations:[] Root: Options:map[] PrivilegedWithoutHostDevices:false PrivilegedWithoutHostDevicesAllDevicesAllowed:false BaseRuntimeSpec: NetworkPluginConfDir: NetworkPluginMaxConfNum:0 Snapshotter: SandboxMode:} Runtimes:map[runc:{Type:io.containerd.runc.v2 Path: Engine: PodAnnotations:[] ContainerAnnotations:[] Root: Options:map[SystemdCgroup:true] PrivilegedWithoutHostDevices:false PrivilegedWithoutHostDevicesAllDevicesAllowed:false BaseRunti
meSpec: NetworkPluginConfDir: NetworkPluginMaxConfNum:0 Snapshotter: SandboxMode:podsandbox}] NoPivot:false DisableSnapshotAnnotations:true DiscardUnpackedLayers:true IgnoreBlockIONotEnabledErrors:false IgnoreRdtNotEnabledErrors:false} CniConfig:{NetworkPluginBinDir:/opt/cni/bin NetworkPluginConfDir:/etc/cni/net.d NetworkPluginMaxConfNum:1 NetworkPluginSetupSerially:false NetworkPluginConfTemplate: IPPreference:} Registry:{ConfigPath:/etc/containerd/certs.d Mirrors:map[] Configs:map[] Auths:map[] Headers:map[]} ImageDecryption:{KeyModel:node} DisableTCPService:true StreamServerAddress: StreamServerPort:10010 StreamIdleTimeout:4h0m0s EnableSelinux:false SelinuxCategoryRange:1024 SandboxImage:registry.k8s.io/pause:3.9 StatsCollectPeriod:10 SystemdCgroup:false EnableTLSStreaming:false X509KeyPairStreaming:{TLSCertFile: TLSKeyFile:} MaxContainerLogLineSize:16384 DisableCgroup:false DisableApparmor:false RestrictOOMScoreAdj:false MaxConcurrentDownloads:3 DisableProcMount:false UnsetSeccompProfile: TolerateMissingH
ugetlbController:true DisableHugetlbController:true DeviceOwnershipFromSecurityContext:false IgnoreImageDefinedVolumes:false NetNSMountsUnderStateDir:false EnableUnprivilegedPorts:false EnableUnprivilegedICMP:false EnableCDI:false CDISpecDirs:[/etc/cdi /var/run/cdi] ImagePullProgressTimeout:5m0s DrainExecSyncIOTimeout:0s ImagePullWithSyncFs:false IgnoreDeprecationWarnings:[]} ContainerdRootDir:/var/lib/containerd ContainerdEndpoint:/run/containerd/containerd.sock RootDir:/var/lib/containerd/io.containerd.grpc.v1.cri StateDir:/run/containerd/io.containerd.grpc.v1.cri}"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.664884973Z" level=info msg="Connect containerd service"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.664996691Z" level=info msg="using legacy CRI server"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.665075782Z" level=info msg="using experimental NRI integration - disable nri plugin to prevent this"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.665279072Z" level=info msg="Get image filesystem path \"/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs\""
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666369327Z" level=info msg="Start subscribing containerd event"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666447979Z" level=info msg="Start recovering state"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666521472Z" level=info msg=serving... address=/run/containerd/containerd.sock.ttrpc
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666620200Z" level=info msg=serving... address=/run/containerd/containerd.sock
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666551140Z" level=info msg="Start event monitor"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666675786Z" level=info msg="Start snapshots syncer"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666688102Z" level=info msg="Start cni network conf syncer for default"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666700719Z" level=info msg="Start streaming server"
	Oct 02 12:44:01 NoKubernetes-115222 systemd[1]: Started containerd container runtime.
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.667191243Z" level=info msg="containerd successfully booted in 0.048143s"
	
	
	==> describe nodes <==
	command /bin/bash -c "sudo /var/lib/minikube/binaries/v0.0.0/kubectl describe nodes --kubeconfig=/var/lib/minikube/kubeconfig" failed with error: /bin/bash -c "sudo /var/lib/minikube/binaries/v0.0.0/kubectl describe nodes --kubeconfig=/var/lib/minikube/kubeconfig": Process exited with status 1
	stdout:
	
	stderr:
	sudo: a terminal is required to read the password; either use the -S option to read from standard input or configure an askpass helper
	sudo: a password is required
	
	
	==> dmesg <==
	[  +0.000008] ll header: 00000000: ff ff ff ff ff ff fe 40 f6 28 01 b6 08 06
	[  +2.473371] IPv4: martian source 10.244.0.1 from 10.244.0.2, on dev eth0
	[  +0.000006] ll header: 00000000: ff ff ff ff ff ff da 82 9e b7 f3 1d 08 06
	[ +12.115583] IPv4: martian source 10.244.0.1 from 10.244.0.3, on dev eth0
	[  +0.000007] ll header: 00000000: ff ff ff ff ff ff fa c9 db f3 a8 cb 08 06
	[  +0.000453] IPv4: martian source 10.244.0.3 from 10.244.0.2, on dev eth0
	[  +0.000004] ll header: 00000000: ff ff ff ff ff ff fe 40 f6 28 01 b6 08 06
	[  +8.697243] IPv4: martian source 10.244.0.1 from 10.244.0.2, on dev eth0
	[  +0.000006] ll header: 00000000: ff ff ff ff ff ff 12 55 f8 67 01 e2 08 06
	[Oct 2 11:45] IPv4: martian source 10.244.0.1 from 10.244.0.2, on dev eth0
	[  +0.000007] ll header: 00000000: ff ff ff ff ff ff 4a b5 6b 10 b4 b6 08 06
	[  +0.072240] IPv4: martian source 10.244.0.1 from 10.244.0.3, on dev eth0
	[  +0.000009] ll header: 00000000: ff ff ff ff ff ff 1a 6c 71 80 c4 33 08 06
	[  +1.016796] IPv4: martian source 10.244.0.1 from 10.244.0.3, on dev eth0
	[  +0.000009] ll header: 00000000: ff ff ff ff ff ff 1a 81 f6 69 f7 42 08 06
	[  +0.000541] IPv4: martian source 10.244.0.3 from 10.244.0.2, on dev eth0
	[  +0.000005] ll header: 00000000: ff ff ff ff ff ff 12 55 f8 67 01 e2 08 06
	[  +6.794314] IPv4: martian source 10.244.0.1 from 10.244.0.3, on dev eth0
	[  +0.000006] ll header: 00000000: ff ff ff ff ff ff 6e d7 18 68 e2 41 08 06
	[  +0.000359] IPv4: martian source 10.244.0.3 from 10.244.0.2, on dev eth0
	[  +0.000016] ll header: 00000000: ff ff ff ff ff ff da 82 9e b7 f3 1d 08 06
	[ +35.779860] IPv4: martian source 10.244.0.1 from 10.244.0.4, on dev eth0
	[  +0.000011] ll header: 00000000: ff ff ff ff ff ff 42 d2 67 3a d3 72 08 06
	[  +0.000391] IPv4: martian source 10.244.0.4 from 10.244.0.3, on dev eth0
	[  +0.000004] ll header: 00000000: ff ff ff ff ff ff 1a 6c 71 80 c4 33 08 06
	
	
	==> kernel <==
	 12:44:02 up  2:26,  0 users,  load average: 4.99, 2.09, 1.54
	Linux NoKubernetes-115222 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux
	PRETTY_NAME="Ubuntu 22.04.5 LTS"
	
	
	==> kubelet <==
	-- No entries --
	
-- /stdout --
helpers_test.go:262: (dbg) Run:  out/minikube-linux-amd64 status --format={{.APIServer}} -p NoKubernetes-115222 -n NoKubernetes-115222
helpers_test.go:262: (dbg) Non-zero exit: out/minikube-linux-amd64 status --format={{.APIServer}} -p NoKubernetes-115222 -n NoKubernetes-115222: exit status 6 (277.805339ms)
-- stdout --
	Stopped
	WARNING: Your kubectl is pointing to stale minikube-vm.
	To fix the kubectl context, run `minikube update-context`
-- /stdout --
** stderr ** 
	E1002 12:44:03.183252  588969 status.go:458] kubeconfig endpoint: get endpoint: "NoKubernetes-115222" does not appear in /home/jenkins/minikube-integration/21139-381342/kubeconfig
** /stderr **
helpers_test.go:262: status error: exit status 6 (may be ok)
helpers_test.go:264: "NoKubernetes-115222" apiserver is not running, skipping kubectl commands (state="Stopped")
helpers_test.go:222: -----------------------post-mortem--------------------------------
helpers_test.go:223: ======>  post-mortem[TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads]: network settings <======
helpers_test.go:230: HOST ENV snapshots: PROXY env: HTTP_PROXY="<empty>" HTTPS_PROXY="<empty>" NO_PROXY="<empty>"
helpers_test.go:238: ======>  post-mortem[TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads]: docker inspect <======
helpers_test.go:239: (dbg) Run:  docker inspect NoKubernetes-115222
helpers_test.go:243: (dbg) docker inspect NoKubernetes-115222:
-- stdout --
	[
	    {
	        "Id": "b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17",
	        "Created": "2025-10-02T12:43:58.931222211Z",
	        "Path": "/usr/local/bin/entrypoint",
	        "Args": [
	            "/sbin/init"
	        ],
	        "State": {
	            "Status": "running",
	            "Running": true,
	            "Paused": false,
	            "Restarting": false,
	            "OOMKilled": false,
	            "Dead": false,
	            "Pid": 586336,
	            "ExitCode": 0,
	            "Error": "",
	            "StartedAt": "2025-10-02T12:43:58.960923912Z",
	            "FinishedAt": "0001-01-01T00:00:00Z"
	        },
	        "Image": "sha256:c6b5532e987b5b4f5fc9cb0336e378ed49c0542bad8cbfc564b71e977a6269de",
	        "ResolvConfPath": "/var/lib/docker/containers/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17/resolv.conf",
	        "HostnamePath": "/var/lib/docker/containers/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17/hostname",
	        "HostsPath": "/var/lib/docker/containers/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17/hosts",
	        "LogPath": "/var/lib/docker/containers/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17/b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17-json.log",
	        "Name": "/NoKubernetes-115222",
	        "RestartCount": 0,
	        "Driver": "overlay2",
	        "Platform": "linux",
	        "MountLabel": "",
	        "ProcessLabel": "",
	        "AppArmorProfile": "unconfined",
	        "ExecIDs": null,
	        "HostConfig": {
	            "Binds": [
	                "/lib/modules:/lib/modules:ro",
	                "NoKubernetes-115222:/var"
	            ],
	            "ContainerIDFile": "",
	            "LogConfig": {
	                "Type": "json-file",
	                "Config": {
	                    "max-size": "100m"
	                }
	            },
	            "NetworkMode": "NoKubernetes-115222",
	            "PortBindings": {
	                "22/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ],
	                "2376/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ],
	                "32443/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ],
	                "5000/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ],
	                "8443/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": ""
	                    }
	                ]
	            },
	            "RestartPolicy": {
	                "Name": "no",
	                "MaximumRetryCount": 0
	            },
	            "AutoRemove": false,
	            "VolumeDriver": "",
	            "VolumesFrom": null,
	            "ConsoleSize": [
	                0,
	                0
	            ],
	            "CapAdd": null,
	            "CapDrop": null,
	            "CgroupnsMode": "private",
	            "Dns": [],
	            "DnsOptions": [],
	            "DnsSearch": [],
	            "ExtraHosts": null,
	            "GroupAdd": null,
	            "IpcMode": "private",
	            "Cgroup": "",
	            "Links": null,
	            "OomScoreAdj": 0,
	            "PidMode": "",
	            "Privileged": true,
	            "PublishAllPorts": false,
	            "ReadonlyRootfs": false,
	            "SecurityOpt": [
	                "seccomp=unconfined",
	                "apparmor=unconfined",
	                "label=disable"
	            ],
	            "Tmpfs": {
	                "/run": "",
	                "/tmp": ""
	            },
	            "UTSMode": "",
	            "UsernsMode": "",
	            "ShmSize": 67108864,
	            "Runtime": "runc",
	            "Isolation": "",
	            "CpuShares": 0,
	            "Memory": 3221225472,
	            "NanoCpus": 0,
	            "CgroupParent": "",
	            "BlkioWeight": 0,
	            "BlkioWeightDevice": [],
	            "BlkioDeviceReadBps": [],
	            "BlkioDeviceWriteBps": [],
	            "BlkioDeviceReadIOps": [],
	            "BlkioDeviceWriteIOps": [],
	            "CpuPeriod": 0,
	            "CpuQuota": 0,
	            "CpuRealtimePeriod": 0,
	            "CpuRealtimeRuntime": 0,
	            "CpusetCpus": "",
	            "CpusetMems": "",
	            "Devices": [],
	            "DeviceCgroupRules": null,
	            "DeviceRequests": null,
	            "MemoryReservation": 0,
	            "MemorySwap": 6442450944,
	            "MemorySwappiness": null,
	            "OomKillDisable": null,
	            "PidsLimit": null,
	            "Ulimits": [],
	            "CpuCount": 0,
	            "CpuPercent": 0,
	            "IOMaximumIOps": 0,
	            "IOMaximumBandwidth": 0,
	            "MaskedPaths": null,
	            "ReadonlyPaths": null
	        },
	        "GraphDriver": {
	            "Data": {
	                "ID": "b4bf0efb9b9471d596a391699f7eb3f1ebd7d9ff494c6577ec18a11e4c051c17",
	                "LowerDir": "/var/lib/docker/overlay2/7caa147e4324acc509479494c329388ece54f4832250947aeb81d55f419ff34b-init/diff:/var/lib/docker/overlay2/74e23e2d414e71bc60aae2442f772d94c45dcc7da38ffe98fa74cb259c3e7865/diff",
	                "MergedDir": "/var/lib/docker/overlay2/7caa147e4324acc509479494c329388ece54f4832250947aeb81d55f419ff34b/merged",
	                "UpperDir": "/var/lib/docker/overlay2/7caa147e4324acc509479494c329388ece54f4832250947aeb81d55f419ff34b/diff",
	                "WorkDir": "/var/lib/docker/overlay2/7caa147e4324acc509479494c329388ece54f4832250947aeb81d55f419ff34b/work"
	            },
	            "Name": "overlay2"
	        },
	        "Mounts": [
	            {
	                "Type": "bind",
	                "Source": "/lib/modules",
	                "Destination": "/lib/modules",
	                "Mode": "ro",
	                "RW": false,
	                "Propagation": "rprivate"
	            },
	            {
	                "Type": "volume",
	                "Name": "NoKubernetes-115222",
	                "Source": "/var/lib/docker/volumes/NoKubernetes-115222/_data",
	                "Destination": "/var",
	                "Driver": "local",
	                "Mode": "z",
	                "RW": true,
	                "Propagation": ""
	            }
	        ],
	        "Config": {
	            "Hostname": "NoKubernetes-115222",
	            "Domainname": "",
	            "User": "",
	            "AttachStdin": false,
	            "AttachStdout": false,
	            "AttachStderr": false,
	            "ExposedPorts": {
	                "22/tcp": {},
	                "2376/tcp": {},
	                "32443/tcp": {},
	                "5000/tcp": {},
	                "8443/tcp": {}
	            },
	            "Tty": true,
	            "OpenStdin": false,
	            "StdinOnce": false,
	            "Env": [
	                "container=docker",
	                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
	            ],
	            "Cmd": null,
	            "Image": "gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1",
	            "Volumes": null,
	            "WorkingDir": "/",
	            "Entrypoint": [
	                "/usr/local/bin/entrypoint",
	                "/sbin/init"
	            ],
	            "OnBuild": null,
	            "Labels": {
	                "created_by.minikube.sigs.k8s.io": "true",
	                "mode.minikube.sigs.k8s.io": "NoKubernetes-115222",
	                "name.minikube.sigs.k8s.io": "NoKubernetes-115222",
	                "role.minikube.sigs.k8s.io": ""
	            },
	            "StopSignal": "SIGRTMIN+3"
	        },
	        "NetworkSettings": {
	            "Bridge": "",
	            "SandboxID": "684712f731d0b43f5cd680d6bec133d72160dd1584ed185f8bc18deca1caec9b",
	            "SandboxKey": "/var/run/docker/netns/684712f731d0",
	            "Ports": {
	                "22/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33390"
	                    }
	                ],
	                "2376/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33391"
	                    }
	                ],
	                "32443/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33394"
	                    }
	                ],
	                "5000/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33392"
	                    }
	                ],
	                "8443/tcp": [
	                    {
	                        "HostIp": "127.0.0.1",
	                        "HostPort": "33393"
	                    }
	                ]
	            },
	            "HairpinMode": false,
	            "LinkLocalIPv6Address": "",
	            "LinkLocalIPv6PrefixLen": 0,
	            "SecondaryIPAddresses": null,
	            "SecondaryIPv6Addresses": null,
	            "EndpointID": "",
	            "Gateway": "",
	            "GlobalIPv6Address": "",
	            "GlobalIPv6PrefixLen": 0,
	            "IPAddress": "",
	            "IPPrefixLen": 0,
	            "IPv6Gateway": "",
	            "MacAddress": "",
	            "Networks": {
	                "NoKubernetes-115222": {
	                    "IPAMConfig": {
	                        "IPv4Address": "192.168.85.2"
	                    },
	                    "Links": null,
	                    "Aliases": null,
	                    "MacAddress": "06:12:47:d1:15:94",
	                    "DriverOpts": null,
	                    "GwPriority": 0,
	                    "NetworkID": "b83a7484a0439ab1227e41307c48bfff54849f6fa122889143471cb7e6e2a017",
	                    "EndpointID": "ca65baae7c15a7c93d2ff87d7e4421a583cf63977e817f6bde2e82559b78c05f",
	                    "Gateway": "192.168.85.1",
	                    "IPAddress": "192.168.85.2",
	                    "IPPrefixLen": 24,
	                    "IPv6Gateway": "",
	                    "GlobalIPv6Address": "",
	                    "GlobalIPv6PrefixLen": 0,
	                    "DNSNames": [
	                        "NoKubernetes-115222",
	                        "b4bf0efb9b94"
	                    ]
	                }
	            }
	        }
	    }
	]
-- /stdout --
helpers_test.go:247: (dbg) Run:  out/minikube-linux-amd64 status --format={{.Host}} -p NoKubernetes-115222 -n NoKubernetes-115222
helpers_test.go:247: (dbg) Non-zero exit: out/minikube-linux-amd64 status --format={{.Host}} -p NoKubernetes-115222 -n NoKubernetes-115222: exit status 6 (310.324364ms)
-- stdout --
	Running
	WARNING: Your kubectl is pointing to stale minikube-vm.
	To fix the kubectl context, run `minikube update-context`
-- /stdout --
** stderr ** 
	E1002 12:44:03.506855  589106 status.go:458] kubeconfig endpoint: get endpoint: "NoKubernetes-115222" does not appear in /home/jenkins/minikube-integration/21139-381342/kubeconfig
** /stderr **
helpers_test.go:247: status error: exit status 6 (may be ok)
helpers_test.go:252: <<< TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads FAILED: start of post-mortem logs <<<
helpers_test.go:253: ======>  post-mortem[TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads]: minikube logs <======
helpers_test.go:255: (dbg) Run:  out/minikube-linux-amd64 -p NoKubernetes-115222 logs -n 25
helpers_test.go:260: TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads logs: 
-- stdout --
	
	==> Audit <==
	┌─────────┬───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┬─────────────────────────────┬─────────┬─────────┬─────────────────────┬─────────────────────┐
	│ COMMAND │                                                             ARGS                                                              │           PROFILE           │  USER   │ VERSION │     START TIME      │      END TIME       │
	├─────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼─────────────────────────────┼─────────┼─────────┼─────────────────────┼─────────────────────┤
	│ stop    │ -p scheduled-stop-180674 --schedule 5m                                                                                        │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 5m                                                                                        │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --cancel-scheduled                                                                                   │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │ 02 Oct 25 12:41 UTC │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │                     │
	│ stop    │ -p scheduled-stop-180674 --schedule 15s                                                                                       │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:41 UTC │ 02 Oct 25 12:42 UTC │
	│ delete  │ -p scheduled-stop-180674                                                                                                      │ scheduled-stop-180674       │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:42 UTC │
	│ start   │ -p insufficient-storage-013758 --memory=3072 --output=json --wait=true --driver=docker  --container-runtime=containerd        │ insufficient-storage-013758 │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │                     │
	│ delete  │ -p insufficient-storage-013758                                                                                                │ insufficient-storage-013758 │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:42 UTC │
	│ start   │ -p offline-containerd-106797 --alsologtostderr -v=1 --memory=3072 --wait=true --driver=docker  --container-runtime=containerd │ offline-containerd-106797   │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p NoKubernetes-115222 --no-kubernetes --kubernetes-version=v1.28.0 --driver=docker  --container-runtime=containerd           │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │                     │
	│ start   │ -p NoKubernetes-115222 --memory=3072 --alsologtostderr -v=5 --driver=docker  --container-runtime=containerd                   │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p running-upgrade-583950 --memory=3072 --vm-driver=docker  --container-runtime=containerd                                    │ running-upgrade-583950      │ jenkins │ v1.32.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p stopped-upgrade-250350 --memory=3072 --vm-driver=docker  --container-runtime=containerd                                    │ stopped-upgrade-250350      │ jenkins │ v1.32.0 │ 02 Oct 25 12:42 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p NoKubernetes-115222 --no-kubernetes --memory=3072 --alsologtostderr -v=5 --driver=docker  --container-runtime=containerd   │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:43 UTC │
	│ delete  │ -p offline-containerd-106797                                                                                                  │ offline-containerd-106797   │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p running-upgrade-583950 --memory=3072 --alsologtostderr -v=1 --driver=docker  --container-runtime=containerd                │ running-upgrade-583950      │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │                     │
	│ stop    │ stopped-upgrade-250350 stop                                                                                                   │ stopped-upgrade-250350      │ jenkins │ v1.32.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p missing-upgrade-404552 --memory=3072 --driver=docker  --container-runtime=containerd                                       │ missing-upgrade-404552      │ jenkins │ v1.32.0 │ 02 Oct 25 12:43 UTC │                     │
	│ start   │ -p stopped-upgrade-250350 --memory=3072 --alsologtostderr -v=1 --driver=docker  --container-runtime=containerd                │ stopped-upgrade-250350      │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │                     │
	│ delete  │ -p NoKubernetes-115222                                                                                                        │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:43 UTC │
	│ start   │ -p NoKubernetes-115222 --no-kubernetes --memory=3072 --alsologtostderr -v=5 --driver=docker  --container-runtime=containerd   │ NoKubernetes-115222         │ jenkins │ v1.37.0 │ 02 Oct 25 12:43 UTC │ 02 Oct 25 12:44 UTC │
	└─────────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┴─────────────────────────────┴─────────┴─────────┴─────────────────────┴─────────────────────┘
	
	
	==> Last Start <==
	Log file created at: 2025/10/02 12:43:58
	Running on machine: ubuntu-20-agent-6
	Binary: Built with gc go1.24.6 for linux/amd64
	Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg
	I1002 12:43:58.143632  585741 out.go:360] Setting OutFile to fd 1 ...
	I1002 12:43:58.143942  585741 out.go:408] TERM=,COLORTERM=, which probably does not support color
	I1002 12:43:58.143954  585741 out.go:374] Setting ErrFile to fd 2...
	I1002 12:43:58.143958  585741 out.go:408] TERM=,COLORTERM=, which probably does not support color
	I1002 12:43:58.144172  585741 root.go:338] Updating PATH: /home/jenkins/minikube-integration/21139-381342/.minikube/bin
	I1002 12:43:58.144654  585741 out.go:368] Setting JSON to false
	I1002 12:43:58.146105  585741 start.go:130] hostinfo: {"hostname":"ubuntu-20-agent-6","uptime":8776,"bootTime":1759400262,"procs":281,"os":"linux","platform":"ubuntu","platformFamily":"debian","platformVersion":"22.04","kernelVersion":"6.8.0-1041-gcp","kernelArch":"x86_64","virtualizationSystem":"kvm","virtualizationRole":"guest","hostId":"591c9f12-2938-3743-e2bf-c56a050d43d1"}
	I1002 12:43:58.146235  585741 start.go:140] virtualization: kvm guest
	I1002 12:43:58.148027  585741 out.go:179] * [NoKubernetes-115222] minikube v1.37.0 on Ubuntu 22.04 (kvm/amd64)
	I1002 12:43:58.149166  585741 out.go:179]   - MINIKUBE_LOCATION=21139
	I1002 12:43:58.149202  585741 notify.go:220] Checking for updates...
	I1002 12:43:58.151111  585741 out.go:179]   - MINIKUBE_SUPPRESS_DOCKER_PERFORMANCE=true
	I1002 12:43:58.152058  585741 out.go:179]   - KUBECONFIG=/home/jenkins/minikube-integration/21139-381342/kubeconfig
	I1002 12:43:58.153034  585741 out.go:179]   - MINIKUBE_HOME=/home/jenkins/minikube-integration/21139-381342/.minikube
	I1002 12:43:58.157294  585741 out.go:179]   - MINIKUBE_BIN=out/minikube-linux-amd64
	I1002 12:43:58.158404  585741 out.go:179]   - MINIKUBE_FORCE_SYSTEMD=
	I1002 12:43:56.113021  581986 kic_runner.go:191] docker (temp): /home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa.pub --> /home/docker/.ssh/authorized_keys (381 bytes)
	I1002 12:43:56.140789  581986 cli_runner.go:164] Run: docker container inspect missing-upgrade-404552 --format={{.State.Status}}
	I1002 12:43:56.158961  581986 kic_runner.go:93] Run: chown docker:docker /home/docker/.ssh/authorized_keys
	I1002 12:43:56.158975  581986 kic_runner.go:114] Args: [docker exec --privileged missing-upgrade-404552 chown docker:docker /home/docker/.ssh/authorized_keys]
	I1002 12:43:56.208006  581986 cli_runner.go:164] Run: docker container inspect missing-upgrade-404552 --format={{.State.Status}}
	I1002 12:43:56.228078  581986 machine.go:88] provisioning docker machine ...
	I1002 12:43:56.228119  581986 ubuntu.go:169] provisioning hostname "missing-upgrade-404552"
	I1002 12:43:56.228181  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:56.247049  581986 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:56.248203  581986 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x808a40] 0x80b720 <nil>  [] 0s} 127.0.0.1 33385 <nil> <nil>}
	I1002 12:43:56.248220  581986 main.go:141] libmachine: About to run SSH command:
	sudo hostname missing-upgrade-404552 && echo "missing-upgrade-404552" | sudo tee /etc/hostname
	I1002 12:43:56.379515  581986 main.go:141] libmachine: SSH cmd err, output: <nil>: missing-upgrade-404552
	
	I1002 12:43:56.379591  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:56.398347  581986 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:56.398900  581986 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x808a40] 0x80b720 <nil>  [] 0s} 127.0.0.1 33385 <nil> <nil>}
	I1002 12:43:56.398922  581986 main.go:141] libmachine: About to run SSH command:
	
			if ! grep -xq '.*\smissing-upgrade-404552' /etc/hosts; then
				if grep -xq '127.0.1.1\s.*' /etc/hosts; then
					sudo sed -i 's/^127.0.1.1\s.*/127.0.1.1 missing-upgrade-404552/g' /etc/hosts;
				else 
					echo '127.0.1.1 missing-upgrade-404552' | sudo tee -a /etc/hosts; 
				fi
			fi
	I1002 12:43:56.515082  581986 main.go:141] libmachine: SSH cmd err, output: <nil>: 
	I1002 12:43:56.515103  581986 ubuntu.go:175] set auth options {CertDir:/home/jenkins/minikube-integration/21139-381342/.minikube CaCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem CaPrivateKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem CaCertRemotePath:/etc/docker/ca.pem ServerCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ServerKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem ClientKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem ServerCertRemotePath:/etc/docker/server.pem ServerKeyRemotePath:/etc/docker/server-key.pem ClientCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem ServerCertSANs:[] StorePath:/home/jenkins/minikube-integration/21139-381342/.minikube}
	I1002 12:43:56.515141  581986 ubuntu.go:177] setting up certificates
	I1002 12:43:56.515151  581986 provision.go:83] configureAuth start
	I1002 12:43:56.515204  581986 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" missing-upgrade-404552
	I1002 12:43:56.534704  581986 provision.go:138] copyHostCerts
	I1002 12:43:56.534762  581986 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem, removing ...
	I1002 12:43:56.534772  581986 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem
	I1002 12:43:56.534864  581986 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem (1082 bytes)
	I1002 12:43:56.535019  581986 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem, removing ...
	I1002 12:43:56.535027  581986 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem
	I1002 12:43:56.535078  581986 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem (1123 bytes)
	I1002 12:43:56.535178  581986 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem, removing ...
	I1002 12:43:56.535198  581986 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem
	I1002 12:43:56.535249  581986 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem (1675 bytes)
	I1002 12:43:56.535380  581986 provision.go:112] generating server cert: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ca-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem private-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem org=jenkins.missing-upgrade-404552 san=[192.168.76.2 127.0.0.1 localhost 127.0.0.1 minikube missing-upgrade-404552]
	I1002 12:43:56.697791  581986 provision.go:172] copyRemoteCerts
	I1002 12:43:56.697863  581986 ssh_runner.go:195] Run: sudo mkdir -p /etc/docker /etc/docker /etc/docker
	I1002 12:43:56.697903  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:56.715683  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:56.800884  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem --> /etc/docker/server-key.pem (1679 bytes)
	I1002 12:43:56.826192  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /etc/docker/ca.pem (1082 bytes)
	I1002 12:43:56.849217  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem --> /etc/docker/server.pem (1241 bytes)
	I1002 12:43:56.871626  581986 provision.go:86] duration metric: configureAuth took 356.465021ms
	I1002 12:43:56.871642  581986 ubuntu.go:193] setting minikube options for container-runtime
	I1002 12:43:56.871791  581986 config.go:182] Loaded profile config "missing-upgrade-404552": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:56.871796  581986 machine.go:91] provisioned docker machine in 643.706565ms
	I1002 12:43:56.871801  581986 client.go:171] LocalClient.Create took 5.519165598s
	I1002 12:43:56.871817  581986 start.go:167] duration metric: libmachine.API.Create for "missing-upgrade-404552" took 5.51922352s
	I1002 12:43:56.871822  581986 start.go:300] post-start starting for "missing-upgrade-404552" (driver="docker")
	I1002 12:43:56.871849  581986 start.go:329] creating required directories: [/etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs]
	I1002 12:43:56.871902  581986 ssh_runner.go:195] Run: sudo mkdir -p /etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs
	I1002 12:43:56.871947  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:56.890200  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:56.982746  581986 ssh_runner.go:195] Run: cat /etc/os-release
	I1002 12:43:56.985888  581986 main.go:141] libmachine: Couldn't set key VERSION_CODENAME, no corresponding struct field found
	I1002 12:43:56.985921  581986 main.go:141] libmachine: Couldn't set key PRIVACY_POLICY_URL, no corresponding struct field found
	I1002 12:43:56.985938  581986 main.go:141] libmachine: Couldn't set key UBUNTU_CODENAME, no corresponding struct field found
	I1002 12:43:56.985944  581986 info.go:137] Remote host: Ubuntu 22.04.3 LTS
	I1002 12:43:56.985953  581986 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/addons for local assets ...
	I1002 12:43:56.986014  581986 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/files for local assets ...
	I1002 12:43:56.986100  581986 filesync.go:149] local asset: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem -> 3849552.pem in /etc/ssl/certs
	I1002 12:43:56.986211  581986 ssh_runner.go:195] Run: sudo mkdir -p /etc/ssl/certs
	I1002 12:43:56.994580  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem --> /etc/ssl/certs/3849552.pem (1708 bytes)
	I1002 12:43:57.020319  581986 start.go:303] post-start completed in 148.471394ms
	I1002 12:43:57.020659  581986 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" missing-upgrade-404552
	I1002 12:43:57.038303  581986 profile.go:148] Saving config to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/config.json ...
	I1002 12:43:57.038500  581986 ssh_runner.go:195] Run: sh -c "df -h /var | awk 'NR==2{print $5}'"
	I1002 12:43:57.038529  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:57.055262  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:57.135385  581986 ssh_runner.go:195] Run: sh -c "df -BG /var | awk 'NR==2{print $4}'"
	I1002 12:43:57.139586  581986 start.go:128] duration metric: createHost completed in 5.789433984s
	I1002 12:43:57.139600  581986 start.go:83] releasing machines lock for "missing-upgrade-404552", held for 5.789554894s
	I1002 12:43:57.139657  581986 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" missing-upgrade-404552
	I1002 12:43:57.156682  581986 ssh_runner.go:195] Run: cat /version.json
	I1002 12:43:57.156713  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:57.156773  581986 ssh_runner.go:195] Run: curl -sS -m 2 https://registry.k8s.io/
	I1002 12:43:57.156843  581986 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" missing-upgrade-404552
	I1002 12:43:57.175099  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:57.175452  581986 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33385 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/missing-upgrade-404552/id_rsa Username:docker}
	I1002 12:43:57.348369  581986 ssh_runner.go:195] Run: systemctl --version
	I1002 12:43:57.353082  581986 ssh_runner.go:195] Run: sh -c "stat /etc/cni/net.d/*loopback.conf*"
	I1002 12:43:57.357358  581986 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f -name *loopback.conf* -not -name *.mk_disabled -exec sh -c "grep -q loopback {} && ( grep -q name {} || sudo sed -i '/"type": "loopback"/i \ \ \ \ "name": "loopback",' {} ) && sudo sed -i 's|"cniVersion": ".*"|"cniVersion": "1.0.0"|g' {}" ;
	I1002 12:43:57.384724  581986 cni.go:230] loopback cni configuration patched: "/etc/cni/net.d/*loopback.conf*" found
	I1002 12:43:57.384773  581986 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f ( ( -name *bridge* -or -name *podman* ) -and -not -name *.mk_disabled ) -printf "%p, " -exec sh -c "sudo mv {} {}.mk_disabled" ;
	I1002 12:43:57.409586  581986 cni.go:262] disabled [/etc/cni/net.d/87-podman-bridge.conflist, /etc/cni/net.d/100-crio-bridge.conf] bridge cni config(s)
	I1002 12:43:57.409603  581986 start.go:472] detecting cgroup driver to use...
	I1002 12:43:57.409629  581986 detect.go:199] detected "systemd" cgroup driver on host os
	I1002 12:43:57.409668  581986 ssh_runner.go:195] Run: sudo systemctl stop -f crio
	I1002 12:43:57.421371  581986 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service crio
	I1002 12:43:57.431690  581986 docker.go:203] disabling cri-docker service (if available) ...
	I1002 12:43:57.431744  581986 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.socket
	I1002 12:43:57.444305  581986 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.service
	I1002 12:43:57.457523  581986 ssh_runner.go:195] Run: sudo systemctl disable cri-docker.socket
	I1002 12:43:57.524054  581986 ssh_runner.go:195] Run: sudo systemctl mask cri-docker.service
	I1002 12:43:57.598654  581986 docker.go:219] disabling docker service ...
	I1002 12:43:57.598703  581986 ssh_runner.go:195] Run: sudo systemctl stop -f docker.socket
	I1002 12:43:57.616246  581986 ssh_runner.go:195] Run: sudo systemctl stop -f docker.service
	I1002 12:43:57.627447  581986 ssh_runner.go:195] Run: sudo systemctl disable docker.socket
	I1002 12:43:57.704488  581986 ssh_runner.go:195] Run: sudo systemctl mask docker.service
	I1002 12:43:57.785554  581986 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service docker
	I1002 12:43:57.797381  581986 ssh_runner.go:195] Run: /bin/bash -c "sudo mkdir -p /etc && printf %s "runtime-endpoint: unix:///run/containerd/containerd.sock
	" | sudo tee /etc/crictl.yaml"
	I1002 12:43:57.813760  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)sandbox_image = .*$|\1sandbox_image = "registry.k8s.io/pause:3.9"|' /etc/containerd/config.toml"
	I1002 12:43:57.829274  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)restrict_oom_score_adj = .*$|\1restrict_oom_score_adj = false|' /etc/containerd/config.toml"
	I1002 12:43:57.839084  581986 containerd.go:145] configuring containerd to use "systemd" as cgroup driver...
	I1002 12:43:57.839137  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)SystemdCgroup = .*$|\1SystemdCgroup = true|g' /etc/containerd/config.toml"
	I1002 12:43:57.848695  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runtime.v1.linux"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:43:57.858251  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i '/systemd_cgroup/d' /etc/containerd/config.toml"
	I1002 12:43:57.868364  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runc.v1"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:43:57.878878  581986 ssh_runner.go:195] Run: sh -c "sudo rm -rf /etc/cni/net.mk"
	I1002 12:43:57.888799  581986 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)conf_dir = .*$|\1conf_dir = "/etc/cni/net.d"|g' /etc/containerd/config.toml"
	I1002 12:43:57.898544  581986 ssh_runner.go:195] Run: sudo sysctl net.bridge.bridge-nf-call-iptables
	I1002 12:43:57.929002  581986 ssh_runner.go:195] Run: sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
	I1002 12:43:57.938069  581986 ssh_runner.go:195] Run: sudo systemctl daemon-reload
	I1002 12:43:58.002820  581986 ssh_runner.go:195] Run: sudo systemctl restart containerd
	I1002 12:43:58.106242  581986 start.go:519] Will wait 60s for socket path /run/containerd/containerd.sock
	I1002 12:43:58.106289  581986 ssh_runner.go:195] Run: stat /run/containerd/containerd.sock
	I1002 12:43:58.110402  581986 start.go:540] Will wait 60s for crictl version
	I1002 12:43:58.110447  581986 ssh_runner.go:195] Run: which crictl
	I1002 12:43:58.114209  581986 ssh_runner.go:195] Run: sudo /usr/bin/crictl version
	I1002 12:43:58.149277  581986 start.go:556] Version:  0.1.0
	RuntimeName:  containerd
	RuntimeVersion:  1.6.24
	RuntimeApiVersion:  v1
	I1002 12:43:58.149349  581986 ssh_runner.go:195] Run: containerd --version
	I1002 12:43:58.174699  581986 ssh_runner.go:195] Run: containerd --version
	I1002 12:43:58.203536  581986 out.go:177] * Preparing Kubernetes v1.28.3 on containerd 1.6.24 ...
	I1002 12:43:58.160098  585741 config.go:182] Loaded profile config "missing-upgrade-404552": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:58.160250  585741 config.go:182] Loaded profile config "running-upgrade-583950": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:58.160372  585741 config.go:182] Loaded profile config "stopped-upgrade-250350": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:58.160403  585741 start.go:1898] No Kubernetes flag is set, setting Kubernetes version to v0.0.0
	I1002 12:43:58.160483  585741 driver.go:421] Setting default libvirt URI to qemu:///system
	I1002 12:43:58.184848  585741 docker.go:123] docker version: linux-28.4.0:Docker Engine - Community
	I1002 12:43:58.184952  585741 cli_runner.go:164] Run: docker system info --format "{{json .}}"
	I1002 12:43:58.239813  585741 info.go:266] docker info: {ID:TS6T:UINC:MIYS:RZPA:KS6T:4JQK:7JHN:D6RA:LDP2:MHAE:G32M:C5NQ Containers:3 ContainersRunning:3 ContainersPaused:0 ContainersStopped:0 Images:4 Driver:overlay2 DriverStatus:[[Backing Filesystem extfs] [Supports d_type true] [Using metacopy false] [Native Overlay Diff true] [userxattr false]] SystemStatus:<nil> Plugins:{Volume:[local] Network:[bridge host ipvlan macvlan null overlay] Authorization:<nil> Log:[awslogs fluentd gcplogs gelf journald json-file local splunk syslog]} MemoryLimit:true SwapLimit:true KernelMemory:false KernelMemoryTCP:false CPUCfsPeriod:true CPUCfsQuota:true CPUShares:true CPUSet:true PidsLimit:true IPv4Forwarding:true BridgeNfIptables:false BridgeNfIP6Tables:false Debug:false NFd:68 OomKillDisable:false NGoroutines:77 SystemTime:2025-10-02 12:43:58.229497835 +0000 UTC LoggingDriver:json-file CgroupDriver:cgroupfs NEventsListener:0 KernelVersion:6.8.0-1041-gcp OperatingSystem:Ubuntu 22.04.5 LTS OSType:linux Architecture:
x86_64 IndexServerAddress:https://index.docker.io/v1/ RegistryConfig:{AllowNondistributableArtifactsCIDRs:[] AllowNondistributableArtifactsHostnames:[] InsecureRegistryCIDRs:[::1/128 127.0.0.0/8] IndexConfigs:{DockerIo:{Name:docker.io Mirrors:[] Secure:true Official:true}} Mirrors:[]} NCPU:8 MemTotal:33652174848 GenericResources:<nil> DockerRootDir:/var/lib/docker HTTPProxy: HTTPSProxy: NoProxy: Name:ubuntu-20-agent-6 Labels:[] ExperimentalBuild:false ServerVersion:28.4.0 ClusterStore: ClusterAdvertise: Runtimes:{Runc:{Path:runc}} DefaultRuntime:runc Swarm:{NodeID: NodeAddr: LocalNodeState:inactive ControlAvailable:false Error: RemoteManagers:<nil>} LiveRestoreEnabled:false Isolation: InitBinary:docker-init ContainerdCommit:{ID:b98a3aace656320842a23f4a392a33f46af97866 Expected:} RuncCommit:{ID:v1.3.0-0-g4ca628d1 Expected:} InitCommit:{ID:de40ad0 Expected:} SecurityOptions:[name=apparmor name=seccomp,profile=builtin name=cgroupns] ProductLicense: Warnings:<nil> ServerErrors:[] ClientInfo:{Debug:false Plugins:[
map[Name:buildx Path:/usr/libexec/docker/cli-plugins/docker-buildx SchemaVersion:0.1.0 ShortDescription:Docker Buildx Vendor:Docker Inc. Version:v0.29.0] map[Name:compose Path:/usr/libexec/docker/cli-plugins/docker-compose SchemaVersion:0.1.0 ShortDescription:Docker Compose Vendor:Docker Inc. Version:v2.39.4] map[Name:model Path:/usr/libexec/docker/cli-plugins/docker-model SchemaVersion:0.1.0 ShortDescription:Docker Model Runner Vendor:Docker Inc. Version:v0.1.40] map[Name:scan Path:/usr/libexec/docker/cli-plugins/docker-scan SchemaVersion:0.1.0 ShortDescription:Docker Scan Vendor:Docker Inc. Version:v0.23.0]] Warnings:<nil>}}
	I1002 12:43:58.239954  585741 docker.go:318] overlay module found
	I1002 12:43:58.242138  585741 out.go:179] * Using the docker driver based on user configuration
	I1002 12:43:58.243209  585741 start.go:304] selected driver: docker
	I1002 12:43:58.243231  585741 start.go:924] validating driver "docker" against <nil>
	I1002 12:43:58.243242  585741 start.go:935] status for docker: {Installed:true Healthy:true Running:false NeedsImprovement:false Error:<nil> Reason: Fix: Doc: Version:}
	I1002 12:43:58.243815  585741 cli_runner.go:164] Run: docker system info --format "{{json .}}"
	I1002 12:43:58.299577  585741 info.go:266] docker info: {ID:TS6T:UINC:MIYS:RZPA:KS6T:4JQK:7JHN:D6RA:LDP2:MHAE:G32M:C5NQ Containers:3 ContainersRunning:3 ContainersPaused:0 ContainersStopped:0 Images:4 Driver:overlay2 DriverStatus:[[Backing Filesystem extfs] [Supports d_type true] [Using metacopy false] [Native Overlay Diff true] [userxattr false]] SystemStatus:<nil> Plugins:{Volume:[local] Network:[bridge host ipvlan macvlan null overlay] Authorization:<nil> Log:[awslogs fluentd gcplogs gelf journald json-file local splunk syslog]} MemoryLimit:true SwapLimit:true KernelMemory:false KernelMemoryTCP:false CPUCfsPeriod:true CPUCfsQuota:true CPUShares:true CPUSet:true PidsLimit:true IPv4Forwarding:true BridgeNfIptables:false BridgeNfIP6Tables:false Debug:false NFd:68 OomKillDisable:false NGoroutines:77 SystemTime:2025-10-02 12:43:58.288929094 +0000 UTC LoggingDriver:json-file CgroupDriver:cgroupfs NEventsListener:0 KernelVersion:6.8.0-1041-gcp OperatingSystem:Ubuntu 22.04.5 LTS OSType:linux Architecture:
x86_64 IndexServerAddress:https://index.docker.io/v1/ RegistryConfig:{AllowNondistributableArtifactsCIDRs:[] AllowNondistributableArtifactsHostnames:[] InsecureRegistryCIDRs:[::1/128 127.0.0.0/8] IndexConfigs:{DockerIo:{Name:docker.io Mirrors:[] Secure:true Official:true}} Mirrors:[]} NCPU:8 MemTotal:33652174848 GenericResources:<nil> DockerRootDir:/var/lib/docker HTTPProxy: HTTPSProxy: NoProxy: Name:ubuntu-20-agent-6 Labels:[] ExperimentalBuild:false ServerVersion:28.4.0 ClusterStore: ClusterAdvertise: Runtimes:{Runc:{Path:runc}} DefaultRuntime:runc Swarm:{NodeID: NodeAddr: LocalNodeState:inactive ControlAvailable:false Error: RemoteManagers:<nil>} LiveRestoreEnabled:false Isolation: InitBinary:docker-init ContainerdCommit:{ID:b98a3aace656320842a23f4a392a33f46af97866 Expected:} RuncCommit:{ID:v1.3.0-0-g4ca628d1 Expected:} InitCommit:{ID:de40ad0 Expected:} SecurityOptions:[name=apparmor name=seccomp,profile=builtin name=cgroupns] ProductLicense: Warnings:<nil> ServerErrors:[] ClientInfo:{Debug:false Plugins:[
map[Name:buildx Path:/usr/libexec/docker/cli-plugins/docker-buildx SchemaVersion:0.1.0 ShortDescription:Docker Buildx Vendor:Docker Inc. Version:v0.29.0] map[Name:compose Path:/usr/libexec/docker/cli-plugins/docker-compose SchemaVersion:0.1.0 ShortDescription:Docker Compose Vendor:Docker Inc. Version:v2.39.4] map[Name:model Path:/usr/libexec/docker/cli-plugins/docker-model SchemaVersion:0.1.0 ShortDescription:Docker Model Runner Vendor:Docker Inc. Version:v0.1.40] map[Name:scan Path:/usr/libexec/docker/cli-plugins/docker-scan SchemaVersion:0.1.0 ShortDescription:Docker Scan Vendor:Docker Inc. Version:v0.23.0]] Warnings:<nil>}}
	I1002 12:43:58.299663  585741 start.go:1898] No Kubernetes flag is set, setting Kubernetes version to v0.0.0
	I1002 12:43:58.299741  585741 start_flags.go:327] no existing cluster config was found, will generate one from the flags 
	I1002 12:43:58.300012  585741 start_flags.go:974] Wait components to verify : map[apiserver:true system_pods:true]
	I1002 12:43:58.301478  585741 out.go:179] * Using Docker driver with root privileges
	I1002 12:43:58.302563  585741 cni.go:84] Creating CNI manager for ""
	I1002 12:43:58.302631  585741 cni.go:143] "docker" driver + "containerd" runtime found, recommending kindnet
	I1002 12:43:58.302645  585741 start_flags.go:336] Found "CNI" CNI - setting NetworkPlugin=cni
	I1002 12:43:58.302670  585741 start.go:1898] No Kubernetes flag is set, setting Kubernetes version to v0.0.0
	I1002 12:43:58.302712  585741 start.go:348] cluster config:
	{Name:NoKubernetes-115222 KeepContext:false EmbedCerts:false MinikubeISO: KicBaseImage:gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 Memory:3072 CPUs:2 DiskSize:20000 Driver:docker HyperkitVpnKitSock: HyperkitVSockPorts:[] DockerEnv:[] ContainerVolumeMounts:[] InsecureRegistry:[] RegistryMirror:[] HostOnlyCIDR:192.168.59.1/24 HypervVirtualSwitch: HypervUseExternalSwitch:false HypervExternalAdapter: KVMNetwork:default KVMQemuURI:qemu:///system KVMGPU:false KVMHidden:false KVMNUMACount:1 APIServerPort:8443 DockerOpt:[] DisableDriverMounts:false NFSShare:[] NFSSharesRoot:/nfsshares UUID: NoVTXCheck:false DNSProxy:false HostDNSResolver:true HostOnlyNicType:virtio NatNicType:virtio SSHIPAddress: SSHUser:root SSHKey: SSHPort:22 KubernetesConfig:{KubernetesVersion:v0.0.0 ClusterName:NoKubernetes-115222 Namespace:default APIServerHAVIP: APIServerName:minikubeCA APIServerNames:[] APIServerIPs:[] DNSDomain:cluster.local ContainerRuntime:containerd C
RISocket: NetworkPlugin:cni FeatureGates: ServiceCIDR:10.96.0.0/12 ImageRepository: LoadBalancerStartIP: LoadBalancerEndIP: CustomIngressCert: RegistryAliases: ExtraOptions:[] ShouldLoadCachedImages:true EnableDefaultCNI:false CNI:} Nodes:[{Name: IP: Port:8443 KubernetesVersion:v0.0.0 ContainerRuntime:containerd ControlPlane:true Worker:true}] Addons:map[] CustomAddonImages:map[] CustomAddonRegistries:map[] VerifyComponents:map[apiserver:true system_pods:true] StartHostTimeout:6m0s ScheduledStop:<nil> ExposedPorts:[] ListenAddress: Network: Subnet: MultiNodeRequested:false ExtraDisks:0 CertExpiration:26280h0m0s MountString: Mount9PVersion:9p2000.L MountGID:docker MountIP: MountMSize:262144 MountOptions:[] MountPort:0 MountType:9p MountUID:docker BinaryMirror: DisableOptimizations:false DisableMetrics:false DisableCoreDNSLog:false CustomQemuFirmwarePath: SocketVMnetClientPath: SocketVMnetPath: StaticIP: SSHAuthSock: SSHAgentPID:0 GPUs: AutoPauseInterval:1m0s}
	I1002 12:43:58.303779  585741 out.go:179] * Starting minikube without Kubernetes in cluster NoKubernetes-115222
	I1002 12:43:58.304702  585741 cache.go:133] Beginning downloading kic base image for docker with containerd
	I1002 12:43:58.308050  585741 out.go:179] * Pulling base image v0.0.48 ...
	I1002 12:43:58.309193  585741 cache.go:58] Skipping Kubernetes image caching due to --no-kubernetes flag
	I1002 12:43:58.309306  585741 image.go:81] Checking for gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 in local docker daemon
	I1002 12:43:58.309348  585741 profile.go:143] Saving config to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/NoKubernetes-115222/config.json ...
	I1002 12:43:58.309398  585741 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/NoKubernetes-115222/config.json: {Name:mkf16e835ee733faaa5453f2629f21d135e155cb Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.331669  585741 image.go:100] Found gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 in local docker daemon, skipping pull
	I1002 12:43:58.331703  585741 cache.go:157] gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 exists in daemon, skipping load
	I1002 12:43:58.331726  585741 cache.go:242] Successfully downloaded all kic artifacts
	I1002 12:43:58.331756  585741 start.go:360] acquireMachinesLock for NoKubernetes-115222: {Name:mkcdc227f947561006dc8621b5597787add0762d Clock:{} Delay:500ms Timeout:10m0s Cancel:<nil>}
	I1002 12:43:58.331820  585741 start.go:364] duration metric: took 43.427µs to acquireMachinesLock for "NoKubernetes-115222"
	I1002 12:43:58.331854  585741 start.go:93] Provisioning new machine with config: &{Name:NoKubernetes-115222 KeepContext:false EmbedCerts:false MinikubeISO: KicBaseImage:gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 Memory:3072 CPUs:2 DiskSize:20000 Driver:docker HyperkitVpnKitSock: HyperkitVSockPorts:[] DockerEnv:[] ContainerVolumeMounts:[] InsecureRegistry:[] RegistryMirror:[] HostOnlyCIDR:192.168.59.1/24 HypervVirtualSwitch: HypervUseExternalSwitch:false HypervExternalAdapter: KVMNetwork:default KVMQemuURI:qemu:///system KVMGPU:false KVMHidden:false KVMNUMACount:1 APIServerPort:8443 DockerOpt:[] DisableDriverMounts:false NFSShare:[] NFSSharesRoot:/nfsshares UUID: NoVTXCheck:false DNSProxy:false HostDNSResolver:true HostOnlyNicType:virtio NatNicType:virtio SSHIPAddress: SSHUser:root SSHKey: SSHPort:22 KubernetesConfig:{KubernetesVersion:v0.0.0 ClusterName:NoKubernetes-115222 Namespace:default APIServerHAVIP: APIServerName:minikubeCA APISe
rverNames:[] APIServerIPs:[] DNSDomain:cluster.local ContainerRuntime:containerd CRISocket: NetworkPlugin:cni FeatureGates: ServiceCIDR:10.96.0.0/12 ImageRepository: LoadBalancerStartIP: LoadBalancerEndIP: CustomIngressCert: RegistryAliases: ExtraOptions:[] ShouldLoadCachedImages:true EnableDefaultCNI:false CNI:} Nodes:[{Name: IP: Port:8443 KubernetesVersion:v0.0.0 ContainerRuntime:containerd ControlPlane:true Worker:true}] Addons:map[] CustomAddonImages:map[] CustomAddonRegistries:map[] VerifyComponents:map[apiserver:true system_pods:true] StartHostTimeout:6m0s ScheduledStop:<nil> ExposedPorts:[] ListenAddress: Network: Subnet: MultiNodeRequested:false ExtraDisks:0 CertExpiration:26280h0m0s MountString: Mount9PVersion:9p2000.L MountGID:docker MountIP: MountMSize:262144 MountOptions:[] MountPort:0 MountType:9p MountUID:docker BinaryMirror: DisableOptimizations:false DisableMetrics:false DisableCoreDNSLog:false CustomQemuFirmwarePath: SocketVMnetClientPath: SocketVMnetPath: StaticIP: SSHAuthSock: SSHAgentPID:0
GPUs: AutoPauseInterval:1m0s} &{Name: IP: Port:8443 KubernetesVersion:v0.0.0 ContainerRuntime:containerd ControlPlane:true Worker:true}
	I1002 12:43:58.331941  585741 start.go:125] createHost starting for "" (driver="docker")
	I1002 12:43:55.053661  583227 out.go:252] * Restarting existing docker container for "stopped-upgrade-250350" ...
	I1002 12:43:55.053739  583227 cli_runner.go:164] Run: docker start stopped-upgrade-250350
	I1002 12:43:55.339095  583227 cli_runner.go:164] Run: docker container inspect stopped-upgrade-250350 --format={{.State.Status}}
	I1002 12:43:55.363276  583227 kic.go:430] container "stopped-upgrade-250350" state is running.
	I1002 12:43:55.363776  583227 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" stopped-upgrade-250350
	I1002 12:43:55.386979  583227 profile.go:143] Saving config to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/stopped-upgrade-250350/config.json ...
	I1002 12:43:55.387263  583227 machine.go:93] provisionDockerMachine start ...
	I1002 12:43:55.387348  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:55.411893  583227 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:55.412281  583227 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33380 <nil> <nil>}
	I1002 12:43:55.412318  583227 main.go:141] libmachine: About to run SSH command:
	hostname
	I1002 12:43:55.413003  583227 main.go:141] libmachine: Error dialing TCP: ssh: handshake failed: read tcp 127.0.0.1:43828->127.0.0.1:33380: read: connection reset by peer
	I1002 12:43:58.531506  583227 main.go:141] libmachine: SSH cmd err, output: <nil>: stopped-upgrade-250350
	
	I1002 12:43:58.531541  583227 ubuntu.go:182] provisioning hostname "stopped-upgrade-250350"
	I1002 12:43:58.531614  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:58.551286  583227 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:58.551592  583227 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33380 <nil> <nil>}
	I1002 12:43:58.551620  583227 main.go:141] libmachine: About to run SSH command:
	sudo hostname stopped-upgrade-250350 && echo "stopped-upgrade-250350" | sudo tee /etc/hostname
	I1002 12:43:58.680783  583227 main.go:141] libmachine: SSH cmd err, output: <nil>: stopped-upgrade-250350
	
	I1002 12:43:58.680891  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:58.700423  583227 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:58.700708  583227 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33380 <nil> <nil>}
	I1002 12:43:58.700734  583227 main.go:141] libmachine: About to run SSH command:
	
			if ! grep -xq '.*\sstopped-upgrade-250350' /etc/hosts; then
				if grep -xq '127.0.1.1\s.*' /etc/hosts; then
					sudo sed -i 's/^127.0.1.1\s.*/127.0.1.1 stopped-upgrade-250350/g' /etc/hosts;
				else 
					echo '127.0.1.1 stopped-upgrade-250350' | sudo tee -a /etc/hosts; 
				fi
			fi
	I1002 12:43:58.824367  583227 main.go:141] libmachine: SSH cmd err, output: <nil>: 
	I1002 12:43:58.824399  583227 ubuntu.go:188] set auth options {CertDir:/home/jenkins/minikube-integration/21139-381342/.minikube CaCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem CaPrivateKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem CaCertRemotePath:/etc/docker/ca.pem ServerCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ServerKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem ClientKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem ServerCertRemotePath:/etc/docker/server.pem ServerKeyRemotePath:/etc/docker/server-key.pem ClientCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem ServerCertSANs:[] StorePath:/home/jenkins/minikube-integration/21139-381342/.minikube}
	I1002 12:43:58.824423  583227 ubuntu.go:190] setting up certificates
	I1002 12:43:58.824433  583227 provision.go:84] configureAuth start
	I1002 12:43:58.824510  583227 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" stopped-upgrade-250350
	I1002 12:43:58.844605  583227 provision.go:143] copyHostCerts
	I1002 12:43:58.844675  583227 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem, removing ...
	I1002 12:43:58.844700  583227 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem
	I1002 12:43:58.844760  583227 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem (1082 bytes)
	I1002 12:43:58.844913  583227 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem, removing ...
	I1002 12:43:58.844927  583227 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem
	I1002 12:43:58.844973  583227 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem (1123 bytes)
	I1002 12:43:58.845069  583227 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem, removing ...
	I1002 12:43:58.845079  583227 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem
	I1002 12:43:58.845118  583227 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem (1675 bytes)
	I1002 12:43:58.845208  583227 provision.go:117] generating server cert: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ca-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem private-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem org=jenkins.stopped-upgrade-250350 san=[127.0.0.1 192.168.103.2 localhost minikube stopped-upgrade-250350]
	I1002 12:43:58.204603  581986 cli_runner.go:164] Run: docker network inspect missing-upgrade-404552 --format "{"Name": "{{.Name}}","Driver": "{{.Driver}}","Subnet": "{{range .IPAM.Config}}{{.Subnet}}{{end}}","Gateway": "{{range .IPAM.Config}}{{.Gateway}}{{end}}","MTU": {{if (index .Options "com.docker.network.driver.mtu")}}{{(index .Options "com.docker.network.driver.mtu")}}{{else}}0{{end}}, "ContainerIPs": [{{range $k,$v := .Containers }}"{{$v.IPv4Address}}",{{end}}]}"
	I1002 12:43:58.223594  581986 ssh_runner.go:195] Run: grep 192.168.76.1	host.minikube.internal$ /etc/hosts
	I1002 12:43:58.227716  581986 ssh_runner.go:195] Run: /bin/bash -c "{ grep -v $'\thost.minikube.internal$' "/etc/hosts"; echo "192.168.76.1	host.minikube.internal"; } > /tmp/h.$$; sudo cp /tmp/h.$$ "/etc/hosts""
	I1002 12:43:58.240322  581986 preload.go:132] Checking if preload exists for k8s version v1.28.3 and runtime containerd
	I1002 12:43:58.240665  581986 ssh_runner.go:195] Run: sudo crictl images --output json
	I1002 12:43:58.277665  581986 containerd.go:604] all images are preloaded for containerd runtime.
	I1002 12:43:58.277685  581986 containerd.go:518] Images already preloaded, skipping extraction
	I1002 12:43:58.277745  581986 ssh_runner.go:195] Run: sudo crictl images --output json
	I1002 12:43:58.313507  581986 containerd.go:604] all images are preloaded for containerd runtime.
	I1002 12:43:58.313522  581986 cache_images.go:84] Images are preloaded, skipping loading
	I1002 12:43:58.313581  581986 ssh_runner.go:195] Run: sudo crictl info
	I1002 12:43:58.352177  581986 cni.go:84] Creating CNI manager for ""
	I1002 12:43:58.352192  581986 cni.go:143] "docker" driver + "containerd" runtime found, recommending kindnet
	I1002 12:43:58.352221  581986 kubeadm.go:87] Using pod CIDR: 10.244.0.0/16
	I1002 12:43:58.352243  581986 kubeadm.go:176] kubeadm options: {CertDir:/var/lib/minikube/certs ServiceCIDR:10.96.0.0/12 PodSubnet:10.244.0.0/16 AdvertiseAddress:192.168.76.2 APIServerPort:8443 KubernetesVersion:v1.28.3 EtcdDataDir:/var/lib/minikube/etcd EtcdExtraArgs:map[] ClusterName:missing-upgrade-404552 NodeName:missing-upgrade-404552 DNSDomain:cluster.local CRISocket:/run/containerd/containerd.sock ImageRepository: ComponentOptions:[{Component:apiServer ExtraArgs:map[enable-admission-plugins:NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota] Pairs:map[certSANs:["127.0.0.1", "localhost", "192.168.76.2"]]} {Component:controllerManager ExtraArgs:map[allocate-node-cidrs:true leader-elect:false] Pairs:map[]} {Component:scheduler ExtraArgs:map[leader-elect:false] Pairs:map[]}] FeatureArgs:map[] NodeIP:192.168.76.2 CgroupDriver:systemd ClientCAFile:/var/lib/minikube/certs/ca.crt S
taticPodPath:/etc/kubernetes/manifests ControlPlaneAddress:control-plane.minikube.internal KubeProxyOptions:map[] ResolvConfSearchRegression:false KubeletConfigOpts:map[hairpinMode:hairpin-veth runtimeRequestTimeout:15m] PrependCriSocketUnix:true}
	I1002 12:43:58.352404  581986 kubeadm.go:181] kubeadm config:
	apiVersion: kubeadm.k8s.io/v1beta3
	kind: InitConfiguration
	localAPIEndpoint:
	  advertiseAddress: 192.168.76.2
	  bindPort: 8443
	bootstrapTokens:
	  - groups:
	      - system:bootstrappers:kubeadm:default-node-token
	    ttl: 24h0m0s
	    usages:
	      - signing
	      - authentication
	nodeRegistration:
	  criSocket: unix:///run/containerd/containerd.sock
	  name: "missing-upgrade-404552"
	  kubeletExtraArgs:
	    node-ip: 192.168.76.2
	  taints: []
	---
	apiVersion: kubeadm.k8s.io/v1beta3
	kind: ClusterConfiguration
	apiServer:
	  certSANs: ["127.0.0.1", "localhost", "192.168.76.2"]
	  extraArgs:
	    enable-admission-plugins: "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota"
	controllerManager:
	  extraArgs:
	    allocate-node-cidrs: "true"
	    leader-elect: "false"
	scheduler:
	  extraArgs:
	    leader-elect: "false"
	certificatesDir: /var/lib/minikube/certs
	clusterName: mk
	controlPlaneEndpoint: control-plane.minikube.internal:8443
	etcd:
	  local:
	    dataDir: /var/lib/minikube/etcd
	    extraArgs:
	      proxy-refresh-interval: "70000"
	kubernetesVersion: v1.28.3
	networking:
	  dnsDomain: cluster.local
	  podSubnet: "10.244.0.0/16"
	  serviceSubnet: 10.96.0.0/12
	---
	apiVersion: kubelet.config.k8s.io/v1beta1
	kind: KubeletConfiguration
	authentication:
	  x509:
	    clientCAFile: /var/lib/minikube/certs/ca.crt
	cgroupDriver: systemd
	hairpinMode: hairpin-veth
	runtimeRequestTimeout: 15m
	clusterDomain: "cluster.local"
	# disable disk resource management by default
	imageGCHighThresholdPercent: 100
	evictionHard:
	  nodefs.available: "0%"
	  nodefs.inodesFree: "0%"
	  imagefs.available: "0%"
	failSwapOn: false
	staticPodPath: /etc/kubernetes/manifests
	---
	apiVersion: kubeproxy.config.k8s.io/v1alpha1
	kind: KubeProxyConfiguration
	clusterCIDR: "10.244.0.0/16"
	metricsBindAddress: 0.0.0.0:10249
	conntrack:
	  maxPerCore: 0
	# Skip setting "net.netfilter.nf_conntrack_tcp_timeout_established"
	  tcpEstablishedTimeout: 0s
	# Skip setting "net.netfilter.nf_conntrack_tcp_timeout_close"
	  tcpCloseWaitTimeout: 0s
	
	I1002 12:43:58.352466  581986 kubeadm.go:976] kubelet [Unit]
	Wants=containerd.service
	
	[Service]
	ExecStart=
	ExecStart=/var/lib/minikube/binaries/v1.28.3/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --config=/var/lib/kubelet/config.yaml --container-runtime-endpoint=unix:///run/containerd/containerd.sock --hostname-override=missing-upgrade-404552 --kubeconfig=/etc/kubernetes/kubelet.conf --node-ip=192.168.76.2
	
	[Install]
	 config:
	{KubernetesVersion:v1.28.3 ClusterName:missing-upgrade-404552 Namespace:default APIServerName:minikubeCA APIServerNames:[] APIServerIPs:[] DNSDomain:cluster.local ContainerRuntime:containerd CRISocket: NetworkPlugin:cni FeatureGates: ServiceCIDR:10.96.0.0/12 ImageRepository: LoadBalancerStartIP: LoadBalancerEndIP: CustomIngressCert: RegistryAliases: ExtraOptions:[] ShouldLoadCachedImages:true EnableDefaultCNI:false CNI: NodeIP: NodePort:8443 NodeName:}
	I1002 12:43:58.352507  581986 ssh_runner.go:195] Run: sudo ls /var/lib/minikube/binaries/v1.28.3
	I1002 12:43:58.361810  581986 binaries.go:44] Found k8s binaries, skipping transfer
	I1002 12:43:58.361897  581986 ssh_runner.go:195] Run: sudo mkdir -p /etc/systemd/system/kubelet.service.d /lib/systemd/system /var/tmp/minikube
	I1002 12:43:58.371174  581986 ssh_runner.go:362] scp memory --> /etc/systemd/system/kubelet.service.d/10-kubeadm.conf (394 bytes)
	I1002 12:43:58.389160  581986 ssh_runner.go:362] scp memory --> /lib/systemd/system/kubelet.service (352 bytes)
	I1002 12:43:58.411246  581986 ssh_runner.go:362] scp memory --> /var/tmp/minikube/kubeadm.yaml.new (2110 bytes)
	I1002 12:43:58.431430  581986 ssh_runner.go:195] Run: grep 192.168.76.2	control-plane.minikube.internal$ /etc/hosts
	I1002 12:43:58.435174  581986 ssh_runner.go:195] Run: /bin/bash -c "{ grep -v $'\tcontrol-plane.minikube.internal$' "/etc/hosts"; echo "192.168.76.2	control-plane.minikube.internal"; } > /tmp/h.$$; sudo cp /tmp/h.$$ "/etc/hosts""
	I1002 12:43:58.446345  581986 certs.go:56] Setting up /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552 for IP: 192.168.76.2
	I1002 12:43:58.446370  581986 certs.go:190] acquiring lock for shared ca certs: {Name:mk0b4dee9533eb55e655fb5dad1d990d151f0d2d Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.446509  581986 certs.go:199] skipping minikubeCA CA generation: /home/jenkins/minikube-integration/21139-381342/.minikube/ca.key
	I1002 12:43:58.446557  581986 certs.go:199] skipping proxyClientCA CA generation: /home/jenkins/minikube-integration/21139-381342/.minikube/proxy-client-ca.key
	I1002 12:43:58.446607  581986 certs.go:319] generating minikube-user signed cert: /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.key
	I1002 12:43:58.446615  581986 crypto.go:68] Generating cert /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.crt with IP's: []
	I1002 12:43:58.560294  581986 crypto.go:156] Writing cert to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.crt ...
	I1002 12:43:58.560321  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.crt: {Name:mk4b5f44166e00d4bcde8b8ec27b665d690df8bf Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.560489  581986 crypto.go:164] Writing key to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.key ...
	I1002 12:43:58.560505  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/client.key: {Name:mkb898b47df603bf78782059d8a04e6f09d8784d Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.560627  581986 certs.go:319] generating minikube signed cert: /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key.31bdca25
	I1002 12:43:58.560638  581986 crypto.go:68] Generating cert /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt.31bdca25 with IP's: [192.168.76.2 10.96.0.1 127.0.0.1 10.0.0.1]
	I1002 12:43:58.650268  581986 crypto.go:156] Writing cert to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt.31bdca25 ...
	I1002 12:43:58.650284  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt.31bdca25: {Name:mk433035fa83eca243b87bc5424c9d7571a6132d Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.650416  581986 crypto.go:164] Writing key to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key.31bdca25 ...
	I1002 12:43:58.650425  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key.31bdca25: {Name:mk57c1c8dfb5d302d0dbc40c6e11db8c54cb244e Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.650490  581986 certs.go:337] copying /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt.31bdca25 -> /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt
	I1002 12:43:58.650569  581986 certs.go:341] copying /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key.31bdca25 -> /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key
	I1002 12:43:58.650622  581986 certs.go:319] generating aggregator signed cert: /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.key
	I1002 12:43:58.650636  581986 crypto.go:68] Generating cert /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.crt with IP's: []
	I1002 12:43:58.956632  581986 crypto.go:156] Writing cert to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.crt ...
	I1002 12:43:58.956649  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.crt: {Name:mk2e1384cf3a50706208d42e4648a292f00d76c3 Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.956803  581986 crypto.go:164] Writing key to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.key ...
	I1002 12:43:58.956814  581986 lock.go:35] WriteFile acquiring /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.key: {Name:mka6a0a0cf42a7c2aae8a0a46864b4a18f560a3f Clock:{} Delay:500ms Timeout:1m0s Cancel:<nil>}
	I1002 12:43:58.957036  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/384955.pem (1338 bytes)
	W1002 12:43:58.957070  581986 certs.go:433] ignoring /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/384955_empty.pem, impossibly tiny 0 bytes
	I1002 12:43:58.957078  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem (1675 bytes)
	I1002 12:43:58.957104  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem (1082 bytes)
	I1002 12:43:58.957124  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem (1123 bytes)
	I1002 12:43:58.957144  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem (1675 bytes)
	I1002 12:43:58.957179  581986 certs.go:437] found cert: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem (1708 bytes)
	I1002 12:43:58.957893  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.crt --> /var/lib/minikube/certs/apiserver.crt (1399 bytes)
	I1002 12:43:58.987145  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/apiserver.key --> /var/lib/minikube/certs/apiserver.key (1679 bytes)
	I1002 12:43:59.011640  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.crt --> /var/lib/minikube/certs/proxy-client.crt (1147 bytes)
	I1002 12:43:59.043029  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/missing-upgrade-404552/proxy-client.key --> /var/lib/minikube/certs/proxy-client.key (1679 bytes)
	I1002 12:43:59.071607  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/ca.crt --> /var/lib/minikube/certs/ca.crt (1111 bytes)
	I1002 12:43:59.095951  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/ca.key --> /var/lib/minikube/certs/ca.key (1679 bytes)
	I1002 12:43:59.120454  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/proxy-client-ca.crt --> /var/lib/minikube/certs/proxy-client-ca.crt (1119 bytes)
	I1002 12:43:59.144500  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/proxy-client-ca.key --> /var/lib/minikube/certs/proxy-client-ca.key (1679 bytes)
	I1002 12:43:59.171360  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem --> /usr/share/ca-certificates/3849552.pem (1708 bytes)
	I1002 12:43:59.201181  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/ca.crt --> /usr/share/ca-certificates/minikubeCA.pem (1111 bytes)
	I1002 12:43:59.228983  581986 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/certs/384955.pem --> /usr/share/ca-certificates/384955.pem (1338 bytes)
	I1002 12:43:59.255173  581986 ssh_runner.go:362] scp memory --> /var/lib/minikube/kubeconfig (738 bytes)
	I1002 12:43:59.274224  581986 ssh_runner.go:195] Run: openssl version
	I1002 12:43:59.281002  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -s /usr/share/ca-certificates/384955.pem && ln -fs /usr/share/ca-certificates/384955.pem /etc/ssl/certs/384955.pem"
	I1002 12:43:59.291291  581986 ssh_runner.go:195] Run: ls -la /usr/share/ca-certificates/384955.pem
	I1002 12:43:59.294937  581986 certs.go:480] hashing: -rw-r--r-- 1 root root 1338 Oct  2 12:20 /usr/share/ca-certificates/384955.pem
	I1002 12:43:59.294983  581986 ssh_runner.go:195] Run: openssl x509 -hash -noout -in /usr/share/ca-certificates/384955.pem
	I1002 12:43:59.301811  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -L /etc/ssl/certs/51391683.0 || ln -fs /etc/ssl/certs/384955.pem /etc/ssl/certs/51391683.0"
	I1002 12:43:59.315896  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -s /usr/share/ca-certificates/3849552.pem && ln -fs /usr/share/ca-certificates/3849552.pem /etc/ssl/certs/3849552.pem"
	I1002 12:43:59.326590  581986 ssh_runner.go:195] Run: ls -la /usr/share/ca-certificates/3849552.pem
	I1002 12:43:59.330767  581986 certs.go:480] hashing: -rw-r--r-- 1 root root 1708 Oct  2 12:20 /usr/share/ca-certificates/3849552.pem
	I1002 12:43:59.330814  581986 ssh_runner.go:195] Run: openssl x509 -hash -noout -in /usr/share/ca-certificates/3849552.pem
	I1002 12:43:59.339545  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -L /etc/ssl/certs/3ec20f2e.0 || ln -fs /etc/ssl/certs/3849552.pem /etc/ssl/certs/3ec20f2e.0"
	I1002 12:43:59.350676  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -s /usr/share/ca-certificates/minikubeCA.pem && ln -fs /usr/share/ca-certificates/minikubeCA.pem /etc/ssl/certs/minikubeCA.pem"
	I1002 12:43:59.361202  581986 ssh_runner.go:195] Run: ls -la /usr/share/ca-certificates/minikubeCA.pem
	I1002 12:43:59.365896  581986 certs.go:480] hashing: -rw-r--r-- 1 root root 1111 Oct  2 12:15 /usr/share/ca-certificates/minikubeCA.pem
	I1002 12:43:59.365958  581986 ssh_runner.go:195] Run: openssl x509 -hash -noout -in /usr/share/ca-certificates/minikubeCA.pem
	I1002 12:43:59.375584  581986 ssh_runner.go:195] Run: sudo /bin/bash -c "test -L /etc/ssl/certs/b5213941.0 || ln -fs /etc/ssl/certs/minikubeCA.pem /etc/ssl/certs/b5213941.0"
	I1002 12:43:59.388988  581986 ssh_runner.go:195] Run: ls /var/lib/minikube/certs/etcd
	I1002 12:43:59.393371  581986 certs.go:353] certs directory doesn't exist, likely first start: ls /var/lib/minikube/certs/etcd: Process exited with status 2
	stdout:
	
	stderr:
	ls: cannot access '/var/lib/minikube/certs/etcd': No such file or directory
	I1002 12:43:59.393422  581986 kubeadm.go:404] StartCluster: {Name:missing-upgrade-404552 KeepContext:false EmbedCerts:false MinikubeISO: KicBaseImage:gcr.io/k8s-minikube/kicbase:v0.0.42@sha256:d35ac07dfda971cabee05e0deca8aeac772f885a5348e1a0c0b0a36db20fcfc0 Memory:3072 CPUs:2 DiskSize:20000 VMDriver: Driver:docker HyperkitVpnKitSock: HyperkitVSockPorts:[] DockerEnv:[] ContainerVolumeMounts:[] InsecureRegistry:[] RegistryMirror:[] HostOnlyCIDR:192.168.59.1/24 HypervVirtualSwitch: HypervUseExternalSwitch:false HypervExternalAdapter: KVMNetwork:default KVMQemuURI:qemu:///system KVMGPU:false KVMHidden:false KVMNUMACount:1 APIServerPort:0 DockerOpt:[] DisableDriverMounts:false NFSShare:[] NFSSharesRoot:/nfsshares UUID: NoVTXCheck:false DNSProxy:false HostDNSResolver:true HostOnlyNicType:virtio NatNicType:virtio SSHIPAddress: SSHUser:root SSHKey: SSHPort:22 KubernetesConfig:{KubernetesVersion:v1.28.3 ClusterName:missing-upgrade-404552 Namespace:default APIServerName:minikubeCA APIServerNames:[] APIServerIP
s:[] DNSDomain:cluster.local ContainerRuntime:containerd CRISocket: NetworkPlugin:cni FeatureGates: ServiceCIDR:10.96.0.0/12 ImageRepository: LoadBalancerStartIP: LoadBalancerEndIP: CustomIngressCert: RegistryAliases: ExtraOptions:[] ShouldLoadCachedImages:true EnableDefaultCNI:false CNI: NodeIP: NodePort:8443 NodeName:} Nodes:[{Name: IP:192.168.76.2 Port:8443 KubernetesVersion:v1.28.3 ContainerRuntime:containerd ControlPlane:true Worker:true}] Addons:map[] CustomAddonImages:map[] CustomAddonRegistries:map[] VerifyComponents:map[apiserver:true system_pods:true] StartHostTimeout:6m0s ScheduledStop:<nil> ExposedPorts:[] ListenAddress: Network: Subnet: MultiNodeRequested:false ExtraDisks:0 CertExpiration:26280h0m0s Mount:false MountString:/home/jenkins:/minikube-host Mount9PVersion:9p2000.L MountGID:docker MountIP: MountMSize:262144 MountOptions:[] MountPort:0 MountType:9p MountUID:docker BinaryMirror: DisableOptimizations:false DisableMetrics:false CustomQemuFirmwarePath: SocketVMnetClientPath: SocketVMnetPath:
StaticIP: SSHAuthSock: SSHAgentPID:0 AutoPauseInterval:1m0s GPUs:}
	I1002 12:43:59.393503  581986 cri.go:54] listing CRI containers in root /run/containerd/runc/k8s.io: {State:paused Name: Namespaces:[kube-system]}
	I1002 12:43:59.393542  581986 ssh_runner.go:195] Run: sudo -s eval "crictl ps -a --quiet --label io.kubernetes.pod.namespace=kube-system"
	I1002 12:43:59.451021  581986 cri.go:89] found id: ""
	I1002 12:43:59.451105  581986 ssh_runner.go:195] Run: sudo ls /var/lib/kubelet/kubeadm-flags.env /var/lib/kubelet/config.yaml /var/lib/minikube/etcd
	I1002 12:43:59.465493  581986 ssh_runner.go:195] Run: sudo cp /var/tmp/minikube/kubeadm.yaml.new /var/tmp/minikube/kubeadm.yaml
	I1002 12:43:59.476153  581986 kubeadm.go:226] ignoring SystemVerification for kubeadm because of docker driver
	I1002 12:43:59.476209  581986 ssh_runner.go:195] Run: sudo ls -la /etc/kubernetes/admin.conf /etc/kubernetes/kubelet.conf /etc/kubernetes/controller-manager.conf /etc/kubernetes/scheduler.conf
	I1002 12:43:59.486438  581986 kubeadm.go:152] config check failed, skipping stale config cleanup: sudo ls -la /etc/kubernetes/admin.conf /etc/kubernetes/kubelet.conf /etc/kubernetes/controller-manager.conf /etc/kubernetes/scheduler.conf: Process exited with status 2
	stdout:
	
	stderr:
	ls: cannot access '/etc/kubernetes/admin.conf': No such file or directory
	ls: cannot access '/etc/kubernetes/kubelet.conf': No such file or directory
	ls: cannot access '/etc/kubernetes/controller-manager.conf': No such file or directory
	ls: cannot access '/etc/kubernetes/scheduler.conf': No such file or directory
	I1002 12:43:59.486490  581986 ssh_runner.go:286] Start: /bin/bash -c "sudo env PATH="/var/lib/minikube/binaries/v1.28.3:$PATH" kubeadm init --config /var/tmp/minikube/kubeadm.yaml  --ignore-preflight-errors=DirAvailable--etc-kubernetes-manifests,DirAvailable--var-lib-minikube,DirAvailable--var-lib-minikube-etcd,FileAvailable--etc-kubernetes-manifests-kube-scheduler.yaml,FileAvailable--etc-kubernetes-manifests-kube-apiserver.yaml,FileAvailable--etc-kubernetes-manifests-kube-controller-manager.yaml,FileAvailable--etc-kubernetes-manifests-etcd.yaml,Port-10250,Swap,NumCPU,Mem,SystemVerification,FileContent--proc-sys-net-bridge-bridge-nf-call-iptables"
	I1002 12:43:59.551520  581986 kubeadm.go:322] [init] Using Kubernetes version: v1.28.3
	I1002 12:43:59.551643  581986 kubeadm.go:322] [preflight] Running pre-flight checks
	I1002 12:43:59.591707  581986 kubeadm.go:322] [preflight] The system verification failed. Printing the output from the verification:
	I1002 12:43:59.591787  581986 kubeadm.go:322] [0;37mKERNEL_VERSION[0m: [0;32m6.8.0-1041-gcp[0m
	I1002 12:43:59.591870  581986 kubeadm.go:322] [0;37mOS[0m: [0;32mLinux[0m
	I1002 12:43:59.591928  581986 kubeadm.go:322] [0;37mCGROUPS_CPU[0m: [0;32menabled[0m
	I1002 12:43:59.591995  581986 kubeadm.go:322] [0;37mCGROUPS_CPUSET[0m: [0;32menabled[0m
	I1002 12:43:59.592058  581986 kubeadm.go:322] [0;37mCGROUPS_DEVICES[0m: [0;32menabled[0m
	I1002 12:43:59.592121  581986 kubeadm.go:322] [0;37mCGROUPS_FREEZER[0m: [0;32menabled[0m
	I1002 12:43:59.592183  581986 kubeadm.go:322] [0;37mCGROUPS_MEMORY[0m: [0;32menabled[0m
	I1002 12:43:59.592259  581986 kubeadm.go:322] [0;37mCGROUPS_PIDS[0m: [0;32menabled[0m
	I1002 12:43:59.592332  581986 kubeadm.go:322] [0;37mCGROUPS_HUGETLB[0m: [0;32menabled[0m
	I1002 12:43:59.592396  581986 kubeadm.go:322] [0;37mCGROUPS_IO[0m: [0;32menabled[0m
	I1002 12:43:59.679277  581986 kubeadm.go:322] [preflight] Pulling images required for setting up a Kubernetes cluster
	I1002 12:43:59.679443  581986 kubeadm.go:322] [preflight] This might take a minute or two, depending on the speed of your internet connection
	I1002 12:43:59.679569  581986 kubeadm.go:322] [preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
	I1002 12:43:59.903212  581986 kubeadm.go:322] [certs] Using certificateDir folder "/var/lib/minikube/certs"
	I1002 12:43:59.136103  583227 provision.go:177] copyRemoteCerts
	I1002 12:43:59.136189  583227 ssh_runner.go:195] Run: sudo mkdir -p /etc/docker /etc/docker /etc/docker
	I1002 12:43:59.136246  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.156567  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	I1002 12:43:59.246326  583227 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /etc/docker/ca.pem (1082 bytes)
	I1002 12:43:59.274495  583227 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem --> /etc/docker/server.pem (1233 bytes)
	I1002 12:43:59.300439  583227 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem --> /etc/docker/server-key.pem (1679 bytes)
	I1002 12:43:59.330525  583227 provision.go:87] duration metric: took 506.064105ms to configureAuth
	I1002 12:43:59.330574  583227 ubuntu.go:206] setting minikube options for container-runtime
	I1002 12:43:59.330764  583227 config.go:182] Loaded profile config "stopped-upgrade-250350": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v1.28.3
	I1002 12:43:59.330785  583227 machine.go:96] duration metric: took 3.943503201s to provisionDockerMachine
	I1002 12:43:59.330798  583227 start.go:293] postStartSetup for "stopped-upgrade-250350" (driver="docker")
	I1002 12:43:59.330812  583227 start.go:322] creating required directories: [/etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs]
	I1002 12:43:59.330941  583227 ssh_runner.go:195] Run: sudo mkdir -p /etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs
	I1002 12:43:59.331092  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.351859  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	I1002 12:43:59.448934  583227 ssh_runner.go:195] Run: cat /etc/os-release
	I1002 12:43:59.456107  583227 main.go:141] libmachine: Couldn't set key VERSION_CODENAME, no corresponding struct field found
	I1002 12:43:59.456147  583227 main.go:141] libmachine: Couldn't set key PRIVACY_POLICY_URL, no corresponding struct field found
	I1002 12:43:59.456232  583227 main.go:141] libmachine: Couldn't set key UBUNTU_CODENAME, no corresponding struct field found
	I1002 12:43:59.456282  583227 info.go:137] Remote host: Ubuntu 22.04.3 LTS
	I1002 12:43:59.456298  583227 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/addons for local assets ...
	I1002 12:43:59.456371  583227 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/files for local assets ...
	I1002 12:43:59.456465  583227 filesync.go:149] local asset: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem -> 3849552.pem in /etc/ssl/certs
	I1002 12:43:59.456602  583227 ssh_runner.go:195] Run: sudo mkdir -p /etc/ssl/certs
	I1002 12:43:59.469387  583227 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem --> /etc/ssl/certs/3849552.pem (1708 bytes)
	I1002 12:43:59.498329  583227 start.go:296] duration metric: took 167.514063ms for postStartSetup
	I1002 12:43:59.498398  583227 ssh_runner.go:195] Run: sh -c "df -h /var | awk 'NR==2{print $5}'"
	I1002 12:43:59.498445  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.524891  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	I1002 12:43:59.614708  583227 ssh_runner.go:195] Run: sh -c "df -BG /var | awk 'NR==2{print $4}'"
	I1002 12:43:59.620747  583227 fix.go:56] duration metric: took 4.606941327s for fixHost
	I1002 12:43:59.620768  583227 start.go:83] releasing machines lock for "stopped-upgrade-250350", held for 4.606988432s
	I1002 12:43:59.620860  583227 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" stopped-upgrade-250350
	I1002 12:43:59.645041  583227 ssh_runner.go:195] Run: cat /version.json
	I1002 12:43:59.645101  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.645135  583227 ssh_runner.go:195] Run: curl -sS -m 2 https://registry.k8s.io/
	I1002 12:43:59.645202  583227 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" stopped-upgrade-250350
	I1002 12:43:59.668811  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	I1002 12:43:59.668811  583227 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33380 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/stopped-upgrade-250350/id_rsa Username:docker}
	W1002 12:43:59.878768  583227 out.go:285] ! Image was not built for the current minikube version. To resolve this you can delete and recreate your minikube cluster using the latest images. Expected minikube version: v1.32.0 -> Actual minikube version: v1.37.0
	I1002 12:43:59.878919  583227 ssh_runner.go:195] Run: systemctl --version
	I1002 12:43:59.883939  583227 ssh_runner.go:195] Run: sh -c "stat /etc/cni/net.d/*loopback.conf*"
	I1002 12:43:59.888632  583227 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f -name *loopback.conf* -not -name *.mk_disabled -exec sh -c "grep -q loopback {} && ( grep -q name {} || sudo sed -i '/"type": "loopback"/i \ \ \ \ "name": "loopback",' {} ) && sudo sed -i 's|"cniVersion": ".*"|"cniVersion": "1.0.0"|g' {}" ;
	I1002 12:43:59.909328  583227 cni.go:230] loopback cni configuration patched: "/etc/cni/net.d/*loopback.conf*" found
	I1002 12:43:59.909394  583227 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f ( ( -name *bridge* -or -name *podman* ) -and -not -name *.mk_disabled ) -printf "%p, " -exec sh -c "sudo mv {} {}.mk_disabled" ;
	I1002 12:43:59.918213  583227 cni.go:259] no active bridge cni configs found in "/etc/cni/net.d" - nothing to disable
	I1002 12:43:59.918236  583227 start.go:495] detecting cgroup driver to use...
	I1002 12:43:59.918270  583227 detect.go:190] detected "systemd" cgroup driver on host os
	I1002 12:43:59.918309  583227 ssh_runner.go:195] Run: sudo systemctl stop -f crio
	I1002 12:43:59.931237  583227 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service crio
	I1002 12:43:59.942383  583227 docker.go:218] disabling cri-docker service (if available) ...
	I1002 12:43:59.942437  583227 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.socket
	I1002 12:43:59.954435  583227 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.service
	I1002 12:43:59.965821  583227 ssh_runner.go:195] Run: sudo systemctl disable cri-docker.socket
	I1002 12:44:00.032939  583227 ssh_runner.go:195] Run: sudo systemctl mask cri-docker.service
	I1002 12:44:00.097207  583227 docker.go:234] disabling docker service ...
	I1002 12:44:00.097264  583227 ssh_runner.go:195] Run: sudo systemctl stop -f docker.socket
	I1002 12:44:00.109382  583227 ssh_runner.go:195] Run: sudo systemctl stop -f docker.service
	I1002 12:44:00.120575  583227 ssh_runner.go:195] Run: sudo systemctl disable docker.socket
	I1002 12:44:00.187758  583227 ssh_runner.go:195] Run: sudo systemctl mask docker.service
	I1002 12:44:00.253215  583227 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service docker
	I1002 12:44:00.264420  583227 ssh_runner.go:195] Run: /bin/bash -c "sudo mkdir -p /etc && printf %s "runtime-endpoint: unix:///run/containerd/containerd.sock
	" | sudo tee /etc/crictl.yaml"
	I1002 12:44:00.282072  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)sandbox_image = .*$|\1sandbox_image = "registry.k8s.io/pause:3.9"|' /etc/containerd/config.toml"
	I1002 12:44:00.291907  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)restrict_oom_score_adj = .*$|\1restrict_oom_score_adj = false|' /etc/containerd/config.toml"
	I1002 12:44:00.301493  583227 containerd.go:146] configuring containerd to use "systemd" as cgroup driver...
	I1002 12:44:00.301553  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)SystemdCgroup = .*$|\1SystemdCgroup = true|g' /etc/containerd/config.toml"
	I1002 12:44:00.311281  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runtime.v1.linux"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:44:00.321045  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i '/systemd_cgroup/d' /etc/containerd/config.toml"
	I1002 12:44:00.330693  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runc.v1"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:44:00.340571  583227 ssh_runner.go:195] Run: sh -c "sudo rm -rf /etc/cni/net.mk"
	I1002 12:44:00.349706  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)conf_dir = .*$|\1conf_dir = "/etc/cni/net.d"|g' /etc/containerd/config.toml"
	I1002 12:44:00.359285  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i '/^ *enable_unprivileged_ports = .*/d' /etc/containerd/config.toml"
	I1002 12:44:00.369202  583227 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)\[plugins."io.containerd.grpc.v1.cri"\]|&\n\1  enable_unprivileged_ports = true|' /etc/containerd/config.toml"
	I1002 12:44:00.379878  583227 ssh_runner.go:195] Run: sudo sysctl net.bridge.bridge-nf-call-iptables
	I1002 12:44:00.388529  583227 ssh_runner.go:195] Run: sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
	I1002 12:44:00.396849  583227 ssh_runner.go:195] Run: sudo systemctl daemon-reload
	I1002 12:44:00.464051  583227 ssh_runner.go:195] Run: sudo systemctl restart containerd
	I1002 12:44:00.576008  583227 start.go:542] Will wait 60s for socket path /run/containerd/containerd.sock
	I1002 12:44:00.576088  583227 ssh_runner.go:195] Run: stat /run/containerd/containerd.sock
	I1002 12:44:00.580326  583227 start.go:563] Will wait 60s for crictl version
	I1002 12:44:00.580379  583227 ssh_runner.go:195] Run: which crictl
	I1002 12:44:00.583808  583227 ssh_runner.go:195] Run: sudo /usr/bin/crictl version
	I1002 12:44:00.618127  583227 start.go:579] Version:  0.1.0
	RuntimeName:  containerd
	RuntimeVersion:  1.6.24
	RuntimeApiVersion:  v1
	I1002 12:44:00.618181  583227 ssh_runner.go:195] Run: containerd --version
	I1002 12:44:00.642732  583227 ssh_runner.go:195] Run: containerd --version
	I1002 12:44:00.670321  583227 out.go:179] * Preparing Kubernetes v1.28.3 on containerd 1.6.24 ...
	I1002 12:43:59.904975  581986 out.go:204]   - Generating certificates and keys ...
	I1002 12:43:59.905096  581986 kubeadm.go:322] [certs] Using existing ca certificate authority
	I1002 12:43:59.905187  581986 kubeadm.go:322] [certs] Using existing apiserver certificate and key on disk
	I1002 12:44:00.230450  581986 kubeadm.go:322] [certs] Generating "apiserver-kubelet-client" certificate and key
	I1002 12:44:00.342809  581986 kubeadm.go:322] [certs] Generating "front-proxy-ca" certificate and key
	I1002 12:44:00.427885  581986 kubeadm.go:322] [certs] Generating "front-proxy-client" certificate and key
	I1002 12:44:00.668257  581986 kubeadm.go:322] [certs] Generating "etcd/ca" certificate and key
	I1002 12:44:00.911902  581986 kubeadm.go:322] [certs] Generating "etcd/server" certificate and key
	I1002 12:44:00.912096  581986 kubeadm.go:322] [certs] etcd/server serving cert is signed for DNS names [localhost missing-upgrade-404552] and IPs [192.168.76.2 127.0.0.1 ::1]
	I1002 12:43:58.333473  585741 out.go:252] * Creating docker container (CPUs=2, Memory=3072MB) ...
	I1002 12:43:58.333729  585741 start.go:159] libmachine.API.Create for "NoKubernetes-115222" (driver="docker")
	I1002 12:43:58.333767  585741 client.go:168] LocalClient.Create starting
	I1002 12:43:58.333864  585741 main.go:141] libmachine: Reading certificate data from /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem
	I1002 12:43:58.333920  585741 main.go:141] libmachine: Decoding PEM data...
	I1002 12:43:58.333951  585741 main.go:141] libmachine: Parsing certificate...
	I1002 12:43:58.334022  585741 main.go:141] libmachine: Reading certificate data from /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem
	I1002 12:43:58.334051  585741 main.go:141] libmachine: Decoding PEM data...
	I1002 12:43:58.334068  585741 main.go:141] libmachine: Parsing certificate...
	I1002 12:43:58.334563  585741 cli_runner.go:164] Run: docker network inspect NoKubernetes-115222 --format "{"Name": "{{.Name}}","Driver": "{{.Driver}}","Subnet": "{{range .IPAM.Config}}{{.Subnet}}{{end}}","Gateway": "{{range .IPAM.Config}}{{.Gateway}}{{end}}","MTU": {{if (index .Options "com.docker.network.driver.mtu")}}{{(index .Options "com.docker.network.driver.mtu")}}{{else}}0{{end}}, "ContainerIPs": [{{range $k,$v := .Containers }}"{{$v.IPv4Address}}",{{end}}]}"
	W1002 12:43:58.352366  585741 cli_runner.go:211] docker network inspect NoKubernetes-115222 --format "{"Name": "{{.Name}}","Driver": "{{.Driver}}","Subnet": "{{range .IPAM.Config}}{{.Subnet}}{{end}}","Gateway": "{{range .IPAM.Config}}{{.Gateway}}{{end}}","MTU": {{if (index .Options "com.docker.network.driver.mtu")}}{{(index .Options "com.docker.network.driver.mtu")}}{{else}}0{{end}}, "ContainerIPs": [{{range $k,$v := .Containers }}"{{$v.IPv4Address}}",{{end}}]}" returned with exit code 1
	I1002 12:43:58.352435  585741 network_create.go:284] running [docker network inspect NoKubernetes-115222] to gather additional debugging logs...
	I1002 12:43:58.352469  585741 cli_runner.go:164] Run: docker network inspect NoKubernetes-115222
	W1002 12:43:58.368779  585741 cli_runner.go:211] docker network inspect NoKubernetes-115222 returned with exit code 1
	I1002 12:43:58.368806  585741 network_create.go:287] error running [docker network inspect NoKubernetes-115222]: docker network inspect NoKubernetes-115222: exit status 1
	stdout:
	[]
	
	stderr:
	Error response from daemon: network NoKubernetes-115222 not found
	I1002 12:43:58.368821  585741 network_create.go:289] output of [docker network inspect NoKubernetes-115222]: -- stdout --
	[]
	
	-- /stdout --
	** stderr ** 
	Error response from daemon: network NoKubernetes-115222 not found
	
	** /stderr **
	I1002 12:43:58.368926  585741 cli_runner.go:164] Run: docker network inspect bridge --format "{"Name": "{{.Name}}","Driver": "{{.Driver}}","Subnet": "{{range .IPAM.Config}}{{.Subnet}}{{end}}","Gateway": "{{range .IPAM.Config}}{{.Gateway}}{{end}}","MTU": {{if (index .Options "com.docker.network.driver.mtu")}}{{(index .Options "com.docker.network.driver.mtu")}}{{else}}0{{end}}, "ContainerIPs": [{{range $k,$v := .Containers }}"{{$v.IPv4Address}}",{{end}}]}"
	I1002 12:43:58.386690  585741 network.go:211] skipping subnet 192.168.49.0/24 that is taken: &{IP:192.168.49.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.49.0/24 Gateway:192.168.49.1 ClientMin:192.168.49.2 ClientMax:192.168.49.254 Broadcast:192.168.49.255 IsPrivate:true Interface:{IfaceName:br-c8fe0fd21bd1 IfaceIPv4:192.168.49.1 IfaceMTU:1500 IfaceMAC:46:4d:f9:f4:cc:90} reservation:<nil>}
	I1002 12:43:58.387729  585741 network.go:211] skipping subnet 192.168.58.0/24 that is taken: &{IP:192.168.58.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.58.0/24 Gateway:192.168.58.1 ClientMin:192.168.58.2 ClientMax:192.168.58.254 Broadcast:192.168.58.255 IsPrivate:true Interface:{IfaceName:br-5cdc9a6b29c5 IfaceIPv4:192.168.58.1 IfaceMTU:1500 IfaceMAC:62:2f:0d:cf:b5:61} reservation:<nil>}
	I1002 12:43:58.388265  585741 network.go:211] skipping subnet 192.168.67.0/24 that is taken: &{IP:192.168.67.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.67.0/24 Gateway:192.168.67.1 ClientMin:192.168.67.2 ClientMax:192.168.67.254 Broadcast:192.168.67.255 IsPrivate:true Interface:{IfaceName:br-cb914d436a66 IfaceIPv4:192.168.67.1 IfaceMTU:1500 IfaceMAC:9a:69:ee:9d:31:3d} reservation:<nil>}
	I1002 12:43:58.388874  585741 network.go:211] skipping subnet 192.168.76.0/24 that is taken: &{IP:192.168.76.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.76.0/24 Gateway:192.168.76.1 ClientMin:192.168.76.2 ClientMax:192.168.76.254 Broadcast:192.168.76.255 IsPrivate:true Interface:{IfaceName:br-f4a55d02f8d3 IfaceIPv4:192.168.76.1 IfaceMTU:1500 IfaceMAC:9a:b9:5b:5d:66:d3} reservation:<nil>}
	I1002 12:43:58.389804  585741 network.go:206] using free private subnet 192.168.85.0/24: &{IP:192.168.85.0 Netmask:255.255.255.0 Prefix:24 CIDR:192.168.85.0/24 Gateway:192.168.85.1 ClientMin:192.168.85.2 ClientMax:192.168.85.254 Broadcast:192.168.85.255 IsPrivate:true Interface:{IfaceName: IfaceIPv4: IfaceMTU:0 IfaceMAC:} reservation:0xc001e9dda0}
	I1002 12:43:58.389845  585741 network_create.go:124] attempt to create docker network NoKubernetes-115222 192.168.85.0/24 with gateway 192.168.85.1 and MTU of 1500 ...
	I1002 12:43:58.389896  585741 cli_runner.go:164] Run: docker network create --driver=bridge --subnet=192.168.85.0/24 --gateway=192.168.85.1 -o --ip-masq -o --icc -o com.docker.network.driver.mtu=1500 --label=created_by.minikube.sigs.k8s.io=true --label=name.minikube.sigs.k8s.io=NoKubernetes-115222 NoKubernetes-115222
	I1002 12:43:58.449771  585741 network_create.go:108] docker network NoKubernetes-115222 192.168.85.0/24 created
	I1002 12:43:58.449818  585741 kic.go:121] calculated static IP "192.168.85.2" for the "NoKubernetes-115222" container
	I1002 12:43:58.449907  585741 cli_runner.go:164] Run: docker ps -a --format {{.Names}}
	I1002 12:43:58.470236  585741 cli_runner.go:164] Run: docker volume create NoKubernetes-115222 --label name.minikube.sigs.k8s.io=NoKubernetes-115222 --label created_by.minikube.sigs.k8s.io=true
	I1002 12:43:58.487719  585741 oci.go:103] Successfully created a docker volume NoKubernetes-115222
	I1002 12:43:58.487792  585741 cli_runner.go:164] Run: docker run --rm --name NoKubernetes-115222-preload-sidecar --label created_by.minikube.sigs.k8s.io=true --label name.minikube.sigs.k8s.io=NoKubernetes-115222 --entrypoint /usr/bin/test -v NoKubernetes-115222:/var gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1 -d /var/lib
	I1002 12:43:58.855992  585741 oci.go:107] Successfully prepared a docker volume NoKubernetes-115222
	I1002 12:43:58.856028  585741 preload.go:178] Skipping preload logic due to --no-kubernetes flag
	W1002 12:43:58.856092  585741 cgroups_linux.go:77] Your kernel does not support swap limit capabilities or the cgroup is not mounted.
	W1002 12:43:58.856120  585741 oci.go:252] Your kernel does not support CPU cfs period/quota or the cgroup is not mounted.
	I1002 12:43:58.856155  585741 cli_runner.go:164] Run: docker info --format "'{{json .SecurityOptions}}'"
	I1002 12:43:58.914050  585741 cli_runner.go:164] Run: docker run -d -t --privileged --security-opt seccomp=unconfined --tmpfs /tmp --tmpfs /run -v /lib/modules:/lib/modules:ro --hostname NoKubernetes-115222 --name NoKubernetes-115222 --label created_by.minikube.sigs.k8s.io=true --label name.minikube.sigs.k8s.io=NoKubernetes-115222 --label role.minikube.sigs.k8s.io= --label mode.minikube.sigs.k8s.io=NoKubernetes-115222 --network NoKubernetes-115222 --ip 192.168.85.2 --volume NoKubernetes-115222:/var --security-opt apparmor=unconfined --memory=3072mb -e container=docker --expose 8443 --publish=127.0.0.1::8443 --publish=127.0.0.1::22 --publish=127.0.0.1::2376 --publish=127.0.0.1::5000 --publish=127.0.0.1::32443 gcr.io/k8s-minikube/kicbase:v0.0.48@sha256:7171c97a51623558720f8e5878e4f4637da093e2f2ed589997bedc6c1549b2b1
	I1002 12:43:59.170328  585741 cli_runner.go:164] Run: docker container inspect NoKubernetes-115222 --format={{.State.Running}}
	I1002 12:43:59.190310  585741 cli_runner.go:164] Run: docker container inspect NoKubernetes-115222 --format={{.State.Status}}
	I1002 12:43:59.210643  585741 cli_runner.go:164] Run: docker exec NoKubernetes-115222 stat /var/lib/dpkg/alternatives/iptables
	I1002 12:43:59.256165  585741 oci.go:144] the created container "NoKubernetes-115222" has a running status.
	I1002 12:43:59.256202  585741 kic.go:225] Creating ssh key for kic: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa...
	I1002 12:43:59.313979  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa.pub -> /home/docker/.ssh/authorized_keys
	I1002 12:43:59.314036  585741 kic_runner.go:191] docker (temp): /home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa.pub --> /home/docker/.ssh/authorized_keys (381 bytes)
	I1002 12:43:59.343661  585741 cli_runner.go:164] Run: docker container inspect NoKubernetes-115222 --format={{.State.Status}}
	I1002 12:43:59.364124  585741 kic_runner.go:93] Run: chown docker:docker /home/docker/.ssh/authorized_keys
	I1002 12:43:59.364145  585741 kic_runner.go:114] Args: [docker exec --privileged NoKubernetes-115222 chown docker:docker /home/docker/.ssh/authorized_keys]
	I1002 12:43:59.411233  585741 cli_runner.go:164] Run: docker container inspect NoKubernetes-115222 --format={{.State.Status}}
	I1002 12:43:59.437345  585741 machine.go:93] provisionDockerMachine start ...
	I1002 12:43:59.437476  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:43:59.465028  585741 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:59.465377  585741 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33390 <nil> <nil>}
	I1002 12:43:59.465401  585741 main.go:141] libmachine: About to run SSH command:
	hostname
	I1002 12:43:59.613979  585741 main.go:141] libmachine: SSH cmd err, output: <nil>: NoKubernetes-115222
	
	I1002 12:43:59.614086  585741 ubuntu.go:182] provisioning hostname "NoKubernetes-115222"
	I1002 12:43:59.614195  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:43:59.637770  585741 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:59.638098  585741 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33390 <nil> <nil>}
	I1002 12:43:59.638114  585741 main.go:141] libmachine: About to run SSH command:
	sudo hostname NoKubernetes-115222 && echo "NoKubernetes-115222" | sudo tee /etc/hostname
	I1002 12:43:59.798996  585741 main.go:141] libmachine: SSH cmd err, output: <nil>: NoKubernetes-115222
	
	I1002 12:43:59.799084  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:43:59.820859  585741 main.go:141] libmachine: Using SSH client type: native
	I1002 12:43:59.821181  585741 main.go:141] libmachine: &{{{<nil> 0 [] [] []} docker [0x840040] 0x842d40 <nil>  [] 0s} 127.0.0.1 33390 <nil> <nil>}
	I1002 12:43:59.821206  585741 main.go:141] libmachine: About to run SSH command:
	
			if ! grep -xq '.*\sNoKubernetes-115222' /etc/hosts; then
				if grep -xq '127.0.1.1\s.*' /etc/hosts; then
					sudo sed -i 's/^127.0.1.1\s.*/127.0.1.1 NoKubernetes-115222/g' /etc/hosts;
				else 
					echo '127.0.1.1 NoKubernetes-115222' | sudo tee -a /etc/hosts; 
				fi
			fi
	I1002 12:43:59.959737  585741 main.go:141] libmachine: SSH cmd err, output: <nil>: 
	I1002 12:43:59.959763  585741 ubuntu.go:188] set auth options {CertDir:/home/jenkins/minikube-integration/21139-381342/.minikube CaCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem CaPrivateKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem CaCertRemotePath:/etc/docker/ca.pem ServerCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ServerKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem ClientKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem ServerCertRemotePath:/etc/docker/server.pem ServerKeyRemotePath:/etc/docker/server-key.pem ClientCertPath:/home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem ServerCertSANs:[] StorePath:/home/jenkins/minikube-integration/21139-381342/.minikube}
	I1002 12:43:59.959785  585741 ubuntu.go:190] setting up certificates
	I1002 12:43:59.959797  585741 provision.go:84] configureAuth start
	I1002 12:43:59.959863  585741 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" NoKubernetes-115222
	I1002 12:43:59.977886  585741 provision.go:143] copyHostCerts
	I1002 12:43:59.977920  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem -> /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem
	I1002 12:43:59.977955  585741 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem, removing ...
	I1002 12:43:59.977972  585741 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem
	I1002 12:43:59.978051  585741 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/ca.pem (1082 bytes)
	I1002 12:43:59.978144  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem -> /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem
	I1002 12:43:59.978171  585741 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem, removing ...
	I1002 12:43:59.978181  585741 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem
	I1002 12:43:59.978225  585741 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/cert.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/cert.pem (1123 bytes)
	I1002 12:43:59.978295  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem -> /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem
	I1002 12:43:59.978321  585741 exec_runner.go:144] found /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem, removing ...
	I1002 12:43:59.978330  585741 exec_runner.go:203] rm: /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem
	I1002 12:43:59.978366  585741 exec_runner.go:151] cp: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/key.pem --> /home/jenkins/minikube-integration/21139-381342/.minikube/key.pem (1675 bytes)
	I1002 12:43:59.978453  585741 provision.go:117] generating server cert: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem ca-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem private-key=/home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca-key.pem org=jenkins.NoKubernetes-115222 san=[127.0.0.1 192.168.85.2 NoKubernetes-115222 localhost minikube]
	I1002 12:44:00.173620  585741 provision.go:177] copyRemoteCerts
	I1002 12:44:00.173691  585741 ssh_runner.go:195] Run: sudo mkdir -p /etc/docker /etc/docker /etc/docker
	I1002 12:44:00.173738  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.192246  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.288330  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem -> /etc/docker/ca.pem
	I1002 12:44:00.288391  585741 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/certs/ca.pem --> /etc/docker/ca.pem (1082 bytes)
	I1002 12:44:00.313987  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem -> /etc/docker/server.pem
	I1002 12:44:00.314044  585741 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server.pem --> /etc/docker/server.pem (1224 bytes)
	I1002 12:44:00.338589  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem -> /etc/docker/server-key.pem
	I1002 12:44:00.338650  585741 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/machines/server-key.pem --> /etc/docker/server-key.pem (1675 bytes)
	I1002 12:44:00.363237  585741 provision.go:87] duration metric: took 403.427992ms to configureAuth
	I1002 12:44:00.363265  585741 ubuntu.go:206] setting minikube options for container-runtime
	I1002 12:44:00.363490  585741 config.go:182] Loaded profile config "NoKubernetes-115222": Driver=docker, ContainerRuntime=containerd, KubernetesVersion=v0.0.0
	I1002 12:44:00.363508  585741 machine.go:96] duration metric: took 926.124193ms to provisionDockerMachine
	I1002 12:44:00.363522  585741 client.go:171] duration metric: took 2.029739003s to LocalClient.Create
	I1002 12:44:00.363546  585741 start.go:167] duration metric: took 2.029818501s to libmachine.API.Create "NoKubernetes-115222"
	I1002 12:44:00.363559  585741 start.go:293] postStartSetup for "NoKubernetes-115222" (driver="docker")
	I1002 12:44:00.363574  585741 start.go:322] creating required directories: [/etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs]
	I1002 12:44:00.363623  585741 ssh_runner.go:195] Run: sudo mkdir -p /etc/kubernetes/addons /etc/kubernetes/manifests /var/tmp/minikube /var/lib/minikube /var/lib/minikube/certs /var/lib/minikube/images /var/lib/minikube/binaries /tmp/gvisor /usr/share/ca-certificates /etc/ssl/certs
	I1002 12:44:00.363670  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.382605  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.481623  585741 ssh_runner.go:195] Run: cat /etc/os-release
	I1002 12:44:00.484862  585741 main.go:141] libmachine: Couldn't set key VERSION_CODENAME, no corresponding struct field found
	I1002 12:44:00.484897  585741 main.go:141] libmachine: Couldn't set key PRIVACY_POLICY_URL, no corresponding struct field found
	I1002 12:44:00.484908  585741 main.go:141] libmachine: Couldn't set key UBUNTU_CODENAME, no corresponding struct field found
	I1002 12:44:00.484919  585741 info.go:137] Remote host: Ubuntu 22.04.5 LTS
	I1002 12:44:00.484930  585741 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/addons for local assets ...
	I1002 12:44:00.484995  585741 filesync.go:126] Scanning /home/jenkins/minikube-integration/21139-381342/.minikube/files for local assets ...
	I1002 12:44:00.485089  585741 filesync.go:149] local asset: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem -> 3849552.pem in /etc/ssl/certs
	I1002 12:44:00.485102  585741 vm_assets.go:164] NewFileAsset: /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem -> /etc/ssl/certs/3849552.pem
	I1002 12:44:00.485223  585741 ssh_runner.go:195] Run: sudo mkdir -p /etc/ssl/certs
	I1002 12:44:00.493785  585741 ssh_runner.go:362] scp /home/jenkins/minikube-integration/21139-381342/.minikube/files/etc/ssl/certs/3849552.pem --> /etc/ssl/certs/3849552.pem (1708 bytes)
	I1002 12:44:00.520356  585741 start.go:296] duration metric: took 156.779587ms for postStartSetup
	I1002 12:44:00.520741  585741 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" NoKubernetes-115222
	I1002 12:44:00.540900  585741 profile.go:143] Saving config to /home/jenkins/minikube-integration/21139-381342/.minikube/profiles/NoKubernetes-115222/config.json ...
	I1002 12:44:00.541121  585741 ssh_runner.go:195] Run: sh -c "df -h /var | awk 'NR==2{print $5}'"
	I1002 12:44:00.541161  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.558900  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.653633  585741 ssh_runner.go:195] Run: sh -c "df -BG /var | awk 'NR==2{print $4}'"
	I1002 12:44:00.658171  585741 start.go:128] duration metric: took 2.32621501s to createHost
	I1002 12:44:00.658198  585741 start.go:83] releasing machines lock for "NoKubernetes-115222", held for 2.326350832s
	I1002 12:44:00.658263  585741 cli_runner.go:164] Run: docker container inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}},{{.GlobalIPv6Address}}{{end}}" NoKubernetes-115222
	I1002 12:44:00.676947  585741 ssh_runner.go:195] Run: cat /version.json
	I1002 12:44:00.676960  585741 ssh_runner.go:195] Run: curl -sS -m 2 https://registry.k8s.io/
	I1002 12:44:00.676998  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.677032  585741 cli_runner.go:164] Run: docker container inspect -f "'{{(index (index .NetworkSettings.Ports "22/tcp") 0).HostPort}}'" NoKubernetes-115222
	I1002 12:44:00.696896  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.697291  585741 sshutil.go:53] new ssh client: &{IP:127.0.0.1 Port:33390 SSHKeyPath:/home/jenkins/minikube-integration/21139-381342/.minikube/machines/NoKubernetes-115222/id_rsa Username:docker}
	I1002 12:44:00.875633  585741 ssh_runner.go:195] Run: systemctl --version
	I1002 12:44:00.880342  585741 ssh_runner.go:195] Run: sh -c "stat /etc/cni/net.d/*loopback.conf*"
	I1002 12:44:00.885039  585741 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f -name *loopback.conf* -not -name *.mk_disabled -exec sh -c "grep -q loopback {} && ( grep -q name {} || sudo sed -i '/"type": "loopback"/i \ \ \ \ "name": "loopback",' {} ) && sudo sed -i 's|"cniVersion": ".*"|"cniVersion": "1.0.0"|g' {}" ;
	I1002 12:44:00.913715  585741 cni.go:230] loopback cni configuration patched: "/etc/cni/net.d/*loopback.conf*" found
	I1002 12:44:00.913787  585741 ssh_runner.go:195] Run: sudo find /etc/cni/net.d -maxdepth 1 -type f ( ( -name *bridge* -or -name *podman* ) -and -not -name *.mk_disabled ) -printf "%p, " -exec sh -c "sudo mv {} {}.mk_disabled" ;
	I1002 12:44:00.943890  585741 cni.go:262] disabled [/etc/cni/net.d/87-podman-bridge.conflist, /etc/cni/net.d/100-crio-bridge.conf] bridge cni config(s)
	I1002 12:44:00.943918  585741 start.go:495] detecting cgroup driver to use...
	I1002 12:44:00.943951  585741 detect.go:190] detected "systemd" cgroup driver on host os
	I1002 12:44:00.943999  585741 ssh_runner.go:195] Run: sudo systemctl stop -f crio
	I1002 12:44:00.956343  585741 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service crio
	I1002 12:44:00.967470  585741 docker.go:218] disabling cri-docker service (if available) ...
	I1002 12:44:00.967519  585741 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.socket
	I1002 12:44:00.980623  585741 ssh_runner.go:195] Run: sudo systemctl stop -f cri-docker.service
	I1002 12:44:00.994982  585741 ssh_runner.go:195] Run: sudo systemctl disable cri-docker.socket
	I1002 12:44:01.077417  585741 ssh_runner.go:195] Run: sudo systemctl mask cri-docker.service
	I1002 12:44:01.161042  585741 docker.go:234] disabling docker service ...
	I1002 12:44:01.161112  585741 ssh_runner.go:195] Run: sudo systemctl stop -f docker.socket
	I1002 12:44:01.179656  585741 ssh_runner.go:195] Run: sudo systemctl stop -f docker.service
	I1002 12:44:01.191648  585741 ssh_runner.go:195] Run: sudo systemctl disable docker.socket
	I1002 12:44:01.265615  585741 ssh_runner.go:195] Run: sudo systemctl mask docker.service
	I1002 12:44:01.334379  585741 ssh_runner.go:195] Run: sudo systemctl is-active --quiet service docker
	I1002 12:44:01.346062  585741 ssh_runner.go:195] Run: /bin/bash -c "sudo mkdir -p /etc && printf %s "runtime-endpoint: unix:///run/containerd/containerd.sock
	" | sudo tee /etc/crictl.yaml"
	I1002 12:44:01.362809  585741 binary.go:59] Skipping Kubernetes binary download due to --no-kubernetes flag
	I1002 12:44:01.362921  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)sandbox_image = .*$|\1sandbox_image = "registry.k8s.io/pause:3.9"|' /etc/containerd/config.toml"
	I1002 12:44:01.374360  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)restrict_oom_score_adj = .*$|\1restrict_oom_score_adj = false|' /etc/containerd/config.toml"
	I1002 12:44:01.384353  585741 containerd.go:146] configuring containerd to use "systemd" as cgroup driver...
	I1002 12:44:01.384409  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)SystemdCgroup = .*$|\1SystemdCgroup = true|g' /etc/containerd/config.toml"
	I1002 12:44:01.394457  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runtime.v1.linux"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:44:01.404615  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i '/systemd_cgroup/d' /etc/containerd/config.toml"
	I1002 12:44:01.415284  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i 's|"io.containerd.runc.v1"|"io.containerd.runc.v2"|g' /etc/containerd/config.toml"
	I1002 12:44:01.425249  585741 ssh_runner.go:195] Run: sh -c "sudo rm -rf /etc/cni/net.mk"
	I1002 12:44:01.434805  585741 ssh_runner.go:195] Run: sh -c "sudo sed -i -r 's|^( *)conf_dir = .*$|\1conf_dir = "/etc/cni/net.d"|g' /etc/containerd/config.toml"
	I1002 12:44:01.445128  585741 ssh_runner.go:195] Run: sudo sysctl net.bridge.bridge-nf-call-iptables
	I1002 12:44:01.453731  585741 ssh_runner.go:195] Run: sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
	I1002 12:44:01.462601  585741 ssh_runner.go:195] Run: sudo systemctl daemon-reload
	I1002 12:44:01.560410  585741 ssh_runner.go:195] Run: sudo systemctl restart containerd
	I1002 12:44:01.669135  585741 start.go:542] Will wait 60s for socket path /run/containerd/containerd.sock
	I1002 12:44:01.669211  585741 ssh_runner.go:195] Run: stat /run/containerd/containerd.sock
	I1002 12:44:01.675844  585741 start.go:563] Will wait 60s for crictl version
	I1002 12:44:01.675914  585741 ssh_runner.go:195] Run: which crictl
	I1002 12:44:01.680923  585741 ssh_runner.go:195] Run: sudo /usr/bin/crictl version
	I1002 12:44:01.737546  585741 start.go:579] Version:  0.1.0
	RuntimeName:  containerd
	RuntimeVersion:  1.7.27
	RuntimeApiVersion:  v1
	I1002 12:44:01.737615  585741 ssh_runner.go:195] Run: containerd --version
	I1002 12:44:01.776697  585741 ssh_runner.go:195] Run: containerd --version
	I1002 12:44:01.812218  585741 out.go:179] * Preparing containerd 1.7.27 ...
	I1002 12:44:01.813955  585741 ssh_runner.go:195] Run: rm -f paused
	I1002 12:44:01.820058  585741 out.go:179] * Done! minikube is ready without Kubernetes!
	I1002 12:44:01.823078  585741 out.go:203] ╭──────────────────────────────────────────────────────────╮
	│                                                          │
	│          * Things to try without Kubernetes ...          │
	│                                                          │
	│    - "minikube ssh" to SSH into minikube's node.         │
	│    - "minikube image" to build images without docker.    │
	│                                                          │
	╰──────────────────────────────────────────────────────────╯
	I1002 12:44:01.108666  581986 kubeadm.go:322] [certs] Generating "etcd/peer" certificate and key
	I1002 12:44:01.108847  581986 kubeadm.go:322] [certs] etcd/peer serving cert is signed for DNS names [localhost missing-upgrade-404552] and IPs [192.168.76.2 127.0.0.1 ::1]
	I1002 12:44:01.220230  581986 kubeadm.go:322] [certs] Generating "etcd/healthcheck-client" certificate and key
	I1002 12:44:01.449689  581986 kubeadm.go:322] [certs] Generating "apiserver-etcd-client" certificate and key
	I1002 12:44:01.853921  581986 kubeadm.go:322] [certs] Generating "sa" key and public key
	I1002 12:44:01.854003  581986 kubeadm.go:322] [kubeconfig] Using kubeconfig folder "/etc/kubernetes"
	I1002 12:44:01.996425  581986 kubeadm.go:322] [kubeconfig] Writing "admin.conf" kubeconfig file
	I1002 12:44:02.205313  581986 kubeadm.go:322] [kubeconfig] Writing "kubelet.conf" kubeconfig file
	I1002 12:44:02.404845  581986 kubeadm.go:322] [kubeconfig] Writing "controller-manager.conf" kubeconfig file
	I1002 12:44:02.544713  581986 kubeadm.go:322] [kubeconfig] Writing "scheduler.conf" kubeconfig file
	I1002 12:44:02.545301  581986 kubeadm.go:322] [etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
	I1002 12:44:02.549303  581986 kubeadm.go:322] [control-plane] Using manifest folder "/etc/kubernetes/manifests"
	
	
	==> container status <==
	CONTAINER           IMAGE               CREATED             STATE               NAME                ATTEMPT             POD ID              POD
	
	
	==> containerd <==
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663216733Z" level=info msg="loading plugin \"io.containerd.grpc.v1.version\"..." type=io.containerd.grpc.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663279467Z" level=info msg="loading plugin \"io.containerd.internal.v1.restart\"..." type=io.containerd.internal.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663417335Z" level=info msg="loading plugin \"io.containerd.tracing.processor.v1.otlp\"..." type=io.containerd.tracing.processor.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663498028Z" level=info msg="skip loading plugin \"io.containerd.tracing.processor.v1.otlp\"..." error="skip plugin: tracing endpoint not configured" type=io.containerd.tracing.processor.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663558306Z" level=info msg="loading plugin \"io.containerd.internal.v1.tracing\"..." type=io.containerd.internal.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663654570Z" level=info msg="skip loading plugin \"io.containerd.internal.v1.tracing\"..." error="skip plugin: tracing endpoint not configured" type=io.containerd.internal.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663715445Z" level=info msg="loading plugin \"io.containerd.grpc.v1.healthcheck\"..." type=io.containerd.grpc.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663812713Z" level=info msg="loading plugin \"io.containerd.nri.v1.nri\"..." type=io.containerd.nri.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663908312Z" level=info msg="NRI interface is disabled by configuration."
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.663975835Z" level=info msg="loading plugin \"io.containerd.grpc.v1.cri\"..." type=io.containerd.grpc.v1
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.664550919Z" level=info msg="Start cri plugin with config {PluginConfig:{ContainerdConfig:{Snapshotter:overlayfs DefaultRuntimeName:runc DefaultRuntime:{Type: Path: Engine: PodAnnotations:[] ContainerAnnotations:[] Root: Options:map[] PrivilegedWithoutHostDevices:false PrivilegedWithoutHostDevicesAllDevicesAllowed:false BaseRuntimeSpec: NetworkPluginConfDir: NetworkPluginMaxConfNum:0 Snapshotter: SandboxMode:} UntrustedWorkloadRuntime:{Type: Path: Engine: PodAnnotations:[] ContainerAnnotations:[] Root: Options:map[] PrivilegedWithoutHostDevices:false PrivilegedWithoutHostDevicesAllDevicesAllowed:false BaseRuntimeSpec: NetworkPluginConfDir: NetworkPluginMaxConfNum:0 Snapshotter: SandboxMode:} Runtimes:map[runc:{Type:io.containerd.runc.v2 Path: Engine: PodAnnotations:[] ContainerAnnotations:[] Root: Options:map[SystemdCgroup:true] PrivilegedWithoutHostDevices:false PrivilegedWithoutHostDevicesAllDevicesAllowed:false BaseRunti
meSpec: NetworkPluginConfDir: NetworkPluginMaxConfNum:0 Snapshotter: SandboxMode:podsandbox}] NoPivot:false DisableSnapshotAnnotations:true DiscardUnpackedLayers:true IgnoreBlockIONotEnabledErrors:false IgnoreRdtNotEnabledErrors:false} CniConfig:{NetworkPluginBinDir:/opt/cni/bin NetworkPluginConfDir:/etc/cni/net.d NetworkPluginMaxConfNum:1 NetworkPluginSetupSerially:false NetworkPluginConfTemplate: IPPreference:} Registry:{ConfigPath:/etc/containerd/certs.d Mirrors:map[] Configs:map[] Auths:map[] Headers:map[]} ImageDecryption:{KeyModel:node} DisableTCPService:true StreamServerAddress: StreamServerPort:10010 StreamIdleTimeout:4h0m0s EnableSelinux:false SelinuxCategoryRange:1024 SandboxImage:registry.k8s.io/pause:3.9 StatsCollectPeriod:10 SystemdCgroup:false EnableTLSStreaming:false X509KeyPairStreaming:{TLSCertFile: TLSKeyFile:} MaxContainerLogLineSize:16384 DisableCgroup:false DisableApparmor:false RestrictOOMScoreAdj:false MaxConcurrentDownloads:3 DisableProcMount:false UnsetSeccompProfile: TolerateMissingH
ugetlbController:true DisableHugetlbController:true DeviceOwnershipFromSecurityContext:false IgnoreImageDefinedVolumes:false NetNSMountsUnderStateDir:false EnableUnprivilegedPorts:false EnableUnprivilegedICMP:false EnableCDI:false CDISpecDirs:[/etc/cdi /var/run/cdi] ImagePullProgressTimeout:5m0s DrainExecSyncIOTimeout:0s ImagePullWithSyncFs:false IgnoreDeprecationWarnings:[]} ContainerdRootDir:/var/lib/containerd ContainerdEndpoint:/run/containerd/containerd.sock RootDir:/var/lib/containerd/io.containerd.grpc.v1.cri StateDir:/run/containerd/io.containerd.grpc.v1.cri}"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.664884973Z" level=info msg="Connect containerd service"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.664996691Z" level=info msg="using legacy CRI server"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.665075782Z" level=info msg="using experimental NRI integration - disable nri plugin to prevent this"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.665279072Z" level=info msg="Get image filesystem path \"/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs\""
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666369327Z" level=info msg="Start subscribing containerd event"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666447979Z" level=info msg="Start recovering state"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666521472Z" level=info msg=serving... address=/run/containerd/containerd.sock.ttrpc
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666620200Z" level=info msg=serving... address=/run/containerd/containerd.sock
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666551140Z" level=info msg="Start event monitor"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666675786Z" level=info msg="Start snapshots syncer"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666688102Z" level=info msg="Start cni network conf syncer for default"
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.666700719Z" level=info msg="Start streaming server"
	Oct 02 12:44:01 NoKubernetes-115222 systemd[1]: Started containerd container runtime.
	Oct 02 12:44:01 NoKubernetes-115222 containerd[758]: time="2025-10-02T12:44:01.667191243Z" level=info msg="containerd successfully booted in 0.048143s"
	
	
	==> describe nodes <==
	command /bin/bash -c "sudo /var/lib/minikube/binaries/v0.0.0/kubectl describe nodes --kubeconfig=/var/lib/minikube/kubeconfig" failed with error: /bin/bash -c "sudo /var/lib/minikube/binaries/v0.0.0/kubectl describe nodes --kubeconfig=/var/lib/minikube/kubeconfig": Process exited with status 1
	stdout:
	
	stderr:
	sudo: a terminal is required to read the password; either use the -S option to read from standard input or configure an askpass helper
	sudo: a password is required
	
	
	==> dmesg <==
	[  +0.000008] ll header: 00000000: ff ff ff ff ff ff fe 40 f6 28 01 b6 08 06
	[  +2.473371] IPv4: martian source 10.244.0.1 from 10.244.0.2, on dev eth0
	[  +0.000006] ll header: 00000000: ff ff ff ff ff ff da 82 9e b7 f3 1d 08 06
	[ +12.115583] IPv4: martian source 10.244.0.1 from 10.244.0.3, on dev eth0
	[  +0.000007] ll header: 00000000: ff ff ff ff ff ff fa c9 db f3 a8 cb 08 06
	[  +0.000453] IPv4: martian source 10.244.0.3 from 10.244.0.2, on dev eth0
	[  +0.000004] ll header: 00000000: ff ff ff ff ff ff fe 40 f6 28 01 b6 08 06
	[  +8.697243] IPv4: martian source 10.244.0.1 from 10.244.0.2, on dev eth0
	[  +0.000006] ll header: 00000000: ff ff ff ff ff ff 12 55 f8 67 01 e2 08 06
	[Oct 2 11:45] IPv4: martian source 10.244.0.1 from 10.244.0.2, on dev eth0
	[  +0.000007] ll header: 00000000: ff ff ff ff ff ff 4a b5 6b 10 b4 b6 08 06
	[  +0.072240] IPv4: martian source 10.244.0.1 from 10.244.0.3, on dev eth0
	[  +0.000009] ll header: 00000000: ff ff ff ff ff ff 1a 6c 71 80 c4 33 08 06
	[  +1.016796] IPv4: martian source 10.244.0.1 from 10.244.0.3, on dev eth0
	[  +0.000009] ll header: 00000000: ff ff ff ff ff ff 1a 81 f6 69 f7 42 08 06
	[  +0.000541] IPv4: martian source 10.244.0.3 from 10.244.0.2, on dev eth0
	[  +0.000005] ll header: 00000000: ff ff ff ff ff ff 12 55 f8 67 01 e2 08 06
	[  +6.794314] IPv4: martian source 10.244.0.1 from 10.244.0.3, on dev eth0
	[  +0.000006] ll header: 00000000: ff ff ff ff ff ff 6e d7 18 68 e2 41 08 06
	[  +0.000359] IPv4: martian source 10.244.0.3 from 10.244.0.2, on dev eth0
	[  +0.000016] ll header: 00000000: ff ff ff ff ff ff da 82 9e b7 f3 1d 08 06
	[ +35.779860] IPv4: martian source 10.244.0.1 from 10.244.0.4, on dev eth0
	[  +0.000011] ll header: 00000000: ff ff ff ff ff ff 42 d2 67 3a d3 72 08 06
	[  +0.000391] IPv4: martian source 10.244.0.4 from 10.244.0.3, on dev eth0
	[  +0.000004] ll header: 00000000: ff ff ff ff ff ff 1a 6c 71 80 c4 33 08 06
	
	
	==> kernel <==
	 12:44:04 up  2:26,  0 users,  load average: 4.99, 2.09, 1.54
	Linux NoKubernetes-115222 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux
	PRETTY_NAME="Ubuntu 22.04.5 LTS"
	
	
	==> kubelet <==
	-- No entries --
	
-- /stdout --
helpers_test.go:262: (dbg) Run:  out/minikube-linux-amd64 status --format={{.APIServer}} -p NoKubernetes-115222 -n NoKubernetes-115222
helpers_test.go:262: (dbg) Non-zero exit: out/minikube-linux-amd64 status --format={{.APIServer}} -p NoKubernetes-115222 -n NoKubernetes-115222: exit status 6 (357.598398ms)
-- stdout --
	Stopped
	WARNING: Your kubectl is pointing to stale minikube-vm.
	To fix the kubectl context, run `minikube update-context`
-- /stdout --
** stderr ** 
	E1002 12:44:04.666643  589723 status.go:458] kubeconfig endpoint: get endpoint: "NoKubernetes-115222" does not appear in /home/jenkins/minikube-integration/21139-381342/kubeconfig
** /stderr **
helpers_test.go:262: status error: exit status 6 (may be ok)
helpers_test.go:264: "NoKubernetes-115222" apiserver is not running, skipping kubectl commands (state="Stopped")
--- FAIL: TestNoKubernetes/serial/VerifyNok8sNoK8sDownloads (2.84s)