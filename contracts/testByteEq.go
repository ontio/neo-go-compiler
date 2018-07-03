package contracts

import "neo-go-compiler/vm/api/runtime"

func Main(operation string ,args []interface{}) bool{

	if operation == "test"{
		if len(args) != 2{
			runtime.RuntimeNotify("arg count error")
			return false
		}
		a := args[0].([]byte)
		b := args[1].([]byte)
		ret := bytesEquals(a,b)
		runtime.RuntimeNotify(ret)
		return ret

	}
	return false

}


func bytesEquals(a []byte, b []byte) bool {
	if a == nil && b == nil {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		runtime.RuntimeLog("===i")
		if a[i] != b[i] {
			return false
		}
	}

	return true

}