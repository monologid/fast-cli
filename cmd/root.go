package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fast-cli",
	Short: "",
}

// RootInit is the main root command for fast-cli
func RootInit() error {
	return rootCmd.Execute()
}
