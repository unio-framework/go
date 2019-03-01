package unio

import (
    "reflect"
)

func (u *Util) ArrayContains(ss interface{}, e interface{}) bool {
    s := reflect.ValueOf(ss)
    for i := 0; i < s.Len(); i++ {
        a := s.Index(i).Interface()
        if a == e {
            return true
        }
    }
    return false
}