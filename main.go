package main

import (
	_ "embed"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/bashofmann/project-logging-operator/pkg/version"
	"github.com/rancher/helm-project-operator/pkg/controllers/common"
	"github.com/rancher/helm-project-operator/pkg/operator"
	command "github.com/rancher/wrangler-cli"
	_ "github.com/rancher/wrangler/pkg/generated/controllers/apiextensions.k8s.io"
	_ "github.com/rancher/wrangler/pkg/generated/controllers/networking.k8s.io"
	"github.com/rancher/wrangler/pkg/kubeconfig"
	"github.com/spf13/cobra"
)

const (
	// HelmAPIVersion is the spec.helmApiVersion corresponding to the embedded logging chart (rancher-project-logging)
	HelmAPIVersion = "logging.cattle.io/v1alpha1"

	// ReleaseName is the release name this operator uses to prefix releases and project release namespaces created on
	// deploying the embedded logging chart (rancher-project-logging)
	ReleaseName = "logging"
)

var (
	// SystemNamespaces is the system namespaces scoped for the embedded logging chart (rancher-project-logging)
	SystemNamespaces = []string{"kube-system", "cattle-logging-system"}

	//go:embed bin/rancher-project-logging/rancher-project-logging.tgz.base64
	base64TgzChart string

	debugConfig command.DebugConfig
)

type RancherLoggingOperator struct {
	// Note: all Project Operator are expected to provide these RuntimeOptions
	common.RuntimeOptions

	Kubeconfig string `usage:"Kubeconfig file"`
}

func (f *RancherLoggingOperator) Run(cmd *cobra.Command, args []string) error {
	go func() {
		// required to set up healthz and pprof handlers
		log.Println(http.ListenAndServe("localhost:80", nil))
	}()
	debugConfig.MustSetupDebug()

	cfg := kubeconfig.GetNonInteractiveClientConfig(f.Kubeconfig)

	ctx := cmd.Context()

	if err := operator.Init(ctx, f.Namespace, cfg, common.Options{
		OperatorOptions: common.OperatorOptions{
			HelmAPIVersion:   HelmAPIVersion,
			ReleaseName:      ReleaseName,
			SystemNamespaces: SystemNamespaces,
			ChartContent:     base64TgzChart,
			Singleton:        true, // indicates only one HelmChart can be registered per project defined
		},
		RuntimeOptions: f.RuntimeOptions,
	}); err != nil {
		return err
	}

	<-cmd.Context().Done()
	return nil
}

func main() {
	cmd := command.Command(&RancherLoggingOperator{}, cobra.Command{
		Version: version.FriendlyVersion(),
	})
	cmd = command.AddDebug(cmd, &debugConfig)
	command.Main(cmd)
}
