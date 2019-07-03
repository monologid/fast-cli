package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Module is a schema for default module config
type Module struct {
	Mod        Mod        `yaml:"mod"`
	FastConfig FastConfig `yaml:"fast"`
}

// Mod is a schema related to module project
type Mod struct {
	Name           string `yaml:"name" json:"name"`
	Description    string `yaml:"description" json:"description"`
	AccountID      string `yaml:"accountId" json:"account_id"`
	SecretKey      string `yaml:"secretKey" json:"secret_key"`
	ReleaseVersion string `yaml:"releaseVersion" json:"release_version"`
	Filename       string `json:"filename"`
}

// FastConfig is a schema related to fast platform configuration
type FastConfig struct {
	UploadURL   string `yaml:"uploadUrl"`
	RegisterURL string `yaml:"registerUrl"`
}

// ReadModuleConf is used to read default module config file
func ReadModuleConf(path string) (*Module, error) {
	moduleConfFilePath := "config.yaml"

	if len(path) != 0 {
		moduleConfFilePath = fmt.Sprintf("%s/config.yaml", path)
	}

	tempModConf, err := ioutil.ReadFile(moduleConfFilePath)
	if err != nil {
		return nil, err
	}

	var mod Module

	errUnmarshal := yaml.Unmarshal(tempModConf, &mod)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &mod, nil
}
