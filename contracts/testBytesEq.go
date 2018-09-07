package contracts

import (
	"neo-go-compiler/vm/api/runtime"
	"neo-go-compiler/vm/api/tools"
)

func Main(operation string, args []interface{}) bool {
	if operation == "isEq"{
		if len(args) != 2{
			return false
		}
		addr1 := args[0].([]byte)
		addr2 := args[1].([]byte)
		res := tools.BytesEquals(addr1,addr2)
		if res{
			runtime.RuntimeNotify([]interface{}{"address1 eqs address2"})
			return true
		}else{
			runtime.RuntimeNotify([]interface{}{"address1 not eqs address2"})
			return false
		}
	}
	return false
}
