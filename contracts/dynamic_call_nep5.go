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

	if operation == "CallNep5Contract" {
		if len(args) != 4 {
			return false
		}
		from := args[0].([]byte)
		to := args[1].([]byte)
		value := args[2].(uint64)
		hash := args[3].([]byte)
		return CallNep5Contract()
	}
}

func CallNep5Contract() bool {}