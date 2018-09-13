package contracts

import (
	"github.com/ontio/neo-go-compiler/vm/api/runtime"
)

var(
	from = []byte{1, 200, 0, 22, 183, 87, 117, 144, 137, 104, 11, 241, 62, 118, 76, 124, 240, 132, 24, 149}
	to =[]byte{ 1, 200, 0, 22, 183, 87, 117, 144, 137, 104, 11, 241, 62, 118, 76, 124, 240, 132, 24, 149}
	value = 100
)

func Main(operation string, args []interface{}) interface{} {
	if operation == "test"{
		//address := "ff00000000000000000000000000000000000001"
				//address := []byte{255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1 }

		//method := "transfer"
		//param := make([]interface{} ,3)
		//param[0] = from
		//param[1] = to
		//param[2] = value
		param := []interface{}{from,to,value}
		//res := native.Invoke(0,address,method,param)
		runtime.Notify(param)
		return param
	}

	return false

}