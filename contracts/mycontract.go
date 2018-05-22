
//package mytoken
package contracts

import (
	"neo-go-compiler/vm/api/storage"
	"neo-go-compiler/vm/api/runtime"
)

var owner = []byte{0xaf, 0x12, 0xa8, 0x68, 0x7b, 0x14, 0x94, 0x8b, 0xc4, 0xa0, 0x08, 0x12, 0x8a, 0x55, 0x0a, 0x63, 0x69, 0x5b, 0xc1, 0xa5}

type Token struct {
	Name        string
	Symbol      string
	TotalSupply int
	Owner       []byte
}

func (t Token) AddToCirculation(amount int) bool {
	ctx := storage.GetContext()
	inCirc := storage.Get(ctx, "in_circ").(int)
	inCirc += amount
	storage.Put(ctx, "in_circ", inCirc)
	return true
}

func newToken() Token {
	return Token{
		Name:        "your awesome NEO token",
		Symbol:      "YANT",
		TotalSupply: 1000,
		Owner:       owner,
	}
}

func Main(operation string, args []interface{}) bool {
	token := newToken()
	trigger := runtime.GetTrigger()

	if trigger == runtime.Verification() {
		isOwner := runtime.CheckWitness(token.Owner)
		if isOwner {
			return true
		}
		return false
	}

	if trigger == runtime.Application() {
		if operation == "mintTokens" {
			token.AddToCirculation(100)
		}
	}

	return true
}
