package cli

import "github.com/tifye/hosts-file-editor-cli/pkg"

type Cli interface {
	SetHostsFile(*pkg.HostsFile)
	HostsFile() *pkg.HostsFile
	AccessibleMode() bool
}

type HostsCli struct {
	hostsFile  *pkg.HostsFile
	Accessible bool
}

func (cli *HostsCli) HostsFile() *pkg.HostsFile {
	return cli.hostsFile
}

func (cli *HostsCli) AccessibleMode() bool {
	return cli.Accessible
}

func (cli *HostsCli) SetHostsFile(hf *pkg.HostsFile) {
	cli.hostsFile = hf
}
