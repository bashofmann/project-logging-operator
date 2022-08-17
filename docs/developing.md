# Developing Project Logging Operator

The Project Logging Operator repository is primarily comprised of just two things:
- A simple `main.go` that implements [Helm Project Operator](https://github.com/rancher/helm-project-operator) for the [`rancher-project-logging` chart](../charts/rancher-project-logging)
- A `packages/` directory that corresponds to a [`rancher/charts-build-scripts`](https://github.com/rancher/charts-build-scripts) repository

In **most** circumstances, you will only ever have to make changes to the `packages/` directory; if you need to make changes to the underlying code of the operator that is deployed, it is likely that you intend to make this change in [rancher/helm-project-operator](https://github.com/rancher/helm-project-operator) instead.

## Repository Structure

```bash
## This directory is a [`rancher/charts-build-scripts`](https://github.com/rancher/charts-build-scripts) packages directory. See below for more details.
packages/

## This directory contains **auto-generated** Helm chart archives that can be used to deploy Project Logging Operator in a Kubernetes cluster in 
## the cattle-logging-system namespace, which deploys rancher-project-logging (located under charts/rancher-project-logging) 
## on seeing a ProjectHelmChart with spec.helmApiVersion: logging.cattle.io/v1alpha1.
##
## IMPORTANT: You should never modify the contents of this directory directly; you should always modify `packages` since that will 
## overwrite the changes that are observed in this directory on running a `make charts`.
##
assets/

## This file is an **auto-generated** Helm index.yaml identifying this repository as a valid Helm repository that contains Helm charts.
##
## IMPORTANT: You should never modify the contents of this file directly; you should always modify `packages` since that will 
## overwrite the changes that are observed in this directory on running a `make charts` or `make index`.
##
index.yaml

## This directory contains **auto-generated** Helm charts that can be used to deploy Project Logging Operator in a Kubernetes cluster in 
## the cattle-logging-system namespace, which deploys rancher-project-logging (located under charts/rancher-project-logging) 
## on seeing a ProjectHelmChart with spec.helmApiVersion: logging.cattle.io/v1alpha1.
##
## IMPORTANT: You should never modify the contents of this directory directly; you should always modify `packages` since that will 
## overwrite the changes that are observed in this directory on running a `make charts`.
charts/

  ## The main chart that deploys Project Logging Operator in the cluster.
  project-logging-operator/*
  
  ## A char that deploys a Project Logging ClusterFlow onto the cluster on seeing a valid ProjectHelmChart (which means that it is contained within a Project Registration Namespace
  ## with spec.helmApiVersion set to logging.cattle.io/v1alpha1)
  ##
  ## This chart is not expected to ever be deployed standalone; it is embedded into the Project Logging Operator binary itself.
  rancher-project-logging/*

## This directory will contain additional docs to assist users in getting started with using Helm Project Operator.
docs/

## This directory contains an example ProjectHelmChart that can be deployed to create an example Project Logging ClusterFlow
## Note: the namespace needs to be modified to be a valid Project Registration Namespace, depending on how you deployed the operator.
examples/

## This directory contains the image that is used to build rancher/helm-project-operator, which is hosted on hub.docker.com.
package/
  Dockerfile

## The main entrypoint into Project Logging Operator that implements Helm Project Operator.
main.go

## The Dockerfile used to run CI and other scripts executed by make in a Docker container (powered by https://github.com/rancher/dapper)
Dockerfile.dapper
```

## Making changes to the Helm Charts (`packages/`)

In most situations, the changes made to this repository will primarily be fixes to the Helm charts that either deploy the operator (`project-logging-operator`) or those that are deployed on behalf of the operator (`rancher-project-logging`).

If you need to bump the version of Helm Project Operator embedded into the charts or binaries, generally you will need to bump the version of the Helm Project Operator in the `go.mod` and update the commit hash in `packages/project-logging-operator/generated-changes/dependencies/helmProjectOperator/dependency.yaml`; once done, run `go mod tidy` and make one commit with your changes entitled `Bump Helm Project Operator` followed by one commit with the output of running `unset PACKAGE; make charts` with the commit message `make charts`.

If you need to make changes to the Project Logging Operator chart itself, make the changes directly in the `packages/project-logging-operator/charts`; once done, make one or more commits that only contain your changes to the `packages/project-logging-operator/charts` directory with proper commit messages describing what you changed and make one commit at the end with the output of running `unset PACKAGE; make charts` with the commit message `make charts`.

If you need to make changes to the rancher-project-logging chart, follow the same steps above but start by running `PACKAGE=rancher-project-logging make prepare` to pull in the latest version of your `rancher-project-logging` chart. Before you commit any changes, always make sure you run `PACKAGE=rancher-project-logging make clean` to avoid committing `packages/rancher-project-logging/charts/charts` (but be careful since `make clean` will wipe out any changes you made to that directory! It does the equivalent of `rm -rf packages/rancher-project-logging/charts/charts`).

> Note: In general, it is recommended to use the experimental caching feature for rancher/charts-build-scripts to avoid multiple network calls to pull in the source repositories by storing them in a local cache under `.charts-build-scripts/.cache/*`. You can turn this on by default by setting `export USE_CACHE=1`.

For more information on how to make changes on repositories powered by `rancher/charts-build-scripts`, please read the [docs](https://github.com/rancher/charts-build-scripts/tree/master/templates/template/docs).

## Once you have made a change

If you modified `packages/`, make sure you run `unset PACKAGE; make charts` to generate the latest `charts/`, `assets/` and `index.yaml`.

Also, make sure you run `go mod tidy` if you make any changes to the code.

## Creating a Docker image based off of your changes

To test your changes and create a Docker image to a specific Docker repository with a given tag, you should run `REPO=<my-docker-repo> TAG=<my-docker-tag> make` (e.g. `REPO=arvindiyengar TAG=dev make`), which will run the `./scripts/ci` script that builds, tests, validates, and packages your changes into a local Docker image (if you run `docker images`, it should show up as an image in the format `${REPO}/project-logging-operator:${TAG}`).

If you don't want to run all the steps in CI every time you make a change, you could also run the following one-liner to build and package the image:

```bash
REPO=<my-repo>
TAG=<my-tag>

./scripts/build-chart && GOOS=linux CGO_ENABLED=0 go build -ldflags "-extldflags -static -s" -o bin/project-logging-operator && REPO=${REPO} TAG=${TAG} make package
```

Once the image is successfully packaged, simply run `docker push ${REPO}/project-logging-operator:${TAG}` to push your image to your Docker repository.

## Testing a custom Docker image build

1. Ensure that your `KUBECONFIG` environment variable is pointing to your cluster (e.g. `export KUBECONFIG=<path-to-kubeconfig>; kubectl get nodes` should show the nodes of your cluster) and pull in this repository locally
2. Go to the root of your local copy of this repository and deploy the Project Logging Operator chart as a Helm 3 chart onto your cluster after overriding the image and tag values with your Docker repository and tag: run `helm upgrade --install --set image.repository="${REPO}/project-logging-operator" --set image.tag="${TAG}" --set image.pullPolicy=Always project-logging-operator -n cattle-logging-system charts/project-logging-operator`
> Note: Why do we set the Image Pull Policy to `Always`? If you update the Docker image on your fork, setting the Image Pull Policy to `Always` ensures that running `kubectl rollout restart -n cattle-logging-system deployment/project-logging-operator` is all you need to do to update your running deployment to the new image, since this would ensure redeploying a deployment triggers a image pull that uses your most up-to-date Docker image. Also, since the underlying Helm chart deployed by the operator (e.g. `example-chart`) is directly embedded into the Helm Project Operator image, you also do not need to update the Deployment object itself to see all the HelmCharts in your cluster automatically be updated to the latest embedded version of the chart.
3. Profit!