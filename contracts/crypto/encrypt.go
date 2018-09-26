package crypto

import "crypto/sha1"

func Main(operation string, args []interface{}) bool {
	if operation == "Test" {
		return Test()
	}
	return true
}

func Test() bool {
	sha1.Sum([]byte("123456789"))
	return true
}