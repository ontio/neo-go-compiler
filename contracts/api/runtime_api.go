package api

import "github.com/ontio/neo-go-compiler/vm/api/runtime"

func Main(operation string, args []interface{}) bool {
	if operation == "checkAuth" {
		user := args[0].([]byte)
		checkAuth(user)
		return true
	}
	return false

}

func checkAuth(user []byte) bool {
	runtime.Notify([]interface{}{runtime.GetTime()})
	runtime.CheckWitness(user)
	runtime.Notify([]interface{}{"Hi Blockchain"})
	runtime.Notify([]interface{}{"s1", "s2", 1, 2, 3})
	return true
}
