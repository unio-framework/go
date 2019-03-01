package unio

type(
    Config     struct {}
    Middleware struct {}
    Util       struct {}
)

//noinspection GoUnusedGlobalVariable
var Configs = Config{}
//noinspection GoUnusedGlobalVariable
var Middlewares = Middleware{}
//noinspection GoUnusedGlobalVariable
var Utils = Util{}

/**
Rule structure
*/
type RequestFormatRule func(method string, key string, value interface{}) interface{}