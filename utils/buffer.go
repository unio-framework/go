package utils

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
)

func GetBuffer(c echo.Context) (buffer []byte) {
	buffer, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		TraceError(err)
	}
	ResetBuffer(c, buffer)
	return
}

func ResetBuffer(c echo.Context, buffer []byte) {
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(buffer))
}

func InterfaceToBuffer(data interface{}) []byte {
	buffer, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return buffer
}
