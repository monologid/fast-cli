package cmd

import (
	"github.com/monologid/fast-cli/log"
	"github.com/monologid/fast-cli/module"
	"github.com/spf13/cobra"
)

var projectPath string

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build FAST module project",
	Run:   runBuildCmd,
}

func runBuildCmd(cmd *cobra.Command, args []string) {
	if len(projectPath) == 0 {
		log.Print("make sure that current folder is in the module project folder")
	}

	err := module.Build(projectPath)
	if err != nil {
		log.PrintErr("failed to build module", err)
		return
	}

	log.Print("module has been build successfully")
}

func init() {
	buildCmd.PersistentFlags().StringVarP(&projectPath, "path", "p", "", "set project module folder path")
	rootCmd.AddCommand(buildCmd)
}
