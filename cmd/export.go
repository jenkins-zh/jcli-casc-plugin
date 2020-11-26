package cmd

import (
	appCmd "github.com/jenkins-zh/jenkins-cli/app/cmd"
	"github.com/jenkins-zh/jenkins-cli/client"
	"github.com/spf13/cobra"
)

func NewExportCmd() (cmd *cobra.Command) {
	opt := &CASCExportOption{}
	cmd = &cobra.Command{
		Use:   "export",
		Short: "Export the config from configuration-as-code",
		Long:  "Export the config from configuration-as-code",
		RunE:  opt.RunE,
	}
	return
}

func (c *CASCExportOption) RunE(cmd *cobra.Command, _ []string) (err error) {
	jClient := &client.CASCManager{
		JenkinsCore: client.JenkinsCore{
			RoundTripper: c.RoundTripper,
		},
	}
	appCmd.GetCurrentJenkinsAndClient(&(jClient.JenkinsCore))

	var config string
	if config, err = jClient.Export(); err == nil {
		cmd.Print(config)
	}
	return
}
