package contracts

import (
	"github.com/ontio/neo-go-compiler/vm/api/system"
	"github.com/ontio/neo-go-compiler/vm/api/storage"
)

func Main(operation string, args []interface{}) bool {
	var (
		ctx      = storage.GetStorageContext()
		selfAddr = system.GetExecutingScriptHash()
	)
	if operation == "Destroy"{
		return DestroyContract()
	}

}

func DestroyContract() bool {
	contract

}
