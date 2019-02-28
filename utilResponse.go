package unio

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (u Util) RequestResult(c echo.Context, status int, data interface{}, errors interface{}) error {
	result := echo.Map{
		"status": status,
	}

	if status == http.StatusOK {
		result["status"] = 1
		result["data"] = data
		if errors != nil {
			result["error_message"] = errors
		}
	} else {
		result["status"] = 0
		result["error_message"] = data
	}

	return c.JSON(status, result)
}

func (u Util) TraceError(err error) {
	if err != nil {
		log.Error(err)
	}
}
