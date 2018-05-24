package appcall

func AppCall(contractBytes []byte) interface{} { return nil }

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
