package api

import (
	"github.com/ontio/neo-go-compiler/vm/api/system"
	"github.com/ontio/neo-go-compiler/vm/api/storage"
)

func Main(operation string, args []interface{}) bool {


	if operation == "CallNep5Contract" {
		if len(args) != 4 {
			return false
		}
		from := args[0].([]byte)
		to := args[1].([]byte)
		value := args[2].(uint64)
		hash := args[3].([]byte)
		return CallNep5Contract(from, to, value, hash)
	}
	return false
}

func CallNep5Contract(from []byte, to []byte, value uint64, contractHash []byte) bool {
	if !TransferNEP5(from,to,value,contractHash){
		//抛异常
		//TODO
		return false
	}
	return true
}

func TransferNEP5(from []byte, to []byte, value uint64, contractHash []byte) bool {
	arg := []interface{}{from, to, value}
	contract := contractHash.ToDelegate()
	if !contract("transfer", arg) {
		return false
	}
	return true
}