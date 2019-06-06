package unio

import (
    "encoding/json"
    "errors"
    "github.com/labstack/gommon/log"
    "gopkg.in/mgo.v2/bson"
    "reflect"
    "strconv"
    "time"
)

// Validate if value is string, if is, try to convert to float/32. If error, returns -999
func (u *Util) StringToFloat(value interface{}) (number float32, err error) {
    number = -999
    if s,ok := value.(string); ok {
        price, err := strconv.ParseFloat(s, 32)
        if err == nil { number = float32(price) }
    } else {
        err = errors.New("variable is not string")
    }
    return number, err
}

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

// Convert struct to string key map
func (u *Util) StructToMap(data interface{}) (map[string]interface{}, error) {
    result := make(map[string]interface{})

    b, err := json.Marshal(data)
    if err != nil { return nil, err }

    err = json.Unmarshal(b, &result)
    return result, err
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

// Validate an interface{} var and return an MongoDB ObjectId
func (u *Util) ToObjectId(id interface{}) bson.ObjectId {
    if hexId, ok := id.(bson.ObjectId); ok {
        return hexId
    } else if stringId, ok := id.(string); ok && bson.IsObjectIdHex(stringId) {
        return bson.ObjectIdHex(stringId)
    }
    return ""
}

// ToString Change arg to string
func (u *Util) ToString(arg interface{}) string {
    var tmp = reflect.Indirect(reflect.ValueOf(arg)).Interface()
    switch v := tmp.(type) {
    case int:
        return strconv.Itoa(v)
    case int8:
        return strconv.FormatInt(int64(v), 10)
    case int16:
        return strconv.FormatInt(int64(v), 10)
    case int32:
        return strconv.FormatInt(int64(v), 10)
    case int64:
        return strconv.FormatInt(v, 10)
    case string:
        return v
    case float32:
        return strconv.FormatFloat(float64(v), 'f', -1, 32)
    case float64:
        return strconv.FormatFloat(v, 'f', -1, 64)
    case time.Time:
        return v.Format("2006-01-02 15:04:05")
    case reflect.Value:
        return u.ToString(v.Interface())
    default:
        return ""
    }
}