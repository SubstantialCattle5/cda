/*
Copyright © 2024 NILAY SHARAN nilaynathsharan16@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cda",
	Short: "Automates the process of image drilling in snapshot transfer from DataCenter(DC) to DataRecovery(DR).",
	Long: `cephdrillautomation is a CLI library for Ceph that automate the process of snapshot transfer from DC to DR.
This application is a tool to generate the snaps.`,
	// Uncomment the following line if your bare application
	//// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	var poolName = ""
	//	var vmName = ""
	//	if len(args) >= 1 && args[0] != "" {
	//		poolName = args[0]
	//		vmName = args[1]
	//	}
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cephdrillautomation.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
