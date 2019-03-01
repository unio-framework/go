package unio

import (
    "github.com/labstack/echo"
    "github.com/labstack/gommon/log"
    "net/url"
    "reflect"
)

func (s *Search) SearchFormat(query JSON, rule RequestFormatRule) JSONObject {
    reflectQuery := reflect.ValueOf(query)

    formattedQuery := JSONObject{}
    for _,k := range reflectQuery.MapKeys() {
        key := k.String()

        switch key {
        case "filter":
            formattedQuery[key] = FormatFilters(reflectQuery.MapIndex(k).Interface(), rule)
        //case "result":
        //    formattedQuery[key] = reflectQuery.MapIndex(k).Interface()
        }
    }
    return formattedQuery
}

func (s *Search) GetQuery(c echo.Context) JSON {
    rawQuery := c.QueryString()
    query,err := url.PathUnescape(rawQuery); if err != nil {
        log.Error(err)
    }

    if Utils.IsJSON(query) {
        jsonQuery := Utils.JSONParse(query)

        reflectJsonQuery := reflect.ValueOf(jsonQuery)
        switch reflectJsonQuery.Kind() {
        case reflect.Slice, reflect.Array:
            return reflectJsonQuery.Index(0).Interface()
        }
        return jsonQuery
    } else {
        urlQuery,err := url.ParseQuery(query); if err != nil {
            log.Error(err)
        }

        formattedQuery := JSONObject{}
        for key,value := range urlQuery {
            formattedQuery[key] = Utils.JSONParse(value[0])
        }
        return formattedQuery
    }
}