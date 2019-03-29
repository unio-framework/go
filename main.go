package unio

type(
    Config     struct {}
    Handler    struct {}
    Middleware struct {}
    Search     struct {}
    Util       struct {}
)

//noinspection GoUnusedGlobalVariable
var Configs = Config{}
//noinspection GoUnusedGlobalVariable
var Handlers = Handler{}
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