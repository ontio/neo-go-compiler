package api

import (
	"github.com/ontio/neo-go-compiler/vm/api/runtime"
	"github.com/ontio/neo-go-compiler/vm/api/storage"
)

func Main(operation string, args []interface{}) bool {

	if operation == "transferMulti" {
		return transferMulti(args)
	}
	return false
}

func transferMulti(args []interface{}) bool {
    for i:=0; i< len(args); i++ {
    	state := args[i].(State)
    	if !Transfer(state.From, state.To, state.Amount){
    		return false
		}
	}
	return true
}

func Transfer(from []byte, to []byte, value uint64) bool {
	if value < 0 {
		return false
	}
	if runtime.CheckWitness(from) {
		return false
	}
	if len(to) != 20 {
		return false
	}
	from_value := storage.Get(storage.GetContext(), from).(uint64)
	if from_value < value {
		return false
	}
	if from_value == value {
		storage.Delete(storage.GetContext(), from)
	}else {
		storage.Put(storage.GetContext(), from, from_value - value)
	}
	to_value := storage.Get(storage.GetContext(), to).(uint64)
	storage.Put(storage.GetContext(), to, to_value + value)
	runtime.Notify([]interface{}{from, to, value})
	return true
}

type State struct {
	From []byte
	To []byte
	Amount uint64
}