package api

import "github.com/ontio/neo-go-compiler/vm/api/runtime"

func Main(operation string, args []interface{}) bool {
	if operation == "StructTest" {
		return MapTest()
	}
	return false
}

func StructTest() bool {
	person := Person{}
	person.Name = "bob"
	person.Age = 20
	runtime.Notify([]interface{}{person.Name})
	runtime.Notify([]interface{}{person.Age})

	pbytes := runtime.RuntimeSerialize(person)

	person2 := runtime.RuntimeDeserialize(pbytes).(Person)
	runtime.Notify([]interface{}{person2.Name})
	runtime.Notify([]interface{}{person2.Age})
	return true
}

type Person struct {
	Name string
	Age int
}
