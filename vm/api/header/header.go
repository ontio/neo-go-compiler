package header

func GetHeaderIndex(hash interface{}) int               { return 0 }
func GetHeaderHash(index int) interface{}               { return nil }
func GetHeaderPrevHash(header interface{}) interface{}  { return nil }
func GetHeaderTimestamp(header interface{}) interface{} { return nil }

//ONT
func GetHeaderVersion(header interface{}) int               { return 0 }
func GetHeaderNextConsensus(header interface{}) interface{} { return nil }
func GetHeaderConsensusData(header interface{}) interface{} { return nil }
func GetHeaderMerkleRoot(header interface{}) interface{}    { return nil }
