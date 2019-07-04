package log

import (
	"fmt"

	"github.com/monologid/fast-cli/util"
)

// PrintErr is used to print an error log
func PrintErr(text string, err error) {
	msg := fmt.Sprintf("%s", text)

	var errmsg string
	if err != nil {
		errmsg = err.Error()
		msg = fmt.Sprintf("%s, err=%s", msg, errmsg)
	}

	fmt.Printf("%s - \033[1;31mERROR:\033[0m %s\n", util.GetDateTime(), msg)
}

// Print is used to print a log
func Print(text string) {
	fmt.Printf("%s - INFO: %s\n", util.GetDateTime(), text)
}
