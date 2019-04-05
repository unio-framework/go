package unio

//noinspection GoUnusedExportedFunction
func Ternary(rule bool, trueResult interface{}, falseResult interface{}) interface{} {
    if rule {
        return trueResult
    } else {
        return falseResult
    }
}