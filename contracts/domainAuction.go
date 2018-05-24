package contracts

import (
	"neo-go-compiler/vm/api/appcall"
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

func transONT(from []byte, to []byte, amount int64) bool {
	runtime.Log("====transONT 1")

	transferbytes := system.SerializeTransfer(from, to, amount)
	runtime.Log("====transONT 3")
	contractBytes := system.SerializeContract(0,
		[]byte(""),
		"ff00000000000000000000000000000000000001",
		"transfer",
		transferbytes)
	runtime.Log("====transONT 7")
	runtime.Notify(contractBytes)
	appcall.AppCall(contractBytes)
	runtime.Log("====transONT 8")

	return true
}

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

	if operation == "transfer" {
		if len(args) != 3 {
			runtime.Notify("args count error!")
		}
		from := args[0].([]byte)
		to := args[1].([]byte)
		amount := args[2].(int64)
		res := transONT(from, to, amount)
		runtime.Notify("transfer done")
		return res

	}

	runtime.Notify(operation + "no supported!")

	return false

}
