package core

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type HostEntry struct {
	IP       string
	Hostname string
}

func ParseHostsFile(reader io.Reader) ([]HostEntry, error) {
	var entries []HostEntry
	scanner := bufio.NewScanner(reader)
	var lineNum int
	for scanner.Scan() {
		lineNum += 1

		line := scanner.Text()
		cmtStart := strings.Index(line, "#")
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

	return entries, nil
}
