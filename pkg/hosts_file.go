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
		line := fmt.Sprintf("%s\t%s\t#%s", entry.IP, entry.Hostname, entry.Comment)
		target.Write([]byte(line + "\r\n"))
	}
}

func (hf *HostsFile) AddEntry(entry HostEntry) {
	hf.Entries = append(hf.Entries, entry)
}

func FilterOut(entries []HostEntry, hostname, ip string) []HostEntry {
	if hostname == "" && ip == "" {
		return nil
	}

	var filtered []HostEntry
	for _, entry := range entries {
		if (hostname == "" || entry.Hostname != hostname) &&
			(ip == "" || entry.IP != ip) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}
