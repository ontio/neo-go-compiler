package blockchain

import "neo-go-compiler/vm/api/types"

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
