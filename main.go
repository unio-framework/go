package unio

type(
    Config     struct {}
    Middleware struct {}
    Search     struct {}
    Util       struct {}
)

//noinspection GoUnusedGlobalVariable
var Configs = Config{}
//noinspection GoUnusedGlobalVariable
var Middlewares = Middleware{}
//noinspection GoUnusedGlobalVariable
var Utils = Util{}
//noinspection GoUnusedGlobalVariable
var Searchs = Search{}

/**
Rule structure
*/
type RequestFormatRule func(method string, key string, value interface{}) interface{}

/**
Populate structure
*/
type SearchPopulate func(field string, value interface{}) (interface{}, error)