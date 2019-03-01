package unio

import (
    "encoding/json"
    "github.com/labstack/gommon/log"
    "reflect"
)

type JSON       interface{}
type JSONObject map[string]JSON
type JSONArray  []JSON

func (u *Util) IsJSON(s interface{}) bool {
    return u.JSONParse(s) != nil
}

func (u *Util) JSONParse(s interface{}) JSON {
    var js JSON

    var b []byte
    var err error

    switch reflect.ValueOf(s).Kind() {
    case reflect.String:
        b = []byte(s.(string))
    case reflect.Slice, reflect.Array, reflect.Map:
        b,err = json.Marshal(s); if err != nil {
            log.Error(err)
            return nil
        }
    }

    err = json.Unmarshal(b, &js); if err != nil {
        log.Error(err)
        return nil
    }
    return js
}