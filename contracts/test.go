
//package mytoken
package contracts

import (
	"github.com/ontio/neo-go-compiler/vm/api/runtime"
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
		runtime.Notify([]interface{}{a})

		b := args[1].(int)
		runtime.Notify([]interface{}{b})

		res := Add(a,b)
		runtime.Notify([]interface{}{res})
		return true
	}

	if operation == "compare" {
		if len(args) != 2 {
			return false
		}
		a := args[0].(int)
		b := args[1].(int)

		res :=  Compare(a,b)
		runtime.Notify([]interface{}{res})
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