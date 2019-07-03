package module

import "github.com/monologid/fast-cli/config"

// Create is used to create a FAST module project folder
func Create(name string) error {
	conf := &config.Config{ModName: name}

	err := conf.CreateModFolder()
	if err != nil {
		return err
	}

	err = conf.CreateConfigFile()
	if err != nil {
		return err
	}

	err = conf.CreateMainModFile()
	if err != nil {
		return err
	}

	return nil
}
