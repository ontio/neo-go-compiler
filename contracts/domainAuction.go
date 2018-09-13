package contracts

import (
	"github.com/ontio/neo-go-compiler/vm/api/runtime"
	"github.com/ontio/neo-go-compiler/vm/api/storage"
	"github.com/ontio/neo-go-compiler/vm/api/system"
	"github.com/ontio/neo-go-compiler/vm/api/native"
	"github.com/ontio/neo-go-compiler/vm/api/tools"
)

type transfer struct{
	From []byte
	To []byte
	Amount int64
}

func transONT(from []byte, to []byte, amount int64) bool {
	if runtime.RuntimeCheckWitness(from) == false{return false}
	contractAddr:=[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	param := transfer{from,to,amount}
	ver := 1
	bs :=native.Invoke([]interface{}{param},"transfer",contractAddr,ver)
	if bs != nil && tools.BytesEquals(bs,[]byte("1")){
		return true
	}else{
		return false
	}
}

func Main(operation string, args []interface{}) bool {

	var (
		ctx      = storage.GetStorageContext()
		selfAddr = system.GetExecutingScriptHash()
	)


	if operation == "query"{
		if len(args) != 1 {
			runtime.RuntimeNotify([]interface{}{"query args count error!"})
			return false
		}
		url := args[0].(string)
		owner := storage.GetStorage(ctx,url)
		runtime.RuntimeNotify([]interface{}{owner})
		return true
	}

	if operation == "register" {
		if len(args) != 2 {
			runtime.RuntimeNotify([]interface{}{"register args count error!"})
			return false
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		if runtime.RuntimeCheckWitness(addr) == false {
			runtime.RuntimeNotify([]interface{}{"Not a valide address!"})
			return false
		}
		if storage.GetStorage(ctx, url) == nil {
			storage.PutStorage(ctx, url, addr)
			runtime.RuntimeNotify([]interface{}{"register succeed!"})
			return true
		} else {
			runtime.RuntimeNotify([]interface{}{"already registered!"})
			return false
		}
	}



	if operation == "sell" {
		if len(args) != 3 {
			runtime.RuntimeNotify([]interface{}{"sell args count error!"})
			return false
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		price := args[2].(int64)

		if runtime.RuntimeCheckWitness(addr) == false {
			runtime.RuntimeNotify([]interface{}{"Not a valide address!"})
			return false
		}
		owner := storage.GetStorage(ctx, url)

		isOwner := tools.BytesEquals(owner.([]byte), addr)

		if isOwner == false {
			runtime.RuntimeNotify([]interface{}{"Not owner! "})
			return false
		} else {
			storage.PutStorage(ctx, url, selfAddr)
			storage.PutStorage(ctx, "Original_Owner_"+url, addr)
			storage.PutStorage(ctx, "Price_"+url, price)
			runtime.RuntimeNotify([]interface{}{"Sell succeed!"})
			return true
		}
	}

	if operation == "buy" {
		if len(args) != 3 {
			runtime.RuntimeNotify([]interface{}{"buy args count error!"})
			return false
		}
		url := args[0].(string)
		addr := args[1].([]byte)
		price := args[2].(int64)

		if runtime.RuntimeCheckWitness(addr) == false {
			runtime.RuntimeNotify([]interface{}{"Not a valide address!"})
			return false
		}
		owner := storage.GetStorage(ctx, url)
		isOwner := tools.BytesEquals(owner.([]byte), selfAddr)
		if isOwner == false {
			runtime.RuntimeNotify([]interface{}{"url is not in sale "})
			return false
		}
		currentPrice := storage.GetStorage(ctx, "Price_"+url)
		if currentPrice.(int64) >= price {
			runtime.RuntimeNotify([]interface{}{"Price is lower than current price"})
			return false
		}
		prevBuyer := storage.GetStorage(ctx,"TP_"+url)
		if prevBuyer != nil{
			if transONT(selfAddr,prevBuyer.([]byte),currentPrice.(int64) ) == false{
				runtime.RuntimeNotify([]interface{}{"refund to prebuyer failed"})
				return false
			}
		}

		if runtime.RuntimeCheckWitness(addr) == false{
			runtime.RuntimeNotify([]interface{}{"CheckWitness  failed"})
			return false
		}

		if transONT(addr,selfAddr,price) == true{
			storage.PutStorage(ctx, "Price_"+url,price)
			storage.PutStorage(ctx, "TP_"+url,addr)
			return true
		}else{
			runtime.RuntimeNotify([]interface{}{"transfer ont failed! "})
			return false
		}
		runtime.RuntimeNotify([]interface{}{"buy succeed! "})
		return true
	}

	if operation == "queryTopPrice"{
		if len(args) != 1{
			runtime.RuntimeNotify([]interface{}{"queryTopPrice args count error!"})
			return false
		}
		url:= args[0].(string)

		tp := storage.GetStorage(ctx,"Price_"+url)
		if tp == nil{
			runtime.RuntimeNotify([]interface{}{0,[]byte{}})
		}else{
			au := storage.GetStorage(ctx,"TP_" + url)
			runtime.RuntimeNotify([]interface{}{tp,au})
		}
		return true

	}

	if operation == "deal"{
		if len(args) != 2{
			runtime.RuntimeNotify([]interface{}{"deal args count error!"})
			return false
		}

		url := args[0].(string)
		addr := args[1].([]byte)

		originOwner := storage.GetStorage(ctx,"Original_Owner_"+url)
		if originOwner == nil{
			runtime.RuntimeNotify([]interface{}{"deal get Original_Owner failed!"})
			return false
		}
		if tools.BytesEquals(originOwner.([]byte),addr) == false{
			runtime.RuntimeNotify([]interface{}{"deal not the origin owner!"})
			return false
		}

		price := storage.GetStorage(ctx,"Price_"+url)
		if price == nil{
			runtime.RuntimeNotify([]interface{}{"deal get price failed!"})
			return false
		}
		buyer := storage.GetStorage(ctx,"TP_"+url)
		if buyer == nil{
			runtime.RuntimeNotify([]interface{}{"deal get buyer failed!"})
			return false
		}

		if transONT(selfAddr,buyer.([]byte),price.(int64)) == false{
			runtime.RuntimeNotify([]interface{}{"deal transfer ont  failed!"})
			return false
		}
		storage.DeleteSorage(ctx,"Price_url")
		storage.DeleteSorage(ctx,"Original_Owner_"+url)
		storage.DeleteSorage(ctx,"TP_"+url)

		storage.PutStorage(ctx,url,buyer)

		return true
	}





	runtime.RuntimeNotify([]interface{}{"not supported method!"})

	return false

}
