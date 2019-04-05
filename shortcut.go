package unio

//noinspection GoUnusedExportedFunction
func DeleteMapKeys(m map[string]interface{}, keys... string) {
    for _, key := range keys {
        delete(m, key)
    }
}

//noinspection GoUnusedExportedFunction
func Ternary(rule bool, trueResult interface{}, falseResult interface{}) interface{} {
    if rule {
        return trueResult
    } else {
        return falseResult
    }
}