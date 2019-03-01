package unio

import (
    "reflect"
)

func (s *Search) FormatFilters(f interface{}, rule RequestFormatRule) (filters JSONObject) {
    rFilters := reflect.ValueOf(f)
    filters = JSONObject{}

    if rFilters.Kind() == reflect.Map {
        for _,key := range rFilters.MapKeys() {
            value := formatFilter(rFilters.MapIndex(key).Interface())
            if value != nil { value = rule("GET", key.String(), value) }
            if value != nil {
                filters[key.String()] = value
            }
        }
    }

    return
}

func (s *Search) formatFilter(content interface{}) interface{} {
    if Utils.IsJSON(content) {
        rContent := reflect.ValueOf(content)
        if rContent.Kind() == reflect.Map {
            formattedContent := map[string]interface{}{}
            for _,key := range rContent.MapKeys() {
                fKey := filterMap[key.String()]
                formattedContent[fKey] = rContent.MapIndex(key).Interface()
            }
            return formattedContent
        }
    } else {
        return content
    }

    return nil
}