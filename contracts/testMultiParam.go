package contracts

import "github.com/ontio/neo-go-compiler/vm/api/runtime"

func Main(operation string, args []interface{}) interface{} {
	if operation == "test" {
		if len(args) != 4 {
			return false
		}
		a := args[0].(string)
		b := args[1].(string)
		c := args[2].(string)
		d := args[3].(string)

		return test(a, b, c, d)

	}
	return false
}

func test(a string, b string, c string, d string) string {
	res := a + b + c + d
	runtime.RuntimeNotify([]interface{}{res})
	return res
}
