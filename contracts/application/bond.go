package application

import (
	"github.com/ontio/neo-go-compiler/vm/api/runtime"
	"github.com/ontio/neo-go-compiler/vm/api/storage"
	"github.com/ontio/neo-go-compiler/vm/api/tools"
	"github.com/ontio/neo-go-compiler/vm/api/native"
)

var (
	ontAddr = "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV".ToScriptHash()
	admin = "AeS7aUsTmf7egcGQGS88LZAGD8gNFmCJnD".ToScriptHash()
	bondPrefix = []byte("Bond_")
	bondInvestorPrefix = []byte("BondInvestor_")
	bondPaidPrefix = []byte("BondPaid_")
	minInterval = 259200
	minIssueCap = 100000
	minInvestCap = 1000
	minRound = 6
)


func Main(operation string, args []interface{}) bool {
	if operation == "IssueBond" {
		return IssueBond()
	}
	return false
}

func IssueBond(bondName string, parValue int, purchaseEndTime int,
	interval int, round int, couponRate int, totalCap int, Account []byte) bool {
	if runtime.CheckWitness(admin) {
		return false
	}
	if purchaseEndTime < runtime.GetTime() {
		return false
	}
	if totalCap < minIssueCap || round < minRound || couponRate <= 0 || interval < minInterval {
		return false
	}
	if !validateAddress(Account) {
		return false
	}
	bond := BondItem{}
	bond.purchaseEndTime = purchaseEndTime
	bond.CouponRate = couponRate
	bond.Interval = interval
	bond.TotalCap = totalCap
	bond.remainCap = totalCap
	bond.Round = round
	bond.Maturity = purchaseEndTime + round * interval
	b := runtime.RuntimeSerialize(bond)
	storage.Put(storage.GetContext(), tools.Cat(bondPrefix, []byte(bondName)), b)
	return true
}

func InvestBond(bondName string, account []byte, bondNumber int) bool {
    if !runtime.CheckWitness(account) {
    	return false
	}
	if bondNumber <= 0 || !validateAddress(account) || !validateBond(bondName) {
		return false
	}
	bond := runtime.RuntimeDeserialize(GetBond(bondName)).(BondItem)
	if runtime.GetTime() > bond.purchaseEndTime {
		runtime.Notify([]interface{}{"bond subscription has been ended."})
		return false
	}
	investValue := bondNumber * bond.ParValue
	if bond.remainCap < investValue {
		runtime.Notify([]interface{}{"bond remain invest capacity not enough."})
		return false
	}
	ret := native.Invoke(0, ontAddr, "transfer", []interface{}{Transfer{account, bond.Account, investValue}})
    if ret[0] != 1 {
    	return false
	}
	bond.remainCap = bond.TotalCap - investValue
	investorKey := tools.Cat(bondInvestorPrefix, []byte(bondName))
	balance := storage.Get(storage.GetContext(), investorKey).(int)
	storage.Put(storage.GetContext(), investorKey, balance + investValue)
	return true
}

func PayInterstOrPrincipal(bondName string, account []byte) bool {
	if !validateBond(bondName) {
		return false
	}
	investorKey := tools.Cat(bondInvestorPrefix, []byte(bondName))
	investorKey = tools.Cat(investorKey, account)
	balance := storage.Get(storage.GetContext(), investorKey).(int)
	if balance < minInvestCap {
		return false
	}
	bond := runtime.RuntimeDeserialize(GetBond(bondName)).(BondItem)
	paidKey := tools.Cat(bondPaidPrefix, []byte(bondName))
	paidKey = tools.Cat(paidKey, account)
	paidRound := storage.Get(storage.GetContext(), paidKey).(int)
	currentRound := (runtime.GetTime() - bond.purchaseEndTime)/bond.Interval
	if paidRound > bond.Round {
		return false
	}
	if currentRound > bond.Round {
		return false
	}
	investValue := storage.Get(storage.GetContext(), investorKey).(int)
	interest := (currentRound - paidRound) * (investValue * bond.CouponRate)/100
	var ret []byte
	if currentRound == bond.Round {
		ret = native.Invoke(0, ontAddr, "transfer", []interface{}{Transfer{bond.Account,account, uint64(interest + investValue)}})
	} else {
		ret = native.Invoke(0, ontAddr, "transfer", []interface{}{Transfer{bond.Account,account,uint64(interest)}})
	}
	if ret[0] != 1 {
		return false
	}
	storage.Put(storage.GetContext(), paidKey, paidRound + 1)
	return true
}

func GetBond(bondName string) interface{} {
	return storage.Get(storage.GetContext(), tools.Cat(bondPrefix, []byte(bondName)))
}

func validateAddress(address []byte) bool {
	if len(address) != 20 {
		return false
	}
	return true
}
func validateBond(bondName string) bool {
	v := storage.Get(storage.GetContext(), tools.Cat(bondPrefix, []byte(bondName))).([]byte)
	if v == nil || len(v) == 0 {
		return false
	}
    return true
}

type BondItem struct {
	ParValue int
	purchaseEndTime int
	Maturity int
	Interval int
	CouponRate int
	Round int
	TotalCap int
	remainCap int
	Account []byte
}

type Transfer struct {
	From []byte
	To []byte
	Value uint64
}
