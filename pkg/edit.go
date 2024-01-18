package pkg

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Editor struct {
	Writer io.StringWriter
	Reader io.Reader
}

func (e *Editor) AddEntry(entry HostEntry) error {
	entryStr := fmt.Sprintf("%s %s #%s", entry.IP, entry.Hostname, entry.Comment)
	_, err := e.Writer.WriteString(entryStr)
	if err != nil {
		return err
	}

	return nil
}

func (e *Editor) FilterOutEntries(hostname, ip string) ([]HostEntry, error) {
	if hostname == "" && ip == "" {
		var entries []HostEntry
		return entries, nil
	}

	hostsFile, err := ParseHostsFile(e.Reader)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse hosts file got: %w", err)
	}

	filteredEntries := make([]HostEntry, 0)
	for _, entry := range hostsFile.Entries {
		if (hostname == "" || entry.Hostname != hostname) &&
			(ip == "" || entry.IP != ip) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries, nil
}

func (e *Editor) ReplaceWith(entries []HostEntry) {
	scanner := bufio.NewScanner(e.Reader)
	var headerComments []string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if strings.HasPrefix(line, "#") {
			headerComments = append(headerComments, line)
		} else {
			break
		}
	}

	for _, comment := range headerComments {
		e.Writer.WriteString(comment)
	}
	for _, entry := range entries {
		str := fmt.Sprintf("%s %s #%s\n", entry.IP, entry.Hostname, entry.Comment)
		e.Writer.WriteString(str)
	}
}
