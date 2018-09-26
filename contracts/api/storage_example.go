package api

import (
	"github.com/ontio/neo-go-compiler/vm/api/storage"
	"github.com/ontio/neo-go-compiler/vm/api/runtime"
)

func Main(operation string, args []interface{}) bool {
	if operation == "storageTest" {
		return storageTest()
	}
	return false

}

func storageTest() bool {
	ctx := storage.GetContext()

	storage.Put(ctx, "key", 100)
	v := storage.Get(ctx, "key").(uint64)
	runtime.Notify([]interface{}{v})

	storage.Put(ctx, []byte("key"), 100)
	v2 := storage.Get(ctx, []byte("key")).(uint64)
	runtime.Notify([]interface{}{v2})

	storage.Delete(ctx, "key")
	v3 := storage.Get(ctx, "key")
	runtime.Notify([]interface{}{v3})
	return true
}