package tests

import (
	"reflect"
	"testing"
)

func result(t *testing.T, want interface{}, got interface{}) {
	if reflect.DeepEqual(want, got) == false {
		t.Helper()
		t.Errorf("want = %t, got = %t", want, got)
	}
}
