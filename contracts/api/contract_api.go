package api

import (
	"github.com/ontio/neo-go-compiler/vm/api/contract"
)

func Main(operation string, args []interface{}) bool {
	if operation == "Destroy"{
		return DestroyContract()
	}
	if operation == "Migrate" {
		code := args[0].([]byte)
		return Migrate(code)
	}

}

func DestroyContract() bool {
	//TODO
	//contract.
	return false
}

func Migrate(code []byte) bool {
	contract.MigrateContract(code)
	return true
}
