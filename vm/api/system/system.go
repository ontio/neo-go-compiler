package system

func GetScriptContainer()interface{}{return nil}

func GetExecutingScriptHash() []byte{return nil}

func GetCallingScriptHash() []byte{return nil}

func GetEntryScriptHash() []byte{return nil}

func SerializeContract(version int,code []byte,contractAddress string,method string,args []byte) []byte{return nil}

func SerializeTransfer(from []byte,to []byte,amount int64) []byte{return nil}