// Copyright 2019 Leandro Akira Omiya Takagi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package unio

import (
	"encoding/json"
	"reflect"
)

// JSON types
type JSON interface{}
type JSONObject map[string]JSON
type JSONArray []JSON

// Check if a interface has JSON pattern
func (u *Util) IsJSON(s interface{}) bool {
	js, _ := u.JSONParse(s)
	return js != nil
}

// Parse any to JSON structure
func (u *Util) JSONParse(s interface{}) (JSON, error) {
	var js JSON

	var b []byte
	var err error

	switch reflect.ValueOf(s).Kind() {
	case reflect.String:
		b = []byte(s.(string))
	case reflect.Slice, reflect.Array, reflect.Map:
		b, err = json.Marshal(s)
		if err != nil {
			return nil, err
		}
	}

	err = json.Unmarshal(b, &js)
	if err != nil {
		return nil, err
	}
	return js, nil
}
