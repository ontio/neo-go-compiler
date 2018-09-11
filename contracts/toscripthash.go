package contracts

import (
	"neo-go-compiler/vm/api/tools"
)

var owner = tools.ToScriptHash("Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT")
func Main(operation string, args []interface{}) interface{}{


	if operation == "test"{
		return owner
	}
	return "not support"
}