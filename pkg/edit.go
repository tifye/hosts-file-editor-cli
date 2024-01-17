package core

import (
	"fmt"
	"io"
)

type Editor struct {
	Writer io.StringWriter
}

func (e *Editor) AddEntry(entry HostEntry) error {
	entryStr := fmt.Sprintf("%s %s #%s", entry.IP, entry.Hostname, entry.Comment)
	_, err := e.Writer.WriteString(entryStr)
	if err != nil {
		return err
	}

	return nil
}
