package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Config contains field available for config package
type Config struct {
	ModName string
}

// CreateModFolder is used to create module folder
func (c *Config) CreateModFolder() error {
	if len(c.ModName) == 0 {
		return errors.New("invalid module name")
	}

	err := os.Mkdir(c.ModName, 0755)
	if err != nil {
		return errors.New("failed to create module folder, err=" + err.Error())
	}

	return nil
}

// CreateConfigFile is used to create default module config file
func (c *Config) CreateConfigFile() error {
	if len(c.ModName) == 0 {
		return errors.New("invalid module name")
	}

	defaultModConfig := strings.ReplaceAll(DefaultModuleConfig, "{{MOD_NAME}}", c.ModName)
	filepath := fmt.Sprintf("%s/%s", c.ModName, "config.yaml")
	if errCreateFile := ioutil.WriteFile(filepath, []byte(defaultModConfig), 0755); errCreateFile != nil {
		return errors.New("failed to create default module config file, err=" + errCreateFile.Error())
	}

	filepath = fmt.Sprintf("%s/.version", c.ModName)
	if errCreateVersionFile := ioutil.WriteFile(filepath, []byte("0"), 0755); errCreateVersionFile != nil {
		return errors.New("failed to create version module config file, err=" + errCreateVersionFile.Error())
	}

	return nil
}

// CreateMainModFile is used to create main module file
func (c *Config) CreateMainModFile() error {
	if len(c.ModName) == 0 {
		return errors.New("invalid module name")
	}

	filepath := fmt.Sprintf("%s/%s", c.ModName, "main.go")
	if errCreateFile := ioutil.WriteFile(filepath, []byte(DefaultMainModuleFile), 0755); errCreateFile != nil {
		return errors.New("failed to create main module file, err=" + errCreateFile.Error())
	}

	return nil
}
