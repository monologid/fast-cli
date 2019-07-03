package module_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/monologid/fast-cli/config"
	"github.com/monologid/fast-cli/module"
	"github.com/stretchr/testify/assert"
)

var modnameapply = "testapply"

func tearDownForApply() {
	os.RemoveAll("./" + modnameapply)
}

func TestApplyShouldReturnErrorIfPathIsNotValid(t *testing.T) {
	err := module.Apply("xyz")
	assert.Error(t, err)
}

func TestApplyShouldReturnErrorOnUploadWhenFileIsNotFound(t *testing.T) {
	_ = module.Create(modnameapply)

	err := module.Apply(modnameapply)
	assert.Error(t, err)

	tearDownForApply()
}

func TestApplyShouldReturnErrorOnUploadWhenOpenFile(t *testing.T) {
	_ = module.Create(modnameapply)
	_ = module.Build(modnameapply)

	err := module.Apply(modnameapply)
	assert.Error(t, err)

	tearDownForApply()
}

func TestApplyShouldReturnErrorOnUpload(t *testing.T) {
	_ = module.Create(modnameapply)
	_ = module.Build(modnameapply)

	url := "something"
	dir, _ := os.Getwd()
	mod := config.Mod{
		Name:           modnameapply,
		ReleaseVersion: "v1",
	}

	err := module.UploadMod(url, dir+"/"+modnameapply, mod)
	assert.Error(t, err)

	tearDownForApply()
}

func TestApplyShouldReturnErrorWhenParseJSONOnUpload(t *testing.T) {
	_ = module.Create(modnameapply)
	_ = module.Build(modnameapply)

	dir, _ := os.Getwd()
	mod := config.Mod{
		Name:           modnameapply,
		ReleaseVersion: "v1",
	}

	resp := `{"status":"failed"`
	svr := mockServer(resp)

	err := module.UploadMod(svr.URL, dir+"/"+modnameapply, mod)
	assert.Error(t, err)

	tearDownForApply()
}

func TestApplyShouldReturnFailedOnUpload(t *testing.T) {
	_ = module.Create(modnameapply)
	_ = module.Build(modnameapply)

	dir, _ := os.Getwd()
	mod := config.Mod{
		Name:           modnameapply,
		ReleaseVersion: "v1",
	}

	resp := `{"status":"failed"}`
	svr := mockServer(resp)

	err := module.UploadMod(svr.URL, dir+"/"+modnameapply, mod)
	assert.Error(t, err)

	tearDownForApply()
}

func TestApplyShouldReturnSuccessOnUpload(t *testing.T) {
	_ = module.Create(modnameapply)
	_ = module.Build(modnameapply)

	dir, _ := os.Getwd()
	mod := config.Mod{
		Name:           modnameapply,
		ReleaseVersion: "v1",
	}

	resp := `{"status":"success"}`
	svr := mockServer(resp)

	err := module.UploadMod(svr.URL, dir+"/"+modnameapply, mod)
	assert.NoError(t, err)

	tearDownForApply()
}

func TestApplyShouldReturnErrorWhenParseJSONOnRegister(t *testing.T) {
	_ = module.Create(modnameapply)
	_ = module.Build(modnameapply)

	dir, _ := os.Getwd()
	mod := config.Mod{
		Name:           modnameapply,
		ReleaseVersion: "v1",
	}

	resp := `{"status":"failed"`
	svr := mockServer(resp)

	err := module.RegisterMod(svr.URL, dir+"/"+modnameapply, mod)
	assert.Error(t, err)

	tearDownForApply()
}

func TestApplyShouldReturnFailedOnRegisterMod(t *testing.T) {
	_ = module.Create(modnameapply)
	_ = module.Build(modnameapply)

	dir, _ := os.Getwd()
	mod := config.Mod{
		Name:           modnameapply,
		ReleaseVersion: "v1",
	}

	resp := `{"status":"failed"}`
	svr := mockServer(resp)

	err := module.RegisterMod(svr.URL, dir+"/"+modnameapply, mod)
	assert.Error(t, err)

	tearDownForApply()
}

func TestApplyShouldReturnSuccessOnRegisterMod(t *testing.T) {
	_ = module.Create(modnameapply)
	_ = module.Build(modnameapply)

	dir, _ := os.Getwd()
	mod := config.Mod{
		Name:           modnameapply,
		ReleaseVersion: "v1",
	}

	resp := `{"status":"success"}`
	svr := mockServer(resp)

	err := module.RegisterMod(svr.URL, dir+"/"+modnameapply, mod)
	assert.NoError(t, err)

	tearDownForApply()
}

func mockServer(jsonres string) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(jsonres))
	}

	return httptest.NewServer(http.HandlerFunc(handler))
}
