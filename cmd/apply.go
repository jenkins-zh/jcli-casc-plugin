package cmd

import (
	appCmd "github.com/jenkins-zh/jenkins-cli/app/cmd"
	"github.com/jenkins-zh/jenkins-cli/client"

	"github.com/spf13/cobra"
)

func NewApplyCmd() (cmd *cobra.Command) {
	opt := &CASCApplyOption{}
	cmd = &cobra.Command{
		Use:   "apply",
		Short: "Apply config through configuration-as-code",
		Long:  "Apply config through configuration-as-code",
		RunE:  opt.RunE,
	}
	return
}

func (c *CASCApplyOption) RunE(cmd *cobra.Command, _ []string) error {
	jClient := &client.CASCManager{
		JenkinsCore: client.JenkinsCore{
			RoundTripper: c.RoundTripper,
		},
	}
	appCmd.GetCurrentJenkinsAndClient(&(jClient.JenkinsCore))
	return jClient.Apply()
}
