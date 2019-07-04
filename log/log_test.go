package log_test

import (
	"errors"
	"testing"

	"github.com/monologid/fast-cli/log"
	"github.com/monologid/fast-cli/util"
	"github.com/stretchr/testify/assert"
)

func TestGetDateTimeShouldNotReturnEmptyString(t *testing.T) {
	datetime := util.GetDateTime()

	assert.NotEqual(t, 0, len(datetime))
}

func TestShouldPrintErrorLog(t *testing.T) {
	log.PrintErr("something", errors.New("went wrong"))
}

func TestShouldPrintLog(t *testing.T) {
	log.Print("something")
}
