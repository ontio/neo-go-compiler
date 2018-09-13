package contracts

import (
	"github.com/ontio/neo-go-compiler/vm/api/runtime"
	"github.com/ontio/neo-go-compiler/vm/api/storage"
)

type TToken struct {
	TotalSupply int64
	Name        string
}

func (t TToken) GetTotalSupply() int64 {
	return t.TotalSupply
}

func NewTToken() TToken{
	runtime.Log("in NewTToken")
	return TToken{TotalSupply: 200000000, Name: "MYTOKEN"}
}


func Main(operation string, args []interface{}) bool {

	var (
		token_key = "TOTAL_SUPPLY"
		ctx       = storage.GetContext()
		token     = NewTToken()
	)

	if operation == "init" {
		total := storage.Get(ctx, token_key)
		if total != nil {
			runtime.Notify([]interface{}{"Already initialized"})
			return false
		} else {
			storage.Put(ctx, token_key, token.TotalSupply)
			runtime.Notify([]interface{}{"init succeed!"})
			return true
		}
	}

	if operation == "getTotalSupply" {
		totalSupply := storage.Get(ctx, token_key)
		runtime.Notify([]interface{}{totalSupply})
		return true
	}



	runtime.Notify([]interface{}{"operation not supported"})
	return false
}
