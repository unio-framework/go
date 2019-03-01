package tests

import (
	"github.com/unio-framework/go"
	"testing"
)

func TestStringArrayContainsIsValid(t *testing.T) {
	group := []string{"one", "two"}
	got := unio.Utils.ArrayContains(group, "one")
	result(t, true, got)
}

func TestStringArrayContainsIsInvalid(t *testing.T) {
	group := []string{"one", "two"}
	got := unio.Utils.ArrayContains(group, "three")
	result(t, false, got)
}

func TestIntContainsIsValid(t *testing.T) {
	group := []int{1, 2}
	got := unio.Utils.ArrayContains(group, 2)
	result(t, true, got)
}

func TestIntContainsIsInvalid(t *testing.T) {
	group := []int{1, 2}
	got := unio.Utils.ArrayContains(group, 3)
	result(t, false, got)
}

func TestMapContainsIsValid(t *testing.T) {
	toTry := map[string]string{
		"one": "two",
	}
	group := []interface{}{
		toTry,
		map[string]string{
			"three": "four",
		},
	}
	got := unio.Utils.ArrayContains(group, toTry)
	result(t, true, got)
}

func TestMapContainsIsInvalid(t *testing.T) {
	toTry := map[string]string{
		"five": "six",
	}
	group := []interface{}{
		map[string]string{
			"one": "two",
		},
		map[string]string{
			"three": "four",
		},
	}
	got := unio.Utils.ArrayContains(group, toTry)
	result(t, false, got)
}
