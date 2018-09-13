package blockchain

import "github.com/ontio/neo-go-compiler/vm/api/types"

// GetHeight returns current block height.
func GetHeight() uint { return 0 }

// GetHeader returns specific block height header.
func GetHeader(height uint) types.Header { return types.Header{} }

// GetBlock get block by block hash.
func GetBlock(hash []byte) types.Block { return types.Block{} }

// GetTransaction returns the transaction by txid.
func GetTransaction(txid []byte) types.Transaction { return types.Transaction{} }

// GetContract returns smart contract by txid.
func GetContract(script_hash []byte) types.Contract { return types.Contract{} }

func GetBlockchainHeight() int                               { return 0 }
func GetBlockchainHeader(blockchain interface{}) interface{} { return nil }
func GetBlockchainBlock(blockchain interface{}) interface{}  { return nil }
func GetBlockchainTransaction(hash interface{}) interface{}  { return nil }
