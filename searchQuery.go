// Copyright 2019 Leandro Akira Omiya Takagi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package unio

import (
    "github.com/go-bongo/bongo"
    "github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/url"
	"reflect"
)

// Format $_GET string query to JSON structure
// Uses bellow pattern:
// - search: MongoDB custom formatted query
// - filter: List of fields to remove from response
// - result: List of fields to return from response
// - populate: List of fields can be populated
func (s *Search) GetQuery(c echo.Context) JSON {
	rawQuery := c.QueryString()
	query, err := url.PathUnescape(rawQuery)
	if err != nil { log.Error(err) }

	if Utils.IsJSON(query) {
		jsonQuery, _ := Utils.JSONParse(query)

		reflectJsonQuery := reflect.ValueOf(jsonQuery)
		switch reflectJsonQuery.Kind() {
		case reflect.Slice, reflect.Array:
			return reflectJsonQuery.Index(0).Interface()
		}
		return jsonQuery
	} else {
		urlQuery, err := url.ParseQuery(query)
		if err != nil {
			log.Error(err)
		}

		formattedQuery := JSONObject{}
		for key, value := range urlQuery {
			formattedQuery[key], _ = Utils.JSONParse(value[0])
		}
		return formattedQuery
	}
}

// Format query filter parameters with MongoDB pattern
func (s *Search) SearchFormat(query JSON, rule RequestFormatRule) JSONObject {
    reflectQuery := reflect.ValueOf(query)

    formattedQuery := JSONObject{}
    for _, k := range reflectQuery.MapKeys() {
        key := k.String()

        switch key {
        case "search":
            formattedQuery[key] = s.FormatFilters(reflectQuery.MapIndex(k).Interface(), rule)
        case "filter", "result", "populate":
            formattedQuery[key] = reflectQuery.MapIndex(k).Interface()
        }
    }
    return formattedQuery
}

// Run MongoDB search and result filtering
func (s *Search) RunQuery(collection *bongo.Collection, search JSONObject, populate SearchPopulate) []interface{} {
    records := make([]interface{}, 0)
    if search["search"] != nil {
        var model map[string]interface{}
        results := collection.Find(search["search"])
        for results.Next(&model) {
            filterResult(search, &model)
            populateFields(search, model, populate)
            records = append(records, model)
        }
        defer collection.Connection.Session.Close()
    }
    return records
}

// Filter result fields
func filterResult(search JSONObject, model *map[string]interface{}) {
    // Filter fields that not returns to result
    if search["filter"] != nil {
        f, ok := search["filter"].([]interface{}); if ok {
            for _,key := range f {
                s, ok := key.(string); if ok {
                    delete(*model, s)
                }
            }
        }
    }

    // Get all results
    if search["result"] != nil {
        f, ok := search["result"].([]interface{}); if ok {
            for key := range *model {
                if Utils.ArrayContains(f, key) == false {
                    delete(*model, key)
                }
            }
        }
    }
}

// Populate result fields
func populateFields(search JSONObject, model map[string]interface{}, populate SearchPopulate) {
    if search["populate"] != nil && populate != nil {
        f, ok := search["populate"].([]interface{}); if ok {
            for _,k := range f {
                key, ok := k.(string); if ok && Utils.MapKeyExists(model, key) && model[key] != nil {
                    result, err := populate(key, model[key])
                    if err == nil && result != nil { model[key] = result }
                }
            }
        }
    }
}