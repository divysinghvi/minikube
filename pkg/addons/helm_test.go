package addons

import (
	"context"
	"os/exec"
	"strings"
	"testing"

	"k8s.io/minikube/pkg/minikube/assets"
	"k8s.io/minikube/pkg/minikube/command"
)

// mockRunner is a mock implementation of command.Runner for testing
type mockRunner struct{}

func (m *mockRunner) RunCmd(cmd *exec.Cmd) (*command.RunResult, error) {
	// Mock successful execution
	return &command.RunResult{}, nil
}

func (m *mockRunner) StartCmd(cmd *exec.Cmd) (*command.StartedCmd, error) {
	return &command.StartedCmd{}, nil
}

func (m *mockRunner) WaitCmd(startedCmd *command.StartedCmd) (*command.RunResult, error) {
	return &command.RunResult{}, nil
}

func (m *mockRunner) Copy(f assets.CopyableFile) error {
	return nil
}

func (m *mockRunner) CopyFrom(f assets.CopyableFile) error {
	return nil
}

func (m *mockRunner) Remove(f assets.CopyableFile) error {
	return nil
}

func (m *mockRunner) ReadableFile(sourcePath string) (assets.ReadableFile, error) {
	return nil, nil
}

func TestHelmCommand(t *testing.T) {
	tests := []struct {
		description string
		chart       *assets.HelmChart
		enable      bool
		expected    string
		mode        string
	}{
		{
			description: "enable an addon",
			chart: &assets.HelmChart{
				Name:       "addon-name",
				Repo:       "addon-repo/addon-chart",
				Namespace:  "addon-namespace",
				Values:     []string{"key=value"},
				ValueFiles: []string{"/etc/kubernetes/addons/values.yaml"},
			},
			enable:   true,
			expected: "sudo KUBECONFIG=/var/lib/minikube/kubeconfig helm upgrade --install addon-name addon-repo/addon-chart --create-namespace --namespace addon-namespace --set key=value --values /etc/kubernetes/addons/values.yaml",
		},
		{
			description: "enable an addon without namespace",
			chart: &assets.HelmChart{
				Name:       "addon-name",
				Repo:       "addon-repo/addon-chart",
				Values:     []string{"key=value"},
				ValueFiles: []string{"/etc/kubernetes/addons/values.yaml"},
			},
			enable:   true,
			expected: "sudo KUBECONFIG=/var/lib/minikube/kubeconfig helm upgrade --install addon-name addon-repo/addon-chart --create-namespace --set key=value --values /etc/kubernetes/addons/values.yaml",
		},
		{
			description: "disable an addon",
			chart: &assets.HelmChart{
				Name:      "addon-name",
				Namespace: "addon-namespace",
			},
			enable:   false,
			expected: "sudo KUBECONFIG=/var/lib/minikube/kubeconfig helm uninstall addon-name --namespace addon-namespace",
			mode:     "cpu",
		},
	}

	runner := &mockRunner{}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			command, err := helmUninstallOrInstall(context.Background(), test.chart, runner, test.enable)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			actual := strings.Join(command.Args, " ")
			if actual != test.expected {
				t.Errorf("helm command mismatch:\nexpected: %s\nactual:   %s", test.expected, actual)
			}
		})
	}
}

func TestRepoNameExtraction(t *testing.T) {
	tests := []struct {
		repo     string
		expected string
	}{
		{"kubernetes-dashboard/kubernetes-dashboard", "kubernetes-dashboard"},
		{"my-repo/my-chart", "my-repo"},
		{"single-name", "single-name"}, // No slash, should return whole string
	}

	for _, test := range tests {
		t.Run(test.repo, func(t *testing.T) {
			result := test.repo
			if idx := strings.Index(test.repo, "/"); idx > 0 {
				result = test.repo[:idx]
			}
			if result != test.expected {
				t.Errorf("repo name extraction mismatch for %s:\nexpected: %s\nactual:   %s", test.repo, test.expected, result)
			}
		})
	}
}
