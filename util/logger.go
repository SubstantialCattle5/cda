package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CreateLogDirectories ensures that all required log directories are created.
func CreateLogDirectories() error {
	directoryFiles := []string{"checksums"}
	subDirectoryFiles := []string{"export", "verify", "export-diff"}

	for _, directoryFile := range directoryFiles {
		for _, subDirectoryFile := range subDirectoryFiles {
			dir := filepath.Join(directoryFile, subDirectoryFile)
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dir, err)
			}
		}
	}
	return nil
}

// CreateLogFile creates a log file based on Type and logFileName.
func CreateLogFile(Type string, logFileName string) (*os.File, error) {
	var directory string

	switch Type {
	case "export", "verify", "export-diff":
		directory = filepath.Join("checksums", Type)
	default:
		return nil, fmt.Errorf("invalid log file Type: %s", Type)
	}

	// Ensure the directory for the log file exists
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", directory, err)
	}

	// Construct the full path for the log file
	logFilePath := filepath.Join(directory, logFileName)

	// Create the file or open it if it exists
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %s: %w", logFilePath, err)
	}

	return file, nil
}

func logContent(hash string, timestamp string, filePath string) error {
	logContent := fmt.Sprintf("%s %s\n", timestamp, hash)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(logContent)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}
	fmt.Printf("Successfully wrote to file %s\n", filePath)
	return nil
}

// SaveFile saves a file with given imageName, poolName, and hash based on Type.
func SaveFile(Type string, imageName string, poolName string, hash string) error {
	// Ensure the log directories are created
	err := CreateLogDirectories()
	if err != nil {
		return fmt.Errorf("failed to create log directories: %w", err)
	}

	// Generate the filename
	filename := strings.Join([]string{poolName, imageName}, "_") + ".log"
	fmt.Printf("Saving it into file %s...........\n", filename)

	// Create the log file
	file, err := CreateLogFile(Type, filename)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer file.Close()

	// Write to the log file
	timestamp := time.Now() // replace with actual timestamp
	err = logContent(hash, timestamp.Format("20060102150405"), file.Name())
	if err != nil {
		return fmt.Errorf("failed to write log content: %w", err)
	}

	return nil
}

//
//func LoadFile() {
//
//}
//
//func deleteFile() {
//
//}
//
//func ChecksumLoggerSave(data []byte) {
//
//}
//func ChecksumLoggerFetch(data []byte) {
//
//}
