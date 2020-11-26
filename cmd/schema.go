package cmd

import (
	appCmd "github.com/jenkins-zh/jenkins-cli/app/cmd"
	"github.com/jenkins-zh/jenkins-cli/client"

	"github.com/spf13/cobra"
)

func NewSchemaCmd() (cascSchemaCmd *cobra.Command) {
	opt := &CASCSchemaOption{}
	cascSchemaCmd = &cobra.Command{
		Use:   "schema",
		Short: "Get the schema of configuration-as-code",
		Long:  "Get the schema of configuration-as-code",
		RunE:  opt.RunE,
	}
	return
}

func (c *CASCSchemaOption) RunE(cmd *cobra.Command, _ []string) (err error) {
	jClient := &client.CASCManager{
		JenkinsCore: client.JenkinsCore{
			RoundTripper: c.RoundTripper,
		},
	}
	appCmd.GetCurrentJenkinsAndClient(&(jClient.JenkinsCore))

	var config string
	if config, err = jClient.Schema(); err == nil {
		cmd.Print(config)
	}
	return
}
