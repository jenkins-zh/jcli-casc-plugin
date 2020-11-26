package main

import (
	"github.com/jenkins-zh/jcli-casc-plugin/cmd"
	"os"
)

func main() {
	command := cmd.NewCasCCMD()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
