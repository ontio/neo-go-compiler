package contracts

import (
	"neo-go-compiler/vm/api/runtime"
	"neo-go-compiler/vm/api/native"
	"neo-go-compiler/vm/api/tools"
)

type transfer struct{
	From []byte
	To []byte
	Amount int64
}


func Main(operation string,args []interface{}) interface{}{

	if operation == "transfer"{
		if len(args) != 3{return false}
		from := args[0].([]byte)
		to := args[1].([]byte)
		amount := args[2].(int64)
		return transONT(from,to,amount)
	}
	return false
}


func transONT(from []byte, to []byte, amount int64) bool {
	if runtime.RuntimeCheckWitness(from) == false{return false}
	contractAddr:=[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	param := transfer{from,to,amount}
	ver := 1
	bs :=native.Invoke([]interface{}{param},"transfer",contractAddr,ver)
	if bs != nil && tools.BytesEquals(bs,[]byte("1")){
		return true
	}else{
		return false
	}
}
