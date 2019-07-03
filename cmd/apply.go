package cmd

import (
	"github.com/monologid/fast-cli/log"
	"github.com/monologid/fast-cli/module"
	"github.com/spf13/cobra"
)

var applyProjectPath string

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a release module to FAST platform",
	Run:   runApplyCmd,
}

func runApplyCmd(cmd *cobra.Command, args []string) {
	if len(applyProjectPath) == 0 {
		log.Print("make sure that current folder is in the module project folder")
	}

	err := module.Apply(applyProjectPath)
	if err != nil {
		log.PrintErr("failed to apply module", err)
		return
	}

	log.Print("module has been applied successfully")
}

func init() {
	applyCmd.PersistentFlags().StringVarP(&applyProjectPath, "path", "p", "", "set project module folder path")
	rootCmd.AddCommand(applyCmd)
}
