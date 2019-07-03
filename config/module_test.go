package config_test

import (
	"testing"

	"github.com/monologid/fast-cli/config"
	"github.com/stretchr/testify/assert"
)

func TestReadModuleConfShouldReturnErrorIfFileNotFound(t *testing.T) {
	_, err := config.ReadModuleConf("")
	assert.Error(t, err)
}

func TestReadModuleConfShouldReturnSuccess(t *testing.T) {
	tearUp()

	conf := &config.Config{
		ModName: modname,
	}

	_ = conf.CreateModFolder()
	_ = conf.CreateMainModFile()
	_ = conf.CreateConfigFile()

	mod, err := config.ReadModuleConf(modname)
	assert.NoError(t, err)
	assert.Equal(t, modname, mod.Mod.Name)

	tearDown()
}
