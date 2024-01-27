package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Backup struct {
	Time     time.Time
	Filepath string
	Comment  string
}

var pathToBackups *string

func GetBackupsDirPath() (string, error) {
	if pathToBackups != nil {
		return *pathToBackups, nil
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("Failed to determine cache directory %w", err)
	}

	return fmt.Sprintf("%s\\hosts-file-editor-cli", cacheDir), nil
}

func GetListOfBackups() ([]Backup, error) {
	backupsPath, err := GetBackupsDirPath()
	if err != nil {
		return nil, err
	}

	dirEntries, err := os.ReadDir(backupsPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read backups directory: %w", err)
	}

	var backups []Backup
	for _, dirEntry := range dirEntries {
		backup, err := parseBackupName(dirEntry.Name())
		if err != nil {
			return nil, fmt.Errorf("Failed to parse backup file %s got: %s", dirEntry.Name(), err)
		}
		if backup == nil {
			continue
		}
		backups = append(backups, *backup)
	}

	sort.Slice(backups, func(ia, ib int) bool {
		return backups[ia].Time.Before(backups[ib].Time)
	})
	return backups, nil
}

func parseBackupName(name string) (*Backup, error) {
	parts := strings.Split(name, "-")
	if len(parts) < 0 || len(parts) < 2 || parts[0] != "backup" {
		return nil, nil
	}

	unixTimeStr := parts[1]
	unixTime, err := strconv.Atoi(unixTimeStr)
	if err != nil {
		return nil, errors.New("Could not parse backup time string")
	}
	time := time.Unix(int64(unixTime), 0)

	comment := strings.Join(parts[2:], " ")
	return &Backup{
		Time:     time,
		Filepath: name,
		Comment:  comment,
	}, nil
}

func CreateBackupFile(hf *HostsFile, cmd string) error {
	backupDir, err := GetBackupsDirPath()
	if err != nil {
		return err
	}

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
