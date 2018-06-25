package contracts

import (
	"neo-go-compiler/vm/api/runtime"
	"neo-go-compiler/vm/api/storage"
	"neo-go-compiler/vm/api/system"
)

type Domain struct {
	Url   string
	Price int64
	Owner []byte
}

func NewDomain(url string, price int64, owner []byte) Domain {
	return Domain{Url: url, Price: price, Owner: owner}
}

func bytesEquals(a []byte, b []byte) bool {

	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true

}

/*func transONT(from []byte, to []byte, amount int64) bool {

	transferbytes := system.SerializeTransfer(from, to, amount)
	runtime.Log("====transONT 3")
	contractBytes := system.SerializeContract(0,
		[]byte(""),
		"ff00000000000000000000000000000000000001",
		"transfer",
		transferbytes)
	runtime.Notify(contractBytes)
	appcall.AppCall(contractBytes)

	return true
}
*/
func Main(operation string, args []interface{}) bool {

	var (
		ctx      = storage.GetContext()
		selfAddr = system.GetExecutingScriptHash()
	)

	if operation == "register" {
		if len(args) != 2 {
			runtime.Notify("args count error!")
			return false
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		runtime.Notify(addr)
		if runtime.CheckWitness(addr) == false {
			runtime.Notify("Not a valide address!")
			return false
		}
		runtime.Notify(addr)
		if storage.Get(ctx, url) == nil {
			storage.Put(ctx, url, addr)
			runtime.Notify("register succeed!")
			return true
		} else {
			runtime.Notify("already registered!")
			return false
		}
	}

	if operation == "sell" {
		if len(args) != 3 {
			runtime.Notify("args count error!")
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		price := args[2].(int64)
		if runtime.CheckWitness(addr) == false {
			runtime.Notify("Not a valide address!")
			return false
		}
		owner := storage.Get(ctx, url)
		isOwner := bytesEquals(owner.([]byte), addr)

		if isOwner == false {
			runtime.Notify("Not owner! ")
			return false
		} else {
			storage.Put(ctx, url, selfAddr)
			storage.Put(ctx, "Original_Owner_"+url, addr)
			storage.Put(ctx, "Price_"+url, price)
			runtime.Notify("Sell succeed!")
			return true
		}
	}

	if operation == "buy" {
		if len(args) != 3 {
			runtime.Notify("args count error!")
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		price := args[2].(int64)

		if runtime.CheckWitness(addr) == false {
			runtime.Notify("Not a valide address!")
			return false
		}

		owner := storage.Get(ctx, url)
		isOwner := bytesEquals(owner.([]byte), selfAddr)
		if isOwner == false {
			runtime.Notify("url is not in sale ")
			return false
		}

		currentPrice := storage.Get(ctx, "Price_"+url)
		if currentPrice.(int64) >= price {
			runtime.Notify("Price is lower than current price")
			return false
		}

		//tmp
		return true
	}

/*	if operation == "transfer" {
		if len(args) != 3 {
			runtime.Notify("args count error!")
		}
		from := args[0].([]byte)
		to := args[1].([]byte)
		amount := args[2].(int64)
		res := transONT(from, to, amount)
		runtime.Notify("transfer done")
		return res

	}*/

	if operation == "testmap" {
		if len(args) != 2{
			//runtime.Notify("args count error!")
		}
		key := args[0].(string)
		value := args[1].(string)
		//runtime.Log("===testmap0")

		m := make(map[string]string)
		//runtime.Log("===testmap1")
		m[key] = value
		//runtime.Log("===testmap2")
		val2 := m[key]
		//runtime.Log("===testmap3")
		runtime.Notify(val2)

		newkey:="testkey"
		m[newkey]="testvalue"
		val3 := m[newkey]
		runtime.Notify(val3)

		for k,v :=range m {
			runtime.RuntimeNotify(k)
			runtime.RuntimeNotify(v)
		}
		return true
	}

	if operation == "testarray" {
		arr := make([]string,2)
		arr[0] = "aaaa"
		arr[1] = "bbbb"
		runtime.Notify(arr[0])
		runtime.Notify(arr[1])

		for i, a:= range arr{
			runtime.Notify(i)
			runtime.Notify(a)
		}

		return true
	}

	if operation == "teststruct" {
		type Information struct{
			info string
			id string
		}

		type Student struct{
			name string
			age int
			info Information
		}

		a:= Information{"testtest","afad"}
		b:= Student{"jack",10,a}

		runtime.Notify(a.info)
		runtime.Notify(b.info.id)

		return true
	}


	if operation == "teststructarray"{
		type Compsite struct{
			name string
			code int
		}

		array := make([]Compsite,2)

		a := Compsite{name:"aaa",code:1}
		//a := Compsite{"aaa",1}
		b := Compsite{"bbb",2}
		runtime.Log("===teststructarray1")
		array[0] = a
		array[1] = b
		runtime.Log("===teststructarray2")
		c := array[0].name
		runtime.Notify(c)
		runtime.Notify(array[1].code)
		return true

	}



	runtime.Notify("not supported method!")

	return false

}
