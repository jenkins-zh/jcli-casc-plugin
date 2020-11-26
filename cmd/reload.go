package cmd

import (
	appCmd "github.com/jenkins-zh/jenkins-cli/app/cmd"
	"github.com/jenkins-zh/jenkins-cli/client"

	"github.com/spf13/cobra"
)

func (c *CASCReloadOption) RunE(cmd *cobra.Command, _ []string) error {
	jClient := &client.CASCManager{
		JenkinsCore: client.JenkinsCore{
			RoundTripper: c.RoundTripper,
		},
	}
	appCmd.GetCurrentJenkinsAndClient(&(jClient.JenkinsCore))
	return jClient.Reload()
}

func NewReloadCmd() (cmd *cobra.Command) {
	opt := &CASCReloadOption{}
	cmd = &cobra.Command{
		Use:   "reload",
		Short: "Reload config through configuration-as-code",
		Long:  "Reload config through configuration-as-code",
		RunE:  opt.RunE,
	}
	return
}
