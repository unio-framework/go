package unio

import (
    "errors"
    "github.com/labstack/gommon/log"
    "strconv"
    "time"
)

// Validate if value is string, if is, try to convert to int. If error, returns -999
func (u *Util) StringToInt(value interface{}) (number int, err error) {
    number = -999
    if s,ok := value.(string); ok {
        res, err := strconv.Atoi(s)
        if err == nil { number = res }
    } else {
        err = errors.New("variable is not string")
    }
    return number, err
}

// Validate if value is string, if is, try to convert to float/32. If error, returns -999
func (u *Util) StringToFloat(value interface{}) (number float32, err error) {
    number = -999
    if s,ok := value.(string); ok {
        price, err := strconv.ParseFloat(s, 32)
        if err == nil { value = float32(price) }
    } else {
        err = errors.New("variable is not string")
    }
    return number, err
}

// Validate if value is string and try to convert one time format to another format
func (u *Util) TimeConvert(value interface{}, from string, to string) (result string, err error) {
    t, err := time.Parse(from, value.(string))
    if err == nil {
        return t.Format(to), nil
    } else {
        log.Error(err)
    }
    return "", err
}