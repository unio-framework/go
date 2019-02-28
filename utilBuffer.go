package unio

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
)

func (u Utils) GetBuffer(c echo.Context) (buffer []byte) {
	buffer, err := ioutil.ReadAll(c.Request().Body); if err != nil {
		u.TraceError(err)
	}
	u.ResetBuffer(c, buffer)
	return
}

func (u Utils) ResetBuffer(c echo.Context, buffer []byte) {
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(buffer))
}

func (u Utils) InterfaceToBuffer(data interface{}) []byte {
	buffer, err := json.Marshal(data); if err != nil {
		panic(err)
	}
	return buffer
}
