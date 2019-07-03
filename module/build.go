package module

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Build is used to build FAST module project
func Build(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return nil
	}

	tempName := strings.Split(dir, "/")
	modname := tempName[len(tempName)-1]

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

	buildcmd := exec.Command("go", "build", "-buildmode=plugin", "-o", fmt.Sprintf("./%s/%s.v%d.so", releaseFolderPath, modname, version), "main.go")
	errBuildCmd := buildcmd.Run()
	if errBuildCmd != nil {
		return errors.New("go-build " + errBuildCmd.Error())
	}

	if errUpdateVersionFile := ioutil.WriteFile(".version", []byte(strconv.Itoa(version)), 0755); errUpdateVersionFile != nil {
		return errors.New("update-version " + errUpdateVersionFile.Error())
	}

	return nil
}
