package cmd

import (
	"github.com/jenkins-zh/jenkins-cli/app/cmd/common"
	"github.com/jenkins-zh/jenkins-cli/util"
	"net/http"
)

// CASCReloadOption as the options of reload configuration as code
type CASCReloadOption struct {
	RoundTripper http.RoundTripper
}

// CASCSchemaOption as the options of reload configuration as code
type CASCSchemaOption struct {
	RoundTripper http.RoundTripper
}

// CASCOpenOption is the option of casc open cmd
type CASCOpenOption struct {
	ExecContext util.ExecContext

	Browser string
}

// CASCExportOption as the options of reload configuration as code
type CASCExportOption struct {
	RoundTripper http.RoundTripper
}

// CASCApplyOption as the options of apply configuration as code
type CASCApplyOption struct {
	RoundTripper http.RoundTripper
}

// CASCOptions is the option of casc
type CASCOptions struct {
	common.Option
}
