/*
Copyright © 2024 NILAY SHARAN nilaynathsharan16@gmail.com
*/

package cmd

import (
	"cephdrillautomation/util"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"regexp"
	"slices"
	"strings"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "This command generates a hash of the given snapshot",
	Long: `This command generates a hash of the given snapshot and saves it in the local log file.
`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		var rbdCommand, poolName, vmName string
		var err error
		scriptPath := "functions/dccheck.sh"

		if len(args) > 0 && strings.TrimSpace(args[0]) != "" {
			rbdCommand = strings.TrimSpace(args[0])
			if len(args) > 1 {
				poolName = strings.TrimSpace(args[1])
			}
			if len(args) > 2 {
				vmName = strings.TrimSpace(args[2])
			}
		}

		if vmName == "" || poolName == "" || !slices.Contains([]string{"export", "export-diff"}, rbdCommand) {
			var errorMsg string
			if !slices.Contains([]string{"export", "export-diff"}, rbdCommand) {
				errorMsg += "RBD command is wrong\n"
			}
			if vmName == "" {
				errorMsg += "VM Snapshot name is missing\n"
			}
			if poolName == "" {
				errorMsg += "Pool Name is missing\n"
			}
			err = errors.New(errorMsg)
			log.Fatalf("Errors: %s", err)
			return
		}

		fmt.Printf("Running export and checksums for %s/%s\n", poolName, vmName)
		fmt.Printf("Splitting based on 10 of the memory size \n")
		// Execute the shell script
		scriptCmd := exec.Command("bash", scriptPath, rbdCommand, poolName, vmName)
		output, err := scriptCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Failed to execute script: %v\nOutput: %s", err, output)
		}
		hash, err := util.ScriptValidationCheck(strings.TrimSpace(string(output)))
		if err != nil {
			log.Fatalf("Script validation failed: %v", err)
		}
		answers := strings.Split(hash, "\n")
		// Print the output of the script
		answer := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(answers[len(answers)-1], "")
		fmt.Printf("%s\n", answer)
		err = util.SaveFile(rbdCommand, vmName, poolName, answer)
		if err != nil {
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
