package cmd

import "github.com/tifye/hosts-file-editor-cli/core"

type Cli struct {
	Editor *core.Editor
}

func NewCli() *Cli {
	return &Cli{
		Editor: &core.Editor{},
	}
}
