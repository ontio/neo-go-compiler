package system

func GetScriptContainer() interface{} { return nil }

func GetExecutingScriptHash() []byte { return nil }

func GetCallingScriptHash() []byte { return nil }

func GetEntryScriptHash() []byte { return nil }

//func SerializeContract(version int, code []byte, contractAddress string, method string, args []byte) []byte {
//	return nil
//}
//
//func SerializeTransfer(from []byte, to []byte, amount int64) []byte { return nil }

func GetContract(hash interface{}) interface{}  { return nil }
func GetTransactionHeight(hash interface{}) int { return 0 }

func GetStorage() interface{} { return nil }

func Base58ToAddress(base58str string) []byte {return nil}
func AddressToBase58(addr []byte) string {return ""}
func VerifyBase58() bool { return true}