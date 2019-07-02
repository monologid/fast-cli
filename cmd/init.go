package cmd

import (
	"fmt"

	"github.com/monologid/fast-cli/log"
	"github.com/monologid/fast-cli/module"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiate new FAST module project",
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.PrintErr("please define module name", nil)
		return
	}

	modname := args[0]
	err := module.Create(modname)
	if err != nil {
		log.PrintErr("failed to create new module", err)
		return
	}

	log.Print(fmt.Sprintf("New module `%s` has been created", modname))
}

func init() {
	rootCmd.AddCommand(initCmd)
}
