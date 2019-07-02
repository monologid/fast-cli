package module_test

import (
	"os"
	"testing"

	"github.com/monologid/fast-cli/module"
	"github.com/stretchr/testify/assert"
)

var modname = "test"

func tearDown() {
	os.RemoveAll("./" + modname)
}

func TestCreateShouldReturnErrorIfModNameIsInvalid(t *testing.T) {
	err := module.Create("")
	assert.Error(t, err)
}

func TestCreateShouldReturnSuccess(t *testing.T) {
	err := module.Create(modname)
	assert.NoError(t, err)

	errFolderExist := module.Create(modname)
	assert.Error(t, errFolderExist)

	tearDown()
}
