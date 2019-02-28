package unio

import (
	"fmt"
	"github.com/labstack/echo"
	"reflect"
)

/**
Rule structure
*/
type JsonFormatRule func(key string, value interface{}) interface{}

/**
Middleware
Run all JSON body fields, and format the need
*/
func (m Middleware) JsonFormatFields(formatter JsonFormatRule) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if c.Request().Header.Get(echo.HeaderContentType) != echo.MIMEApplicationJSON {
				// Conversion will be done only for JSON request
				return next(c)
			}

			buffer := Utils.GetBuffer(c)
			var raw interface{}
			err = c.Bind(&raw)

			switch reflect.ValueOf(raw).Kind() {
            //noinspection ALL
            case reflect.Slice:
				s := reflect.ValueOf(raw)

				body := []interface{}{}
				for i := 0; i < s.Len(); i++ {
					part := s.Index(i)
					formatted := format(part.Interface(), formatter)
					body = append(body, formatted)
				}
				buffer = Utils.InterfaceToBuffer(body)
			default:
				body := format(raw, formatter)
				buffer = Utils.InterfaceToBuffer(body)
			}

			Utils.ResetBuffer(c, buffer)
			return next(c)
		}
	}
}

/**
Formatter workaround
*/
func format(raw interface{}, formatter JsonFormatRule) interface{} {
	reflectRaw := reflect.ValueOf(raw)

	body := map[string]interface{}{}

	for _, k := range reflectRaw.MapKeys() {
		key := k.String()
		value := reflectRaw.MapIndex(k).Interface()

		switch reflect.ValueOf(value).Kind() {
		case reflect.Slice:
		case reflect.Array:
			s := reflect.ValueOf(value)
			for i := 0; i < s.Len(); i++ {
				part := s.Index(i)
				body[key] = format(part.Interface(), formatter)
			}

		case reflect.Map:
			body[key] = format(value, formatter)

		default:
			body[key] = formatter(key, value)
		}
	}
	return body
}
