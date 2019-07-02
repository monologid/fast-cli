package main

import (
	"fmt"

	"github.com/monologid/fast-cli/cmd"
)

func main() {
	err := cmd.RootInit()
	if err != nil {
		fmt.Println("failed to execute command, err=" + err.Error())
	}
}
