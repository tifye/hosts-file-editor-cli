package pkg

import (
	"fmt"
	"io"
)

type HostEntry struct {
	IP       string
	Hostname string
	Comment  string
}

type HostsFile struct {
	Header  []string
	Entries []HostEntry
}

func (e HostEntry) String() string {
	return fmt.Sprintf("%s %s #%s", e.IP, e.Hostname, e.Comment)
}

func (hf *HostsFile) SaveTo(target io.Writer) {
	for _, line := range hf.Header {
		target.Write([]byte(line + "\r\n"))
	}
	for _, entry := range hf.Entries {
		target.Write([]byte(entry.String() + "\r\n"))
	}
}

func (hf *HostsFile) AddEntry(entry HostEntry) {
	hf.Entries = append(hf.Entries, entry)
}
