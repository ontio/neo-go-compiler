package contracts

import (
	"neo-go-compiler/vm/api/runtime"
	"neo-go-compiler/vm/api/storage"
	"neo-go-compiler/vm/api/system"
	"neo-go-compiler/vm/api/tools"
)

const (
	NAME         = "tokenName"
	SYMBOL       = "SYMBOL"
	DECIMAL      = 8
	FACTOR       = 100000000
	TOTAL_AMOUNT = 100000000
)

var (
	owner           = tools.ToScriptHash("Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT")
	transfer_prefix = []byte("transfer_")
	approve_prefix  = []byte("approve_")
	supply_key      = []byte("totalSupply")
	ctx             = storage.GetStorageContext()
	selfAddr        = system.GetExecutingScriptHash()
)

func Main(operation string, args []interface{}) interface{} {

	if operation == "name" {
		return NAME
	}
	if operation == "totalSupply" {
		return TOTAL_AMOUNT * FACTOR
	}
	if operation == "decimal" {
		return DECIMAL
	}
	if operation == "init" {
		return tokenInit()
	}
	if operation == "symbol" {
		return SYMBOL
	}
	if operation == "balanceOf" {
		if len(args) != 1 {
			return false
		}
		acct := args[0].([]byte)
		return balanceOf(acct)
	}
	if operation == "transfer" {
		if len(args) != 3 {
			return false
		}
		from := args[0].([]byte)
		to := args[1].([]byte)
		amount := args[2].(int64)
		return tokenTransfer(from, to, amount)
	}
	if operation == "transferMuti" {
		return tokenTransferMulti(args)
	}
	if operation == "approve" {
		if len(args) != 3 {
			return false
		}
		owner := args[0].([]byte)
		spender := args[1].([]byte)
		amount := args[2].(int64)
		return approve(owner, spender, amount)
	}
	if operation == "allowance" {
		if len(args) != 2 {
			return false
		}
		owner := args[0].([]byte)
		spender := args[1].([]byte)
		return allowance(owner, spender)
	}
	if operation == "transferFrom" {
		if len(args) != 4 {
			return false
		}
		//spender := args[0].([]byte)
		//from_acct := args[1].([]byte)
		//to_acct := args[2].([]byte)
		//amount := args[3].(int64)
		//runtime.RuntimeLog("==============")
		//runtime.RuntimeLog(spender)
		//runtime.RuntimeLog(from_acct)
		//runtime.RuntimeLog(to_acct)
		//runtime.RuntimeLog(amount)
		//runtime.RuntimeLog("==============")
		return transferFrom(args)
	}

	return false
}

func tokenInit() bool {
	if storage.GetStorage(ctx, supply_key) != nil {
		return false
	} else {
		total := TOTAL_AMOUNT * FACTOR
		storage.PutStorage(ctx, supply_key, total)
		storage.PutStorage(ctx, tools.Cat(transfer_prefix, owner), total)

		runtime.RuntimeNotify([]interface{}{"transfer", "", owner, total})
		return true
	}
}

func balanceOf(acct []byte) int64 {
	balance := storage.GetStorage(ctx, tools.Cat(transfer_prefix, acct))
	if balance == nil {
		return 0
	} else {
		return balance.(int64)
	}
}

func tokenTransfer(from []byte, to []byte, amount int64) bool {

	if tools.BytesEquals(from, to) == true {
		return true
	}

	if amount == 0 {
		return true
	}

	if amount < 0 {
		return false
	}

	if runtime.RuntimeCheckWitness(from) == false {
		return false
	}

	fromkey := tools.Cat(transfer_prefix, from)
	//runtime.RuntimeLog("fromkey")
	//runtime.RuntimeLog(fromkey)

	fromBalance := balanceOf(from)
	if fromBalance < amount {
		return false
	}
	runtime.RuntimeLog("fromBalance")

	runtime.RuntimeLog(fromBalance)
	if fromBalance == amount {
		storage.DeleteSorage(ctx, fromkey)
	} else {
		storage.PutStorage(ctx, fromkey, fromBalance-amount)
	}
	runtime.RuntimeLog("tokenTransfer 7")

	tokey := tools.Cat(transfer_prefix, to)
	tobalance := balanceOf(to)
	storage.PutStorage(ctx, tokey, tobalance+amount)
	runtime.RuntimeLog("tokenTransfer 8")

	runtime.RuntimeNotify([]interface{}{"transfer", from, to, amount})

	return true
}

func tokenTransferMulti(args []interface{}) bool {

	for i := 0; i < len(args); i++ {
		elem := args[i].([]interface{})
		if len(elem) != 3 {
			return false
		}
		from := elem[0].([]byte)
		to := elem[1].([]byte)
		amount := elem[2].(int64)

		if tokenTransfer(from, to, amount) == false {
			return false
		}
	}

	return true
}

func approve(owner []byte, spender []byte, amount int64) bool {
	if amount <= 0 {
		return false
	}

	if runtime.RuntimeCheckWitness(owner) == false {
		return false
	}

	approveKey := tools.Cat(tools.Cat(approve_prefix, owner), spender)
	allowance := storage.GetStorage(ctx, approveKey)
	if allowance != nil {
		storage.PutStorage(ctx, approveKey, allowance.(int64)+amount)
	} else {
		storage.PutStorage(ctx, approveKey, amount)
	}
	runtime.RuntimeNotify([]interface{}{"approve", owner, spender, amount})

	return true
}

func transferFrom(args []interface{}) bool {
	spender := args[0].([]byte)
	from_acct := args[1].([]byte)
	to_acct := args[2].([]byte)
	amount := args[3].(int64)
	runtime.RuntimeLog("======transferFrom========")
	runtime.RuntimeLog(spender)
	runtime.RuntimeLog(from_acct)
	runtime.RuntimeLog(to_acct)
	runtime.RuntimeLog(amount)
	runtime.RuntimeLog("======transferFrom========")

	runtime.RuntimeLog("transferFrom 1")
	if amount <= 0 {
		return false
	}
	runtime.RuntimeLog("transferFrom 2")

	if runtime.RuntimeCheckWitness(spender) == false {
		return false
	}
	runtime.RuntimeLog("transferFrom 3")

	approveKey := tools.Cat(tools.Cat(approve_prefix, from_acct), spender)
	allowanceNo := allowance(from_acct, spender)
	if allowanceNo < amount {
		return false
	}
	runtime.RuntimeLog("transferFrom 4")

	if allowanceNo == amount {
		storage.DeleteSorage(ctx, approveKey)
	}
	runtime.RuntimeLog("transferFrom 5")

	storage.PutStorage(ctx, approveKey, allowanceNo-amount)

	if len(to_acct) != 20 {
		return false
	}
	runtime.RuntimeLog("transferFrom 6")

	fromKey := tools.Cat(transfer_prefix, from_acct)
	fromBalance := balanceOf(from_acct)
	if  fromBalance < amount {
		return false
	}
	runtime.RuntimeLog("transferFrom 7")

	storage.PutStorage(ctx, fromKey, fromBalance-amount)
	runtime.RuntimeLog("transferFrom 8")

	toKey := tools.Cat(transfer_prefix, to_acct)
	toBalance := balanceOf(to_acct)

	storage.PutStorage(ctx, toKey, toBalance+amount)
	runtime.RuntimeLog("transferFrom 9")

	runtime.RuntimeNotify([]interface{}{"transfer", from_acct, to_acct, amount})

	return true
}

func allowance(owner []byte, spender []byte) int64 {
	approveKey := tools.Cat(tools.Cat(approve_prefix, owner), spender)
	allowance := storage.GetStorage(ctx, approveKey)
	if allowance == nil {
		return 0
	} else {
		return allowance.(int64)
	}
}
