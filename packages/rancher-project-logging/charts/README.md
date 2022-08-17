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

For more information on how to configure values and what they mean, please see the comments and options provided on the `values.yaml` packaged with this chart.

## Further Information

For more in-depth documentation of configuration options meanings, please see

- [`rancher-logging`](https://rancher.com/docs/rancher/v2.6/en/logging/)
- [Banzai Logging Operator](https://banzaicloud.com/docs/one-eye/logging-operator/)
