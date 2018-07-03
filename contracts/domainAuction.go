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
	runtime.RuntimeLog("==bytesEquals===")
	if a == nil && b == nil {
		return true
	}
	runtime.RuntimeLog("==bytesEquals 1===")

	if len(a) != len(b) {
		return false
	}
	runtime.RuntimeLog("==bytesEquals 2===")

	for i := 0; i < len(a); i++ {
		runtime.RuntimeLog("===i")
		if a[i] != b[i] {
			runtime.RuntimeLog("return false")
			return false
		}
	}
	runtime.RuntimeLog("==bytesEquals 3===")

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
	runtime.RuntimeNotify(contractBytes)
	appcall.AppCall(contractBytes)

	return true
}
*/
func Main(operation string, args []interface{}) bool {

	var (
		ctx      = storage.GetStorageContext()
		selfAddr = system.GetExecutingScriptHash()
	)


	if operation == "query"{
		if len(args) != 1 {
			runtime.RuntimeNotify("query args count error!")
			return false
		}
		url := args[0].(string)
		owner := storage.GetStorage(ctx,url)
		runtime.RuntimeNotify(owner)
		return true
	}

	if operation == "register" {
		if len(args) != 2 {
			runtime.RuntimeNotify("register args count error!")
			return false
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		if runtime.RuntimeCheckWitness(addr) == false {
			runtime.RuntimeNotify("Not a valide address!")
			return false
		}
		if storage.GetStorage(ctx, url) == nil {
			storage.PutStorage(ctx, url, addr)
			runtime.RuntimeNotify("register succeed!")
			return true
		} else {
			runtime.RuntimeNotify("already registered!")
			return false
		}
	}



	if operation == "sell" {
		if len(args) != 3 {
			runtime.RuntimeNotify("sell args count error!")
			return false
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		price := args[2].(int64)
		runtime.RuntimeLog("00")

		if runtime.RuntimeCheckWitness(addr) == false {
			runtime.RuntimeNotify("Not a valide address!")
			return false
		}
		runtime.RuntimeLog("111111")
		owner := storage.GetStorage(ctx, url)
		runtime.RuntimeNotify(owner)

		runtime.RuntimeLog("11222")
		isOwner := bytesEquals(owner.([]byte), addr)
		runtime.RuntimeLog("2222")

		if isOwner == false {
			runtime.RuntimeLog("3333")

			runtime.RuntimeNotify("Not owner! ")
			return false
		} else {
			runtime.RuntimeLog("4444")

			storage.PutStorage(ctx, url, selfAddr)
			storage.PutStorage(ctx, "Original_Owner_"+url, addr)
			storage.PutStorage(ctx, "Price_"+url, price)
			runtime.RuntimeNotify("Sell succeed!")
			return true
		}
	}

	if operation == "buy" {
		if len(args) != 3 {
			runtime.RuntimeNotify("buy args count error!")
			return false
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		price := args[2].(int64)

		if runtime.RuntimeCheckWitness(addr) == false {
			runtime.RuntimeNotify("Not a valide address!")
			return false
		}

		owner := storage.Get(ctx, url)
		isOwner := bytesEquals(owner.([]byte), selfAddr)
		if isOwner == false {
			runtime.RuntimeNotify("url is not in sale ")
			return false
		}

		currentPrice := storage.GetStorage(ctx, "Price_"+url)
		if currentPrice.(int64) >= price {
			runtime.RuntimeNotify("Price is lower than current price")
			return false
		}

		//tmp
		return true
	}

/*	if operation == "transfer" {
		if len(args) != 3 {
			runtime.RuntimeNotify("args count error!")
		}
		from := args[0].([]byte)
		to := args[1].([]byte)
		amount := args[2].(int64)
		res := transONT(from, to, amount)
		runtime.RuntimeNotify("transfer done")
		return res

	}*/

	if operation == "testmap" {
		if len(args) != 2{
			//runtime.RuntimeNotify("args count error!")
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
		runtime.RuntimeNotify(val2)

		newkey:="testkey"
		m[newkey]="testvalue"
		val3 := m[newkey]
		runtime.RuntimeNotify(val3)

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
		runtime.RuntimeNotify(arr[0])
		runtime.RuntimeNotify(arr[1])

		for i, a:= range arr{
			runtime.RuntimeNotify(i)
			runtime.RuntimeNotify(a)
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

		runtime.RuntimeNotify(a.info)
		runtime.RuntimeNotify(b.info.id)

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
		runtime.RuntimeNotify(c)
		runtime.RuntimeNotify(array[1].code)
		return true

	}



	runtime.RuntimeNotify("not supported method!")

	return false

}
