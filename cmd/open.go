package cmd

import (
	"errors"
	"fmt"
	"os"

	appCmd "github.com/jenkins-zh/jenkins-cli/app/cmd"
	"github.com/jenkins-zh/jenkins-cli/util"
	"github.com/spf13/cobra"
)

func NewOpenCmd() (cmd *cobra.Command) {
	opt := &CASCOpenOption{}
	cmd = &cobra.Command{
		Use:    "open",
		Short:  "Open Configuration as Code page in browser",
		Long:   "Open Configuration as Code page in browser",
		PreRun: opt.PreRun,
		RunE:   opt.RunE,
	}

	cmd.Flags().StringVarP(&opt.Browser, "browser", "b", "",
		"Open Jenkins with a specific browser")
	return
}

func (c *CASCOpenOption) PreRun(_ *cobra.Command, _ []string) {
	if c.Browser == "" {
		c.Browser = os.Getenv("JCLI_BROWSER")
	}
}

func (c *CASCOpenOption) RunE(_ *cobra.Command, _ []string) (err error) {
	jenkins := appCmd.GetCurrentJenkinsFromOptions()
	if jenkins == nil {
		err = errors.New("cannot found Jenkins by --jenkins")
		return
	}

	browser := c.Browser
	err = util.Open(fmt.Sprintf("%s/configuration-as-code", jenkins.URL), browser, c.ExecContext)
	return
}
