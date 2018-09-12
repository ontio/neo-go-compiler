package appcall

func AppCall(contractaddress string, method string, args []interface{}) interface{} { return nil }

//func AppCall(args []interface{},method string,contractaddress []byte ) interface{} { return nil }

type Contract struct {
	Version byte
	Code    []byte
	Address string
	Method  string
	Args    []byte
}

type OntTransfer struct {
	States State
}

type State struct {
	From   []byte
	To     []byte
	Amount int64
}
