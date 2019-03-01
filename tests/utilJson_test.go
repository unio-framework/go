package tests

import (
	"github.com/unio-framework/go"
	"testing"
)

func TestIfStringCanBeJSON(t *testing.T) {
	data := "{\"one\":\"two\"}"
	got := unio.Utils.IsJSON(data)
	result(t, true, got)
}

func TestIfStringCanNotBeJSON(t *testing.T) {
	data := "one=two"
	got := unio.Utils.IsJSON(data)
	result(t, false, got)
}

func TestIfMapCanBeJSON(t *testing.T) {
	data := map[string]string{
		"one":   "two",
		"three": "four",
	}
	got := unio.Utils.IsJSON(data)
	result(t, true, got)
}

func TestIfArrayCanBeJSON(t *testing.T) {
	data := []string{
		"one",
		"two",
	}
	got := unio.Utils.IsJSON(data)
	result(t, true, got)
}

func TestStringToJSONParse(t *testing.T) {
	data := "{\"one\":\"two\"}"
	want := map[string]interface{}{
		"one": "two",
	}
	got, _ := unio.Utils.JSONParse(data)
	result(t, want, got)
}
