package vm

// Syscalls are a mapping between the syscall function name
// and the registerd VM interop API.
var Syscalls = map[string]string{
	//System API
	"GetScriptContainer":     "System.ExecutionEngine.GetScriptContainer",
	"GetExecutingScriptHash": "System.ExecutionEngine.GetExecutingScriptHash",
	"GetCallingScriptHash":   "System.ExecutionEngine.GetCallingScriptHash",
	"GetEntryScriptHash":     "System.ExecutionEngine.GetEntryScriptHash",
	//"SerializeContract":      "System.ExecutionEngine.SerializeContract", //temp for test
	//"SerializeTransfer":      "System.ExecutionEngine.SerializeTransfer", //temp for test
	"GetBlockTransactionCount":  "System.Block.GetTransactionCount",
	"GetBlockTransactions":      "System.Block.GetTransactions",
	"GetBlockTransaction":       "System.Block.GetTransaction",
	"GetBlockchainHeight":       "System.Blockchain.GetHeight",
	"GetBlockchainHeader":       "System.Blockchain.GetHeader",
	"GetBlockchainBlock":        "System.Blockchain.GetBlock",
	"GetBlockchainTransaction":  "System.Blockchain.GetTransaction",
	"GetContract":               "System.Blockchain.GetContract",
	"GetTransactionHeight":      "System.Blockchain.GetTransactionHeight",
	"GetHeaderIndex":            "System.Header.GetIndex",
	"GetHeaderHash":             "System.Header.GetHash",
	"GetHeaderPrevHash":         "System.Header.GetPrevHash",
	"GetHeaderTimestamp":        "System.Header.GetTimestamp",
	"GetTransHash":              "System.Transaction.GetHash",
	"GetContractStorageContext": "System.Contract.GetStorageContext",
	"GetStorage":                "System.Storage.Get",
	"PutStorage":                "System.Storage.Put",
	"DeleteSorage":              "System.Storage.Delete",
	"GetStorageContext":         "System.Storage.GetContext",
	"GetReadOnlyStorageContext": "System.Storage.GetReadOnlyContext",
	"StorageCtxAsReadOnly":      "System.StorageContext.AsReadOnly",
	"RuntimeGetTime":            "System.Runtime.GetTime",
	"RuntimeCheckWitness":       "System.Runtime.CheckWitness",
	"RuntimeNotify":             "System.Runtime.Notify",
	"RuntimeLog":                "System.Runtime.Log",
	"RuntimeGetTrigger":         "System.Runtime.GetTrigger",
	"RuntimeSerialize":          "System.Runtime.Serialize",
	"RuntimeDeserialize":        "System.Runtime.Deserialize",

	// Storage API
	"GetContext": "Neo.Storage.GetContext",
	"Put":        "Neo.Storage.Put",
	"Get":        "Neo.Storage.Get",
	"Delete":     "Neo.Storage.Delete",

	// Runtime API
	"GetTrigger":      "Neo.Runtime.GetTrigger",
	"CheckWitness":    "Neo.Runtime.CheckWitness",
	"GetCurrentBlock": "Neo.Runtime.GetCurrentBlock",
	"GetTime":         "Neo.Runtime.GetTime",
	"Notify":          "Neo.Runtime.Notify",
	"Log":             "Neo.Runtime.Log",

	//Ont API
	"Invoke":                   "Ontology.Native.Invoke",
	"GetAttrUsage":             "Ontology.Attribute.GetUsage",
	"GetAttrData":              "Ontology.Attribute.GetData",
	"GetHeaderVersion":         "Ontology.Header.GetVersion",
	"GetHeaderNextConsensus":   "Ontology.Header.GetNextConsensus",
	"GetHeaderConsensusData":   "Ontology.Header.GetConsensusData",
	"GetHeaderMerkleRoot":      "Ontology.Header.GetMerkleRoot",
	"GetTransactionType":       "Ontology.Transaction.GetType",
	"GetTransactionAttributes": "Ontology.Transaction.GetAttributes",
	"CreateContract":           "Ontology.Contract.Create",
	"MigrateContract":          "Ontology.Contract.Migrate",
	"GetContractScript":        "Ontology.Contract.GetScript",
}
