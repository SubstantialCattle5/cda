/*
Copyright © 2024 Nilay Nath Sharan nilaynathsharan16@gmail.com
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify <export/export-diff> <dc_log_file> <dr_log_file> <pool_name> <image_name>",
	Short: "Compares hashes from DC and DR log files",
	Long:  `This command compares the hashes saved in the DC and DR log files to verify data integrity.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 5 {
			log.Fatalf("usage: %s verify <export/export-diff> <dc_log_file> <dr_log_file> <pool_name> <image_name>", os.Args[0])
		}
		exportCMD, dcLogFile, drLogFile, poolName, imageName := args[0], args[1], args[2], args[3], args[4]

		string1, dateTimeDc, err := checkSumCheck(exportCMD, dcLogFile, poolName, imageName)
		if err != nil {
			log.Fatal(err)
		}
		string2, dateTimeDr, err := checkSumCheck(exportCMD, drLogFile, poolName, imageName)
		if err != nil {
			log.Fatal(err)
		}

		validity := compareHashes(string1, string2)
		fmt.Printf("Verification result: %v\n", validity)

		// Implement verification write logic as needed
		err = verifyWrite(dcLogFile, drLogFile, poolName, imageName, string1, dateTimeDc, string2, dateTimeDr, validity)
		if err != nil {
			return
		}
	},
}

func checkSumCheck(exportCMD string, logFile string, poolName string, imageName string) (string, int, error) {
	// Check if the specified path exists
	var validPath = logFile + "/" + exportCMD
	if _, err := os.Stat(validPath); os.IsNotExist(err) {
		return "", 0, fmt.Errorf("%s does not exist", validPath)
	}

	// Construct the file path
	filePath := logFile + "/" + exportCMD + "/" + poolName + "_" + imageName + ".log"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", 0, fmt.Errorf("%s does not exist", filePath)
	}

	// Read the file contents and process the last line
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", 0, err
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) == 0 {
		return "", 0, fmt.Errorf("%s contains no data", filePath)
	}
	lastLine := lines[len(lines)-1]
	parts := strings.Split(lastLine, " ")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("%s has an invalid format", filePath)
	}
	datetime, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", 0, fmt.Errorf("%s has an invalid datetime", filePath)
	}

	return parts[1], datetime, nil
}

func compareHashes(hash1, hash2 string) bool {
	if len(hash1) != len(hash2) {
		return false
	}
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			return false
		}
	}
	return true
}
func verifyWrite(dcLogFile string, drLogFile string, poolName string, imageName string, string1 string, dateTimeDc int, string2 string, dateTimeDr int, answer bool) error {
	// Ensure directories exist or create them
	if err := ensureDirectoryExists(dcLogFile); err != nil {
		return fmt.Errorf("failed to ensure DC directory exists: %s", err)
	}
	if err := ensureDirectoryExists(drLogFile); err != nil {
		return fmt.Errorf("failed to ensure DR directory exists: %s", err)
	}

	// Construct file paths for verification logs
	dcFilePath := fmt.Sprintf("%s/verify/%s_%s.log", dcLogFile, poolName, imageName)
	drFilePath := fmt.Sprintf("%s/verify/%s_%s.log", drLogFile, poolName, imageName)

	// Write verification results to DC and DR log files
	if err := writeVerificationFile(dcFilePath, string1, dateTimeDc, string2, dateTimeDr, time.Now().Unix(), answer); err != nil {
		return fmt.Errorf("failed to write DC verification file: %s", err)
	}
	if err := writeVerificationFile(drFilePath, string1, dateTimeDc, string2, dateTimeDr, time.Now().Unix(), answer); err != nil {
		return fmt.Errorf("failed to write DR verification file: %s", err)
	}

	return nil
}

func ensureDirectoryExists(path string) error {
	// Create directories if they don't exist
	verifyPath := fmt.Sprintf("%s/verify", path)
	if err := os.MkdirAll(verifyPath, 0755); err != nil {
		return err
	}
	return nil
}

func writeVerificationFile(filePath string, hashDC string, dateTimeDC int, hashDR string, dateTimeDR int, verificationTime int64, verificationCheck bool) error {
	// Open file for writing (create if not exist, append if exists)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %s", err)
	}
	defer file.Close()

	// Format the line according to the specified format
	line := fmt.Sprintf("%d %s %d %s %d %t\n", dateTimeDC, hashDC, dateTimeDR, hashDR, verificationTime, verificationCheck)

	// Write formatted line to file
	if _, err := file.WriteString(line); err != nil {
		return fmt.Errorf("failed to write to file: %s", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	// Here you can define flags and configuration settings for verifyCmd
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
