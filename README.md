# Neo go compiler

this compiler is forked from [CityOfZion/neo-go](https://github.com/CityOfZion/neo-go)

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


## How to write smart contract in GO

You can write your smart contract code in any editors which provides the golang 
grammar check.

The entry function in golang contract must be like following:
```go
func Main(operation string, args []interface{}) interface{} {
}




```


To be continued...