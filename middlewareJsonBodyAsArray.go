package unio

import (
	"encoding/json"
	"github.com/labstack/echo"
	"reflect"
)

func (m Middleware) JsonBodyAsArray(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if c.Request().Header.Get(echo.HeaderContentType) != echo.MIMEApplicationJSON {
			// Conversion will be done only for JSON request
			return next(c)
		}

		buffer := utils.GetBuffer(c)
		var raw interface{}
		err = c.Bind(&raw)

		if reflect.ValueOf(raw).Kind() != reflect.Slice {
			body := []interface{}{raw}
			buffer, err = json.Marshal(body); if err != nil {
				panic(err)
			}
		}

		utils.ResetBuffer(c, buffer)
		return next(c)
	}
}
