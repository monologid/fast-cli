package config_test

import (
	"os"
	"testing"

	"github.com/monologid/fast-cli/config"
	"github.com/stretchr/testify/assert"
)

const modname = "register"

func tearUp() {
	_ = os.Mkdir(modname, 0755)
}

func tearDown() {
	_ = os.RemoveAll("./" + modname)
}

func TestCreateModFolderShouldReturnErrorIfModNameIsInvalid(t *testing.T) {
	conf := &config.Config{}
	err := conf.CreateModFolder()
	assert.Error(t, err)
}

func TestCreateModFolderShouldReturnSuccess(t *testing.T) {
	conf := &config.Config{
		ModName: modname,
	}
	err := conf.CreateModFolder()
	assert.NoError(t, err)

	tearDown()
}

func TestCreateConfigFileShouldReturnErrorIfModNameIsInvalid(t *testing.T) {
	conf := &config.Config{}
	err := conf.CreateConfigFile()
	assert.Error(t, err)
}

func TestCreateConfigFileShouldReturnSuccess(t *testing.T) {
	tearUp()

	conf := &config.Config{
		ModName: modname,
	}

	err := conf.CreateConfigFile()
	assert.NoError(t, err)

	_, err = os.Stat("./" + modname + "/config.yaml")
	isExist := os.IsNotExist(err)
	assert.False(t, isExist)

	tearDown()
}

func TestCreateMainModFileShouldReturnErrorIfModNameIsInvalid(t *testing.T) {
	conf := &config.Config{}
	err := conf.CreateMainModFile()
	assert.Error(t, err)
}

func TestCreateMainModFileShouldReturnSuccess(t *testing.T) {
	tearUp()

	conf := &config.Config{
		ModName: modname,
	}

	err := conf.CreateMainModFile()
	assert.NoError(t, err)

	_, err = os.Stat("./" + modname + "/main.go")
	isExist := os.IsNotExist(err)
	assert.False(t, isExist)

	tearDown()
}
