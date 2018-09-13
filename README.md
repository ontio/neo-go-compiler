# Introduction 

this compiler is forked from [CityOfZion/neo-go](https://github.com/CityOfZion/neo-go), compile go smart contract code to avm file.

## How to use
Use built binary file "neo-go-compile"

```./neo-go-compiler```

use ```--help``` to get detail information

### Compile smart contract
```
./neo-go-compiler contract compile -i <path/contract.go> [-abi] [-o <out put path/file.avm>]
```
-o is optional to indicate the output avm file path

-abi is generage abi file or not, if yes, you will find the <contract>.json under the same directory

### Dump opcode

```
./neo-go-compiler contract opdump -i <path/contract.go> 
```

output :
```
INDEX    OPCODE    DESC             
0        0x56      Opush6           
1        0xc5      Onewarray        
2        0x6b      Otoaltstack      
3        0x6c      Ofromaltstack    
4        0x76      Odup             
5        0x6b      Otoaltstack      
6        0x 0      Opush0           
...

```

## Neo and Ontology
Smart contract for Neo and Ontology has some difference on their APIs

| No   | Ontology            | Neo             |  comments    |
| ---- | ------------------- | --------------- | ------------ |
| 1    | GetStorageContext   | GetContext      |  get storage context    |
| 2    | PutStorage          | Put             |  save to storage        |
| 3    | GetStorage          | Get             |  query from storage     |
| 4    | DeleteSorage        | Delete          |  delete from storage    |
| 5    | RuntimeGetTrigger   | GetTrigger      |  get trigger            |
| 6    | RuntimeCheckWitness | CheckWitness    |  check witness          |
| 7    |       -             | GetCurrentBlock |  get current block      |
| 8    | RuntimeGetTime      | GetTime         |  get time stamp         |
| 9    | RuntimeNotify       | Notify          |  add a notify           |
| 10   | RuntimeLog          | Log             |  add a log              |
| 11   | Invoke              |    -            |  invoke native contract |
| 12   | GetAttrUsage        |    -            |  get attribute usage    |
| 13   | GetAttrData         |    -            |  get attribute data     |
| 14   | GetHeaderVersion    |    -            |  get header version     |
| 15   | GetHeaderNextConsensus    |   -       |  get header next consensus     |
| 16   | GetHeaderConsensusData    |   -       |  get header consensus data     |
| 17   | GetHeaderMerkleRoot |        -        |  get header merkle root     |
| 18   | GetTransactionType  |        -        |  get transaction type     |
| 19   | GetTransactionAttributes  |   -       |  get transaction attributes     |
| 20   | MigrateContract  |           -        |  migrate contract       |
| 21   | GetContractScript   |        -        |  get contract script     |

## Other APIs
the following apis supports both Ontology and Neo

| No    | Name                | comments             |
| ----- | ------------------- | -------------------- |
| 1     | GetScriptContainer       | get script container           |
| 2     | GetExecutingScriptHash   | get current contract hash      |
| 3     | GetCallingScriptHash     | get calling contract hash      |
| 4     | GetEntryScriptHash       | get entry script hash          |
| 5     | GetBlockTransactionCount | get transactions count of block          |
| 6     | GetBlockTransactions     | get transactions of block           |
| 7     | GetBlockTransaction      | get a specified transaction           |
| 8     | GetBlockchainHeight      | get current block height           |
| 9     | GetBlockchainHeader      | get block header           |
| 10    | GetBlockchainBlock       | get block                      |
| 11    | GetContract              | get contract                      |
| 12    | GetTransactionHeight     | get transaction height                      |
| 13    | GetHeaderIndex           | get header height                      |
| 14    | GetHeaderHash            | get header hash                      |
| 15    | GetHeaderPrevHash        | get header previous hash                      |
| 16    | GetHeaderTimestamp       | get header timestamp                      |
| 17    | GetTransHash             | get transaction hash                      |
| 18    | GetContractStorageContext| get contract storage context              |
| 19    | GetReadOnlyStorageContext| get readonly storage context                      |
| 20    | StorageCtxAsReadOnly     | make storage readonly                      |
| 21    | RuntimeSerialize         | serialize collection to bytearray                      |
| 22    | RuntimeDeserialize       | deserialize bytearray                      |


## How to write smart contract in golang

You can write your smart contract code in any editors which provides the golang 
grammar check.

The entry function in golang contract must be like following:
```go
func Main(operation string, args []interface{}) interface{} {
	if operation == "someoperation1"{ 
        do something 
    }
    if operation == "someoperation2"{
        do something
    }
    ...
}

```

"operation" represents the actual method name to invoke

"args" represents the parameters to be passed in .

### Supported Types

Golang neovm contract is a subset of golang,  only support the following types:

string , int , int64, []byte, byte, bool 



### Not supported 

1. Multiple return values
2. Error
3. for ... range   (Ontology need to support "KEYS" opcode)
4. go routine
5. get or set elements from a byte array



### Address

Generally, all addresses in smart contract should be ```[]byte```  

1. As parameter: like C# and python ,you need to pass the Hex format address parameters .

   ```
   addr := args[0].([]byte)
   ```

   ​

2. As constant variables

   ```
   contractAddr:=[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}

   ```

3. Use tools.ToScriptHash()

   ```
   owner = tools.ToScriptHash("Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT")
   ```

   ***Note**: the parameter of ToScriptHash must be a constant string

   ​

### Storage

first , you need to declare a storage context :

Ontology

```go
ctx = storage.GetStorageContext()
```

NEO

```go
ctx = storage.GetContext()
```



then calling Put, Get or Delete method to access the storage

Ontology

```go
storage.PutStorage(ctx, key, val)
result := storage.GetStorage(ctx,key)
storage.DeleteStorage(ctx,key)
```

NEO

```
storage.Put(ctx, key, val)
result := storage.Get(ctx,key)
storage.Delete(ctx,key)
```



### Transfer ONT/ONG

In smart contract , you need to invoke native contract to transfer ONT and ONG

```go
func transONT(from []byte, to []byte, amount int64) bool {
	if runtime.RuntimeCheckWitness(from) == false{return false}
	contractAddr:=[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	param := appcall.State{from,to,amount}
	ver := 1
	bs :=native.Invoke(ver,contractAddr,"transfer",[]interface{}{param})
	if bs != nil && tools.BytesEquals(bs,[]byte("1")){
		return true
	}else{
		return false
	}
}

```

You can also define a struct to carry transfer parameters :

```go
type transferParam struct{
	From []byte
	To []byte
	Amount int64
}
...
param := transferParam(from,to,amount)
```



### AppCall

In order to call an other smart contract, you need to use appcall.AppCall:

```go
 appcall.AppCall("APPWgNbWvUdQjQxeN7RduYweH3caaM1LM1","transfer",args)
 appcall.AppCall("83e69795f9c314a8c4f483e221927f41285a8653","transfer",args)

```

the first parameter must be constant string, either hex or base58 format address is acceptable.

***Note: Do not define the address string like this:**

```go
addr := "APPWgNbWvUdQjQxeN7RduYweH3caaM1LM1"
appcall.AppCall(addr,"transfer",args)
```



the second parameter is method name to call.

the last parameter is the arguments .

 

### Tools

Here are some utility functions :

BytesEquals: to compare to byte arrays equality. (you can't loop the byte array to compare every single byte in neovm)

```go
	if bs != nil && tools.BytesEquals(bs,[]byte("1")){
		return true
	}else{
		return false
	}
```

ToScriptHash: refer to the Address section

Cat: concat two byte arrays :

```go
newbytes:=tools.Cat(transfer_prefix, owner)
```



### Other restrictions

expressions with in ```if``` statement must compare explicitly

```go
if runtime.RuntimeCheckWitness(from) == false {
		return false
}
```



### Examples:

please check the examples under "contracts" directory