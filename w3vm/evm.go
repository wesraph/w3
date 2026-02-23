package w3vm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EVM is the interface for a custom EVM execution backend.
type EVM interface {
	// Execute executes the given message. If isCall is true, the execution
	// should not persist state changes.
	Execute(msg *EVMMessage, isCall bool) (*EVMResult, error)

	Nonce(addr common.Address) (uint64, error)
	Balance(addr common.Address) (*big.Int, error)
	Code(addr common.Address) ([]byte, error)
	StorageAt(addr common.Address, slot common.Hash) (common.Hash, error)

	SetNonce(addr common.Address, nonce uint64)
	SetBalance(addr common.Address, balance *big.Int)
	SetCode(addr common.Address, code []byte)
	SetStorageAt(addr common.Address, slot, val common.Hash)

	Clone() EVM
}

// EVMMessage is a fully resolved EVM message.
type EVMMessage struct {
	From                  common.Address
	To                    *common.Address
	Nonce                 uint64
	Value                 *big.Int
	GasLimit              uint64
	GasPrice              *big.Int
	GasFeeCap             *big.Int
	GasTipCap             *big.Int
	Input                 []byte
	AccessList            types.AccessList
	BlobGasFeeCap         *big.Int
	BlobHashes            []common.Hash
	SetCodeAuthorizations []types.SetCodeAuthorization

	SkipNonceChecks       bool
	SkipTransactionChecks bool
}

// EVMResult is the result of an EVM execution.
type EVMResult struct {
	UsedGas         uint64
	ReturnData      []byte
	Logs            []*types.Log
	Err             error
	ContractAddress *common.Address
}

// EVMFetcherSetter is an optional interface for EVM backends that accept
// a [Fetcher] for lazy state loading.
type EVMFetcherSetter interface {
	SetFetcher(Fetcher)
}
