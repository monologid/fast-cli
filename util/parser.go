package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Parser is an object for parser
type Parser struct {
	ResponseBody io.ReadCloser
	Data         map[string]interface{}
}

// NewParser initiate new parser object
func NewParser() *Parser {
	return &Parser{
		Data: make(map[string]interface{}),
	}
}

// SetResponseBody is used to set response body receive from http request call
func (p *Parser) SetResponseBody(resp io.ReadCloser) {
	p.ResponseBody = resp
}

// Status is used to get status response
func (p *Parser) Status() (string, error) {
	respbody, err := ioutil.ReadAll(p.ResponseBody)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(respbody, &p.Data)
	if err != nil {
		return "", err
	}

	return p.Data["status"].(string), nil
}
