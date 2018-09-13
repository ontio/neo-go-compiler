package contracts

import "github.com/ontio/neo-go-compiler/vm/api/runtime"

func Main(operation string, args []interface{}) interface{} {

	if operation == "testTobytes"{
		//b := []byte("hello world")

		a:= "hi there "
		c:= "ontology"

		s := "const hi there " + a  +  "ontology " + c


		runtime.RuntimeNotify([]interface{}{s})
		return s
	}
	//
	//if operation == "testmap" {
	//	if len(args) != 2{
	//		//runtime.RuntimeNotify("args count error!")
	//	}
	//	key := args[0].(string)
	//	value := args[1].(string)
	//	//runtime.Log("===testmap0")
	//
	//	m := make(map[string]string)
	//	//runtime.Log("===testmap1")
	//	m[key] = value
	//	//runtime.Log("===testmap2")
	//	val2 := m[key]
	//	//runtime.Log("===testmap3")
	//	runtime.RuntimeNotify([]interface{}{val2})
	//
	//	newkey:="testkey"
	//	m[newkey]="testvalue"
	//	val3 := m[newkey]
	//	runtime.RuntimeNotify([]interface{}{val3})
	//
	//	for k,v :=range m {
	//		runtime.RuntimeNotify([]interface{}{k})
	//		runtime.RuntimeNotify([]interface{}{v})
	//	}
	//	return true
	//}
	//
	//if operation == "testarray" {
	//	arr := make([]string,2)
	//	arr[0] = "aaaa"
	//	arr[1] = "bbbb"
	//	runtime.RuntimeNotify([]interface{}{arr[0]})
	//	runtime.RuntimeNotify([]interface{}{arr[1]})
	//
	//	for i, a:= range arr{
	//		runtime.RuntimeNotify([]interface{}{i})
	//		runtime.RuntimeNotify([]interface{}{a})
	//	}
	//
	//	return true
	//}
	//
	//if operation == "teststruct" {
	//	type Information struct{
	//		info string
	//		id string
	//	}
	//
	//	type Student struct{
	//		name string
	//		age int
	//		info Information
	//	}
	//
	//	a:= Information{"testtest","afad"}
	//	b:= Student{"jack",10,a}
	//
	//	runtime.RuntimeNotify([]interface{}{a.info})
	//	runtime.RuntimeNotify([]interface{}{b.info.id})
	//
	//	return true
	//}
	//
	//
	//if operation == "teststructarray"{
	//	type Compsite struct{
	//		name string
	//		code int
	//	}
	//
	//	array := make([]Compsite,2)
	//
	//	a := Compsite{name:"aaa",code:1}
	//	//a := Compsite{"aaa",1}
	//	b := Compsite{"bbb",2}
	//	runtime.Log("===teststructarray1")
	//	array[0] = a
	//	array[1] = b
	//	runtime.Log("===teststructarray2")
	//	c := array[0].name
	//	runtime.RuntimeNotify([]interface{}{c})
	//	runtime.RuntimeNotify([]interface{}{array[1].code})
	//	return true
	//
	//}

	runtime.RuntimeNotify([]interface{}{"operation not supported! "})
	return false
}


