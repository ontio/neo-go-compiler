package api

import "github.com/ontio/neo-go-compiler/vm/api/runtime"

func Main(operation string, args []interface{}) bool {
	if operation == "MapTest" {
		return MapTest()
	}
	return false
}

func MapTest() bool {
	m := make(map[string] int)
	m["k1"] = 100
	m["k2"] = 200
	runtime.Notify([]interface{}{m["k1"]})
	runtime.Notify([]interface{}{m["k2"]})
	mBytes := runtime.RuntimeSerialize(m)

	m2 := runtime.RuntimeDeserialize(mBytes).(map[string]int)
	runtime.Notify([]interface{}{m2["k1"]})
	runtime.Notify([]interface{}{m2["k2"]})
	return true
}
