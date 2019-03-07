package unio

import "reflect"

/**
AnyOf
Check for first not nil value, use as var1 || var2 shortcut
 */
func (u *Util) AnyOf(values ...interface{}) interface{} {
    for _,value := range values {
        if value != nil {
            return value
        }
    }
    return nil
}

/**
In
Compare if data is one of comparators
 */
func (u *Util) In(data interface{}, compares ...interface{}) bool {
    for _,value := range compares {
        if reflect.DeepEqual(data, value) {
            return true
        }
    }
    return false
}