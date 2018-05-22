package contracts

import (
	"neo-go-compiler/vm/api/runtime"
	"neo-go-compiler/vm/api/storage"
	"neo-go-compiler/vm/api/system"
)

type Domain struct{
	Url string
	Price int64
	Owner []byte
}

func NewDomain(url string,price int64,owner []byte) Domain{
	return Domain{Url:url,Price:price,Owner:owner}
}

func bytesEquals(a []byte,b []byte) bool{

	if a == nil && b == nil{
		return true
	}
	if len(a) != len(b){
		return false
	}
	runtime.Log("len(a")
	runtime.Notify(len(a))
	for i:= 0 ; i < len(a);i++{
		runtime.Notify(i)
		runtime.Log("===bytesEquals 6")
		if a[i] != b[i]{
			runtime.Log("===bytesEquals 3")
			return false
		}
	}
	runtime.Log("===bytesEquals 5")

	return true

}


func Main(operation string,args []interface{})bool{

	var ctx = storage.GetContext()

	if operation == "register"{
		if len(args)!= 2{
			runtime.Notify("args count error!")
			return false
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		runtime.Log("===addr1===")
		runtime.Notify(addr)
		if runtime.CheckWitness(addr) == false{
			runtime.Notify("Not a valide address!")
			return false
		}
		runtime.Log("===addr2===")
		runtime.Notify(addr)
		if storage.Get(ctx,url) == nil{
			storage.Put(ctx,url,addr)
			runtime.Notify("register succeed!")
			return true
		}else{
			runtime.Notify("already registered!")
			return false
		}
	}

	if operation == "sell"{
		if len(args) != 3{
			runtime.Notify("args count error!")
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		price := args[2].(int64)
		selfAddr := system.GetExecutingScriptHash()
		if runtime.CheckWitness(addr) == false{
			runtime.Notify("Not a valide address!")
			return false
		}
		runtime.Log("========1")
		owner := storage.Get(ctx,url)
		runtime.Notify(owner)
		runtime.Log("========2")
		isOwner := bytesEquals(owner.([]byte), addr)
		runtime.Log("========isOwner")
		runtime.Notify(isOwner)

		if isOwner == false {
			runtime.Log("========3")
			runtime.Notify("Not owner ")
			return false
		}else{
			runtime.Log("========4")
			storage.Put(ctx,url,selfAddr)
			storage.Put(ctx,"Original_Owner_"+url,addr)
			storage.Put(ctx,"Price_"+url,price)
			runtime.Notify("Sell succeed!")
			runtime.Log("========5")

			return true
		}

	}


	runtime.Notify(operation + "no supported!")

	return false

}