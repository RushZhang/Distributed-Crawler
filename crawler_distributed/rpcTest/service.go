package rpcTest

import "github.com/pkg/errors"

//Service.Method
type DemoService struct {

}

type Args struct {
	A, B int
}

//rpc方法对参数也有要求，一定要是两个参数args、result, 而且result一定要是有指针，返回err
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A)/float64(args.B)
	return nil
}