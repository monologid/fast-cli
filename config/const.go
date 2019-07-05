package config

// FastUploadModuleURL is the default upload url to FAST platform
const FastUploadModuleURL = `http://fast.monolog.id/api/v1/application/module/upload`

// FastURegisterModuleURL is the default url for registering module to FAST platform
const FastURegisterModuleURL = `http://fast.monolog.id/api/v1/application/module`

// XFastSecretKeyHeader is the key header for auth
const XFastSecretKeyHeader = "x-fast-secret-key"

// DefaultMainModuleFile is the default FAST module file
const DefaultMainModuleFile = `
package main

import (
  "errors"
  "net/http"
)

// Define any attributes needed for your API in module struct
type module struct{}

// Define the http method for this module. There are several http methods available, such as:
// GET, POST, PUT, PATCH, DELETE
const method = "GET"

// Define the endpoint path for this module. The path is used for accessing this module.
const path = "/api/v1/sample"

// Method returns http method
// WARNING! DO NOT CHANGE THIS FUNCTION
func (m *module) Method() string {
	return method
}

// Path returns endpoint path
// WARNING! DO NOT CHANGE THIS FUNCTION
func (m *module) Path() string {
	return path
}

// Before is commonly used for validating request header
// For example, validation authorization header, any request body validation, etc
// If there's no validation, please return nil for both interface and error
func (m *module) Before(header http.Header, req map[string]interface{}) (interface{}, error) {
	return nil, nil
}

// Execute returns the result after module execution
func (m *module) Execute(req map[string]interface{}) (interface{}, error) {
	return nil, errors.New("please implement")
}

// You can add more method into module struct according to your requirement.
// Then you can access it in Execute(...) function.
// For example:
//
// func (m *module) checkRequestBody(header http.Header, req map[string]interface{}) error {
// 		if len(header.Get("Authorization")) == 0 {
//				... do something ...
//		}
//
//    return nil
// }

var FastModule module
`

// DefaultModuleConfig is the default config for a module
var DefaultModuleConfig = `
# This is the module config file
# Please replace all necessary field based on its description

mod:

# Module name will be generated. Don't change the value
  name: {{MOD_NAME}}

# Describe the feature of you module
  description: Module description

# Replace the value with your MONOLOG account id
  accountId: xxxxx

# Replace the value with your FAST application secret key
  secretKey: xxxxx

# Replace the value with your intended release version
# This will be used for applying module to FAST platform
  releaseVersion: v1

fast:

# (Optional) Uncomment uploadUrl when you want to upload to sandbox environment,
# otherwise remove this config key if you want to use the defauklt production upload mod url
  uploadUrl: https://fast.monolog.id/api/v1/application/module/upload

# (Optional) Uncomment registerUrl when you want to register your mod into sandbox environment,
# otherwise remove this config key if you want to use the defauklt production upload mod url
  registerUrl: https://fast.monolog.id/api/v1/application/module
`
