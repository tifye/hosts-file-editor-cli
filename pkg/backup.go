package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CreateBackupFile(hf *HostsFile, cmd string) error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return fmt.Errorf("Failed to determine cache directory %w", err)
	}

	backupDir := fmt.Sprintf("%s\\hosts-file-editor-cli", cacheDir)
	err = os.MkdirAll(backupDir, 0777)
	if err != nil {
		return fmt.Errorf("Failed to create backup directory %w", err)
	}

	d := time.Now()
	filename := fmt.Sprintf("backup-%d-before-%s", d.Unix(), cmd)
	path := filepath.Join(backupDir, filename)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Failed to create backup file %w", err)
	}

	hf.SaveTo(file)

	err = file.Close()
	if err != nil {
		log.Printf("Failed to close backup file %s", err)
	}

	return nil
}
