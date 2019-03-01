package tests

import (
	"github.com/unio-framework/go"
	"testing"
)

func TestJsonSearchQuery(t *testing.T) {
	stringQuery := "{\"filter\":{\"packageName\":{\"in\":[\"com.unio.test\"]}}}"
	query, _ := unio.Utils.JSONParse(stringQuery)
	want := unio.JSONObject{
		"filter": map[string]interface{}{
			"packageName": map[string]interface{}{
				"$in": []string{
					"com.unio.test",
				},
			},
		},
	}
	got := unio.Searchs.SearchFormat(query, nil)
	result(t, want["filter"] != nil, got["filter"] != nil)
}

func TestStringSearchQuery(t *testing.T) {
	stringQuery := "{\"packageName\":{\"in\":[\"com.unio.test\"]}}"
	query, _ := unio.Utils.JSONParse(stringQuery)
	js := unio.JSONObject{
		"filter": query,
	}
	want := unio.JSONObject{
		"filter": map[string]interface{}{
			"packageName": map[string]interface{}{
				"$in": []string{
					"com.unio.test",
				},
			},
		},
	}
	got := unio.Searchs.SearchFormat(js, nil)
	result(t, want["filter"] != nil, got["filter"] != nil)
}
