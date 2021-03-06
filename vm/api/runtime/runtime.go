package runtime

import "github.com/ontio/neo-go-compiler/vm/api/types"

// CheckWitness verifies if the invoker is the owner of the contract.
func CheckWitness(hash []byte) bool {
	return true
}

// GetCurrentBlock returns the current block.
func GetCurrentBlock() types.Block { return types.Block{} }

// GetTime returns the timestamp of the most recent block.
func GetTime() int {
	return 0
}

func Notify(arg []interface{}) {}

// Log intructs the VM to log the given message.
func Log(message string) {}

// Application returns the application trigger type.
func Application() byte {
	return 0x10
}

// Verification returns the verification trigger type.
func Verification() byte {
	return 0x00
}

// GetTrigger return the current trigger type. The return in this function
// doesn't really mather, this is just an interop placeholder.
func GetTrigger() interface{} {
	return 0
}

func RuntimeGetTime() interface{}                    { return nil }
func RuntimeCheckWitness(hash []byte) bool           { return false }
func RuntimeNotify(msg []interface{})                {}
func RuntimeLog(msg interface{})                     {}
func RuntimeGetTrigger() interface{}                 { return nil }
func RuntimeSerialize(obj interface{}) interface{}   { return nil }
func RuntimeDeserialize(obj interface{}) interface{} { return nil }
