package core

import "fmt"

type Editor struct {
}

func (e *Editor) AddEntry(entry HostEntry) error {
	fmt.Println("Adding entry")
	return nil
}
