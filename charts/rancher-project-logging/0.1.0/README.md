# rancher-project-logging

This chart installs [`rancher-logging`](https://rancher.com/docs/rancher/v2.6/en/logging) ClusterFlows that automatically collect logs from all namespaces of a project.

## Prerequisites

- Kubernetes 1.16+
- Helm 3+

## Install Chart

This chart is not intended for standalone use; it's intended to be deployed via [Project Logging Operator](https://github.com/bashofmann/project-logging-operator).

## Dependencies

This chart is designed to be deployed alongside an existing rancher-logging deployment in a cluster that has already installed the Banzai Logging CRDs. Specifically, the chart is configured and intended to be deployed alongside [`rancher-logging`](https://rancher.com/docs/rancher/v2.6/en/logging).

### Configuration

Since this chart installs a project-scoped version of [`rancher-monitoring`](https://rancher.com/docs/rancher/v2.6/en/monitoring-alerting/), a Helm chart based off of [`kube-prometheus-stack`](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack), most of the options that apply to either of those charts will apply to this chart (e.g. support for configuring persistent volumes, ingresses, etc.) and can be passed in as part of the `spec.values` of the ProjectHelmChart that deploys this chart; however, certain advanced functionality (such as Thanos support) and options that pose security risks in Project environments (e.g. ability to `ignoreNamespaceSelectors` or modify the existing namepaceSelectors of the Cluster Prometheus, ability to mount additional scrape configs, etc.) have been removed from the `values.yaml` of the chart. For more information on how to configure values and what they mean, please see the comments and options provided on the `values.yaml` packaged with this chart.

## Further Information

For more in-depth documentation of configuration options meanings, please see

- [`rancher-monitoring`](https://rancher.com/docs/rancher/v2.6/en/monitoring-alerting/)
- [Prometheus Operator](https://github.com/prometheus-operator/prometheus-operator)
- [Prometheus](https://prometheus.io/docs/introduction/overview/)
- [Grafana](https://github.com/grafana/helm-charts/tree/main/charts/grafana#grafana-helm-chart)