package pkg

import (
	"fmt"
	"os"
)

func SaveToFile(hf *HostsFile, filepath string) error {
	file, err := os.OpenFile(filepath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Failed to open hosts file for writing %w.\nTry running as administrator", err)
	}

	file.Seek(0, 0)
	hf.SaveTo(file)

	if err = file.Close(); err != nil {
		return fmt.Errorf("Failed to close hosts file %w", err)
	}

	return nil
}
