package cmd

import (
	appCmd "github.com/jenkins-zh/jenkins-cli/app/cmd"
	cmdApp "github.com/jenkins-zh/jenkins-cli/app/cmd"
	"github.com/jenkins-zh/jenkins-cli/app/cmd/common"
	"github.com/spf13/cobra"
	"os"
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
	rootCmd := appCmd.GetRootCommand()
	cascCmd = &cobra.Command{
		Use:               "casc",
		Short:             "Configuration as Code",
		Long:              "Configuration as Code",
		ValidArgsFunction: common.NoFileCompletion,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SetOut(os.Stdout)
			return rootCmd.PersistentPreRunE(cmd, args)
		},
		PersistentPostRun: rootCmd.PersistentPostRun,
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
