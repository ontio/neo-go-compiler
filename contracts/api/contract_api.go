package api

import (
	"github.com/ontio/neo-go-compiler/vm/api/contract"
	"github.com/ontio/neo-go-compiler/vm/api/runtime"
)

func Main(operation string, args []interface{}) bool {
	if operation == "Destroy"{
		return Destroy()
	}
	return true
}

func Destroy() bool {
	res := contract.DestroyContract()
	runtime.Notify([]interface{}{res})
	return true
}
