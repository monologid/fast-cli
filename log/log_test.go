package log_test

import (
	"testing"
	"errors"
	"github.com/monologid/fast-cli/log"
	"github.com/stretchr/testify/assert"
)

func TestGetDateTimeShouldNotReturnEmptyString(t *testing.T) {
	datetime := log.GetDateTime()

	assert.NotEqual(t, 0, len(datetime))
}

func TestShouldPrintErrorLog(t *testing.T) {
	log.PrintErr("something", errors.New("went wrong"))
}

func TestShouldPrintLog(t *testing.T) {
	log.Print("something")
}