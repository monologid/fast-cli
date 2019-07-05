package module

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/monologid/fast-cli/config"
	"github.com/monologid/fast-cli/util"
)

// Apply is used to upload and register a release module to FAST platform
func Apply(path string) error {
	modconf, err := config.ReadModuleConf(path)
	if err != nil {
		return errors.New("read-mod-config " + err.Error())
	}

	fastUploadURL := modconf.FastConfig.UploadURL
	fastRegisterModURL := modconf.FastConfig.RegisterURL

	errUpload := UploadMod(fastUploadURL, path, modconf.Mod)
	if errUpload != nil {
		return errors.New("upload-mod " + errUpload.Error())
	}

	errRegister := RegisterMod(fastRegisterModURL, path, modconf.Mod)
	if errRegister != nil {
		return errors.New("register-mod " + errRegister.Error())
	}

	return nil
}

// UploadMod is used to upload released module
func UploadMod(url string, path string, mod config.Mod) error {
	modname := strings.ToLower(mod.Name)
	releaseVer := strings.ToLower(mod.ReleaseVersion)

	modfilepath := fmt.Sprintf("releases/%s.%s.so", modname, releaseVer)
	if len(path) != 0 {
		modfilepath = fmt.Sprintf("%s/%s", path, modfilepath)
	}

	fmt.Println(modfilepath)

	file, errFile := os.Open(modfilepath)
	if errFile != nil {
		return errors.New("open-file " + errFile.Error())
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, errCreateFormFile := writer.CreateFormFile("file", filepath.Base(modfilepath))
	if errCreateFormFile != nil {
		return errors.New("create-form-file " + errCreateFormFile.Error())
	}

	_, errCopyFile := io.Copy(part, file)
	writer.Close()
	if errCopyFile != nil {
		return errors.New("copy-file " + errCopyFile.Error())
	}

	req, errNewRequest := http.NewRequest(http.MethodPost, url, body)
	if errNewRequest != nil {
		return errors.New("create-new-request-for-upload " + errNewRequest.Error())
	}

	authKey := fmt.Sprintf("Bearer %s", mod.SecretKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set(config.XFastSecretKeyHeader, authKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("call-api-for-upload " + err.Error())
	}
	defer resp.Body.Close()

	parser := util.NewParser()
	parser.SetResponseBody(resp.Body)

	status, err := parser.Status()
	if err != nil {
		return errors.New("parser-resp-for-upload " + err.Error())
	}

	if strings.ToUpper(status) == "ERROR" || strings.ToUpper(status) == "FAILED" {
		return errors.New("failed to upload module")
	}

	return nil
}

// RegisterMod is used to register released module
func RegisterMod(url string, path string, mod config.Mod) error {
	body, err := json.Marshal(mod)
	if err != nil {
		return errors.New("marshal-body-for-register-mod " + err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return errors.New("create-new-request-for-register " + err.Error())
	}

	authorizationToken := fmt.Sprintf("Bearer %s", mod.SecretKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(config.XFastSecretKeyHeader, authorizationToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("call-api-for-register " + err.Error())
	}
	defer resp.Body.Close()

	parser := util.NewParser()
	parser.SetResponseBody(resp.Body)

	status, err := parser.Status()
	if err != nil {
		return errors.New("parser-resp-for-upload " + err.Error())
	}

	if strings.ToUpper(status) == "ERROR" || strings.ToUpper(status) == "FAILED" {
		return errors.New("failed to upload module")
	}

	return nil
}
