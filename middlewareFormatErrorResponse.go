package unio

import (
    "errors"
    "github.com/labstack/echo"
    "net/http"
)

// Middleware to format when returns an error object
func (m *Middleware) FormatErrorResponse(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) (err error) {
        err = next(c)
        if err != nil {
            status := http.StatusUnauthorized
            httpError, ok := err.(*echo.HTTPError); if ok {
                status = httpError.Code
                err = errors.New(httpError.Message.(string))
            }
            return c.JSON(Utils.RequestResult(status, nil, err))
        }
        return err
    }
}