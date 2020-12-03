# Attribution

Attribution, a project helps advertisers achieve self-attribution, only cost some simply configuration to quickly use.

This chart bootstraps a [Attribution](https://github.com/TencentAd/attribution) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+
- Helm 3+

## Install Chart

```console
# Helm
$ helm package charts/attribution/
$ helm install ./attribution-[VERSION].tgz --genertate-name
```
_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

## Dependencies
You can enable to installs additional, dependent charts:

- [bitnami/redis](https://artifacthub.io/packages/helm/bitnami/redis)

_See [helm dependency](https://helm.sh/docs/helm/helm_dependency/) for command documentation._

## Uninstall Chart

```console
# Helm
$ helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

