package unio

import (
	"github.com/labstack/echo"
	"reflect"
    "strings"
)

/**
Middleware
Run all JSON body fields, and format the need
*/
func (m *Middleware) JsonFormatFields(formatter RequestFormatRule) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
		    mimeType := c.Request().Header.Get(echo.HeaderContentType)
		    // Use contains to make more accurate
			if !strings.Contains(mimeType, echo.MIMEApplicationJSON) {
				// Conversion will be done only for JSON request
				return next(c)
			}

			buffer := Utils.GetBuffer(c)
			var raw interface{}
			err = c.Bind(&raw)

			switch reflect.ValueOf(raw).Kind() {
            //noinspection ALL
            case reflect.Slice, reflect.Array:
				s := reflect.ValueOf(raw)

				body := []interface{}{}
				for i := 0; i < s.Len(); i++ {
					part := s.Index(i)
					formatted := format(c, part.Interface(), formatter)
					body = append(body, formatted)
				}
				buffer = Utils.InterfaceToBuffer(body)
			default:
				body := format(c, raw, formatter)
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
func format(c echo.Context, raw interface{}, formatter RequestFormatRule) interface{} {
	reflectRaw := reflect.ValueOf(raw)

	body := map[string]interface{}{}

	for _, k := range reflectRaw.MapKeys() {
		key := k.String()
		value := reflectRaw.MapIndex(k).Interface()

		switch reflect.ValueOf(value).Kind() {
		case reflect.Slice, reflect.Array:
			s := reflect.ValueOf(value)
			for i := 0; i < s.Len(); i++ {
				part := s.Index(i)
				body[key] = format(c, part.Interface(), formatter)
			}

		case reflect.Map:
			body[key] = format(c, value, formatter)

		default:
			body[key] = formatter(c.Request().Method, key, value)
		}
	}
	return body
}
