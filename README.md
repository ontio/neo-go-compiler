# Neo go compiler

this compiler is forked from [CityOfZion/neo-go](https://github.com/CityOfZion/neo-go)

## How to use
Use built binary file "neo-go-compile"

```./neo-go-compiler```

use ```--help``` to get detail information

## Neo and Ontology
Smart contract on Neo and Ontology has some different on their APIs

| No   | Ontology            | Neo             |      |
| ---- | ------------------- | --------------- | ---- |
| 1    | GetStorageContext   | GetContext      |      |
| 2    | PutStorage          | Put             |      |
| 3    | GetStorage          | Get             |      |
| 4    | DeleteSorage        | Delete          |      |
| 5    | RuntimeGetTrigger   | GetTrigger      |      |
| 6    | RuntimeCheckWitness | CheckWitness    |      |
| 7    |                     | GetCurrentBlock |      |
| 8    | RuntimeGetTime      | GetTime         |      |
| 9    | RuntimeNotify       | Notify          |      |
| 10   | RuntimeLog          | Log             |      |

To be continued...