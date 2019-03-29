// Copyright 2019 Leandro Akira Omiya Takagi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package unio

import (
	"reflect"
)

// Search if a value contains inside array
func (u *Util) ArrayContains(ss interface{}, e interface{}) bool {
	s := reflect.ValueOf(ss)
	for i := 0; i < s.Len(); i++ {
		a := s.Index(i).Interface()
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}