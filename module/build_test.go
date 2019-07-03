package module_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/monologid/fast-cli/module"
	"github.com/stretchr/testify/assert"
)

var modnamebuild = "test-build"

func tearDownForBuild() {
	os.RemoveAll("./" + modnamebuild)
	os.RemoveAll("./releases")
}

func TestBuildShouldReturnErrorIfVersionFileNotFound(t *testing.T) {
	_ = module.Create(modnamebuild)
	os.RemoveAll("./" + modnamebuild + "/.version")

	err := module.Build("")
	assert.Error(t, err)

	tearDownForBuild()
}

func TestBuildShouldReturnErrorIfVersionIsNotANumber(t *testing.T) {
	modnamebuild = modnamebuild + "-1"
	_ = module.Create(modnamebuild)

	filepath := fmt.Sprintf("%s/.version", modnamebuild)
	ioutil.WriteFile(filepath, []byte("test"), 0755)

	err := module.Build(modnamebuild)
	assert.Error(t, err)

	tearDownForBuild()
}

func TestBuildShouldReturnSuccess(t *testing.T) {
	modnamebuild = modnamebuild + "-2"
	_ = module.Create(modnamebuild)

	err := module.Build(modnamebuild)
	assert.Error(t, err)

	tearDownForBuild()
}
