project-logging-operator
========

Project Logging Operator is an operator (powered by [`rancher/helm-project-operator`](https://github.com/rancher/helm-project-operator) and [`rancher/charts-build-scripts`](https://github.com/rancher/charts-build-scripts)) that manages deploying one or more Project Logging ClusterFlows that are scoped to project namespaces.

A user can specify that they would like to deploy a Project Logging ClusterFlow by creating a `ProjectHelmChart` CR in a Project Registration Namespace (`cattle-project-<id>`) with `spec.helmApiVersion: logging.cattle.io/v1alpha1`.

> Note: Since this Project Logging ClusterFlow deploys Banzai Logging Operator CRs, an existing Banzai Logging Operator instance must already be deployed in the cluster. It is recommended to use [`rancher-logging`](https://rancher.com/docs/rancher/v2.6/en/logging/) for this. For more information on how the chart works or advanced configurations, please read the [`README.md` on the chart](packages/project-logging-operator/charts/README.md).

For more information on ProjectHelmCharts and how to configure the underlying operator, please read the [`README.md` on the chart](packages/project-logging-operator/charts/README.md) or check out the general docs on Helm Project Operators in [`rancher/helm-project-operator`](https://github.com/rancher/helm-project-operator).

For more information on how to configure the underlying Project Logging ClusterFlow, please read the [`README.md` of the underlying chart](packages/rancher-project-logging/charts/README.md) (`rancher-project-logging`).


## Getting Started

For more information, see the [Getting Started guide](docs/gettingstarted.md).

## Developing

### Which branch do I make changes on?

Project Logging Operator is built and released off the contents of the `main` branch. To make a contribution, open up a PR to the `main` branch.

For more information, see the [Developing guide](docs/developing.md).

## Building

`make`

## Running

`./bin/project-logging-operator`

## License
Copyright (c) 2020 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
