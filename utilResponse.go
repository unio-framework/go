package unio

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
    "reflect"
)

func (u *Util) RequestResult(c echo.Context, status int, data interface{}, errors interface{}) error {
    result := echo.Map{
        "status": status,
    }
    errors = formatErrors(errors)

    if status == http.StatusOK {
        result["status"] = 1
        result["data"] = data
        if errors != nil {
            result["error"] = errors
        }
    } else {
        result["status"] = 0
        result["error"] = u.AnyOf(data, errors)
    }

    return c.JSON(status, result)
}

func (u *Util) TraceError(err error) {
	if err != nil {
		log.Error(err)
	}
}

func formatErrors(errs interface{}) interface{} {
    //noinspection ALL
    errors := []interface{}{}
    if errs == nil { return errors }

    if Utils.In(reflect.ValueOf(errs).Kind(), reflect.Array, reflect.Slice) == false {
        errors = []interface{}{errs}
    }
    for key,value := range errors {
        if reflect.ValueOf(value).Kind() == reflect.String {
            errors[key] = value
        } else {
            errors[key] = value.(error).Error()
        }
    }
    return errors
}