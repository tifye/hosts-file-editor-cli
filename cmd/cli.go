package cmd

import (
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

type Cli struct {
	Editor *pkg.Editor
}

func NewCli() *Cli {
	return &Cli{
		Editor: &pkg.Editor{},
	}
}
