package vm

// Syscalls are a mapping between the syscall function name
// and the registerd VM interop API.
var Syscalls = map[string]string{
	//System API
	"GetScriptContainer":     "System.ExecutionEngine.GetScriptContainer",
	"GetExecutingScriptHash": "System.ExecutionEngine.GetExecutingScriptHash",
	"GetCallingScriptHash":   "System.ExecutionEngine.GetCallingScriptHash",
	"GetEntryScriptHash":     "System.ExecutionEngine.GetEntryScriptHash",
	"SerializeContract":      "System.ExecutionEngine.SerializeContract", //temp for test
	"SerializeTransfer":      "System.ExecutionEngine.SerializeTransfer", //temp for test

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
}
