package unio

import (
    "github.com/go-bongo/bongo"
    "github.com/labstack/echo"
    "net/http"
)

// Shortcut to create a handler with search
func (h *Handler) FindAll(collection *bongo.Collection, rule RequestFormatRule) echo.HandlerFunc {
    return func(c echo.Context) (err error) {
        query := Searchs.GetQuery(c)
        search := Searchs.SearchFormat(query, rule)
        record := Searchs.RunQuery(search, collection)
        return c.JSON(Utils.RequestResult(http.StatusOK, record, err))
    }
}