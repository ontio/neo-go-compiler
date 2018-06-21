
//package mytoken
package contracts

import (
	"neo-go-compiler/vm/api/runtime"
)


func Main(operation string, args []interface{}) bool {

	if operation == "add" {
		if len(args) != 2 {
			return false
		}
		f:= false
		if f ==true{
			return false
		}
		runtime.Notify(args)

		a := args[0].(int)
		runtime.Notify(a)

		b := args[1].(int)
		runtime.Notify(b)

		res := Add(a,b)
		runtime.Notify(res)
		return true
	}

	if operation == "compare" {
		if len(args) != 2 {
			return false
		}
		a := args[0].(int)
		b := args[1].(int)

		res :=  Compare(a,b)
		runtime.Notify(res)
		return true
	}

	return false
}

func Add(a int,b int)int{
	return  a + b
}

func Compare(a,b int)int{
	if (a >= b){
		return a
	}
	return b
}