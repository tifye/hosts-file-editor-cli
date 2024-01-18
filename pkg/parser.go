package pkg

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

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

		var lineComment string
		if cmtStart != -1 {
			lineComment = line[cmtStart+1:]
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
			Comment:  lineComment,
		})
	}

	return &HostsFile{
		Header:  headerLines,
		Entries: entries,
	}, nil
}
