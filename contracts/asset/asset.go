package asset

import "github.com/ontio/neo-go-compiler/vm/api/native"

var(
	OntContract = "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV".ToScriptHash()
	OngContract = "AFmseVrdL9f9oyCzZefL9tG6UbvhfRZMHJ".ToScriptHash()
)

func Main(operation string, args []interface{}) bool {
	if operation == "RecycleAsset" {
		from := args[0].([]byte)
		to := args[1].([]byte)
		ont := args[2].(uint64)
		ong := args[3].(uint64)
		return RecycleAsset(from, to, ont, ong)
	}
	return false
}

func RecycleAsset(from []byte, to []byte, ont uint64, ong uint64) bool {
    transfer := Transfer{from, to, ont}
    ret := native.Invoke(0, OntContract, "transfer", []interface{}{transfer})
    if ret[0] != 1 {
    	return false
	}
	transfer.Value = ong
	ret = native.Invoke(0, OngContract, "transfer", []interface{}{transfer})
	if ret[0] != 1 {
		return false
	}
	return true
}

type Transfer struct {
	From []byte
	To []byte
	Value uint64
}