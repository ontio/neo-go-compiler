package contracts

import (
	"neo-go-compiler/vm/api/appcall"
	"neo-go-compiler/vm/api/runtime"
)

func Main(operation string,args []interface{})interface{}{
	if operation == "test"{
		//contractAddress := tools.ToScriptHash("AdoZ1ZRuovzDRAtyCEFkSTbmo9vaXKAaZc")
		runtime.RuntimeLog("before appcall")
		//return appcall.AppCall("APPWgNbWvUdQjQxeN7RduYweH3caaM1LM1","transfer",args)
		return appcall.AppCall("83e69795f9c314a8c4f483e221927f41285a8653","transfer",args)
	}
	return false
}