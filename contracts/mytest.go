package contracts

import "neo-go-compiler/vm/api/runtime"

func Main(operation string, args []interface{}) interface{} {

	if operation == "testadd"{
		if len(args) != 2{
			runtime.RuntimeNotify("testadd args ")
			return false
		}
		a := args[0].(int)
		b := args[1].(int)
		c := a + b
		runtime.RuntimeNotify(c)
		return c

	}
	if operation == "testsub"{
		if len(args) != 2{
			runtime.RuntimeNotify("testadd args ")
			return false

		}
		a := args[0].(int)
		b := args[1].(int)
		c := a -b
		runtime.RuntimeNotify(c)
		return c

	}
	runtime.RuntimeNotify("operation not supported! ")
	return false
}


