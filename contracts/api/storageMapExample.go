package api

import "github.com/ontio/neo-go-compiler/vm/api/storage"

func Main(operation string, args []interface{}) bool {
	if operation == "StorageMapTest"{
		return StorageMapTest()
	}
	return false
}

func StorageMapTest() bool {
	asset := storage.GetContext()
	return true
}
