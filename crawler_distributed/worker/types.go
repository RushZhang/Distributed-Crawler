package worker

//ParseCity的参数就是nil，ProfileParser的参数就是userName。url和contents是系统自动会分配的
//{"ParseCity", nil}, {"ProfileParser", userName}
type SerializedParser struct {
	FucntionName string
	Args interface{}
}