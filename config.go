package unio

/**
Start all configuratio
*/
func (c Configs) Init() {
    c.LoadEnv()
}