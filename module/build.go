package module

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"github.com/monologid/fast-cli/config"
)

// Build is used to build FAST module project
func Build(path string) error {
	modconf, err := config.ReadModuleConf(path)
	if err != nil {
		return errors.New("read-mod-config " + err.Error())
	}

	releaseFolderPath := "releases"
	versionFilePath := ".version"

	if len(path) != 0 {
		releaseFolderPath = fmt.Sprintf("%s/releases", path)
		versionFilePath = fmt.Sprintf("%s/.version", path)
	}

	if _, errCheckFolder := os.Stat(releaseFolderPath); os.IsNotExist(errCheckFolder) {
		if errCreateFolder := os.Mkdir(releaseFolderPath, 0755); errCreateFolder != nil {
			return errors.New("create-release-folder " + errCreateFolder.Error())
		}
	}

	dataVersion, errReadVersion := ioutil.ReadFile(versionFilePath)
	if errReadVersion != nil {
		return errors.New("read-version " + errReadVersion.Error())
	}

	version, errParseVersion := strconv.Atoi(string(dataVersion))
	if errParseVersion != nil {
		return errParseVersion
	}

	version = version + 1

	outputfile := fmt.Sprintf("%s/%s.v%d.so", releaseFolderPath, modconf.Mod.Name, version)
	mainfile := fmt.Sprintf("%s/main.go", path)

	buildcmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outputfile, mainfile)
	errBuildCmd := buildcmd.Run()
	if errBuildCmd != nil {
		return errors.New("go-build " + errBuildCmd.Error())
	}

	if errUpdateVersionFile := ioutil.WriteFile(".version", []byte(strconv.Itoa(version)), 0755); errUpdateVersionFile != nil {
		return errors.New("update-version " + errUpdateVersionFile.Error())
	}

	return nil
}
