# Attribution

Attribution, 帮助广告主实现自归因的项目，只需要简单的配置，就能快速使用。
本chart将会在[Kubernetes](http://kubernetes.io) cluster 使用[Helm](https://helm.sh) 部署一个Attribution deployment。

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

