package cmd

import (
	cmdApp "github.com/jenkins-zh/jenkins-cli/app/cmd"
	"github.com/jenkins-zh/jenkins-cli/app/cmd/common"
	"github.com/spf13/cobra"
)

// Check do the health check of casc cmd
func (o *CASCOptions) Check() (err error) {
	opt := cmdApp.PluginOptions{
		Option: common.Option{RoundTripper: o.RoundTripper},
	}
	_, err = opt.FindPlugin("configuration-as-code")
	return
}

func NewCasCCMD() (cascCmd *cobra.Command) {
	opt := CASCOptions{}
	cascCmd = &cobra.Command{
		Use:   "jcli casc",
		Short: "Configuration as Code",
		Long:  "Configuration as Code",
	}
	cmdApp.GetHealthCheckRegister().Register("jcli.casc.*", &opt)

	cascCmd.AddCommand(
		NewSchemaCmd(),
		NewReloadCmd(),
		NewExportCmd(),
		NewApplyCmd(),
		NewOpenCmd())
	return
}
