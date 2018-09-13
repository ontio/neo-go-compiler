package contracts

import "github.com/ontio/neo-go-compiler/vm/api/runtime"

func Main(operation string, args []interface{}) bool {

	if operation == "invoke" {
		s := "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"
		runtime.Notify([]interface{}{s})
	}

	if operation == "loop"{


		//for i := 0;i < 10 ;i++ {
		//	runtime.Log("hello")
		//}
		var owner = []byte{0xaf, 0x12, 0xa8, 0x68, 0x7b, 0x14, 0x94, 0x8b, 0xc4, 0xa0, 0x08, 0x12, 0x8a, 0x55, 0x0a, 0x63, 0x69, 0x5b, 0xc1, 0xa5}


		a,b  := owner[0],owner[1]
		if a == b {
			runtime.Notify([]interface{}{"yes"})
		}else{
			runtime.Notify([]interface{}{"no"})
		}


	}

	return true

}