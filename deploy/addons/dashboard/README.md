# Kubernetes Dashboard Addon - Migration to Helm

**⚠️ DEPRECATED**: The static YAML manifests in this directory are deprecated and no longer used by Minikube.

## What Changed?

As of Minikube v1.37.0, the Kubernetes Dashboard addon has been converted from static YAML-based installation to a Helm-based deployment using the official Kubernetes Dashboard Helm chart.

## Migration Details

- **Old Approach**: Static YAML manifests (Dashboard v2.7.0)
- **New Approach**: Helm chart from `https://kubernetes.github.io/dashboard/`
- **Chart**: `kubernetes-dashboard/kubernetes-dashboard` (v7.x)

## Configuration

The dashboard addon is now configured in `pkg/minikube/assets/addons.go` using the HelmChart struct:

```go
&HelmChart{
    Name:      "kubernetes-dashboard",
    Repo:      "kubernetes-dashboard/kubernetes-dashboard",
    RepoURL:   "https://kubernetes.github.io/dashboard/",
    Namespace: "kubernetes-dashboard",
    Values: []string{
        "metricsScraper.enabled=true",
        "nginx.enabled=false",
        "cert-manager.enabled=false",
    },
}
```

## Breaking Changes

Dashboard v7.x **requires authentication**. The `--enable-skip-login` and `--disable-settings-authorizer` flags are no longer supported.

### Accessing the Dashboard

**Recommended method:**
```bash
minikube dashboard
```

**Manual access:**
```bash
# Port forward
kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-kong-proxy 8443:443

# Create authentication token
kubectl -n kubernetes-dashboard create token kubernetes-dashboard
```

## Benefits

1. **Easier Version Upgrades**: Single line change to update dashboard version
2. **Upstream Alignment**: Uses official Kubernetes Dashboard Helm chart
3. **Better Security**: Authentication is now mandatory
4. **Maintainability**: Leverages Helm's release management

## For Minikube Developers

To update the dashboard version, modify the chart reference in `pkg/minikube/assets/addons.go`. The Helm chart automatically pulls the latest compatible version from the repository.

## Legacy Files

The following files in this directory are kept for reference but are no longer used:
- dashboard-ns.yaml
- dashboard-clusterrole.yaml
- dashboard-clusterrolebinding.yaml
- dashboard-configmap.yaml
- dashboard-dp.yaml.tmpl
- dashboard-role.yaml
- dashboard-rolebinding.yaml
- dashboard-sa.yaml
- dashboard-secret.yaml
- dashboard-svc.yaml

These files may be removed in a future Minikube release.
