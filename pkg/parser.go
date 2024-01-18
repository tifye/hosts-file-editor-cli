package pkg

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type HostEntry struct {
	IP       string
	Hostname string
	Comment  string
}

type HostsFile struct {
	Header  string
	Entries []HostEntry
}

func (e HostEntry) String() string {
	return fmt.Sprintf("%s %s %s", e.IP, e.Hostname, e.Comment)
}

func ParseHostsFile(reader io.Reader) (*HostsFile, error) {
	var entries []HostEntry
	scanner := bufio.NewScanner(reader)
	var lineNum int
	readHeader := false
	var headerLines []string
	for scanner.Scan() {
		lineNum += 1

		line := scanner.Text()
		cmtStart := strings.Index(line, "#")
		if readHeader == false && cmtStart == 0 {
			headerLines = append(headerLines, line)
			continue
		}

		readHeader = true

		if cmtStart != -1 {
			line = line[:cmtStart]
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid line(%dn): %s", lineNum, line)
		}

		entries = append(entries, HostEntry{
			IP:       parts[0],
			Hostname: parts[1],
		})
	}

	return &HostsFile{
		Header:  strings.Join(headerLines, "\n"),
		Entries: entries,
	}, nil
}
