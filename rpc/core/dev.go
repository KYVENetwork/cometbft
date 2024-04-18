package core

import (
	ctypes "github.com/KYVENetwork/cometbft/v34/rpc/core/types"
	rpctypes "github.com/KYVENetwork/cometbft/v34/rpc/jsonrpc/types"
)

// UnsafeFlushMempool removes all transactions from the mempool.
func UnsafeFlushMempool(ctx *rpctypes.Context) (*ctypes.ResultUnsafeFlushMempool, error) {
	env.Mempool.Flush()
	return &ctypes.ResultUnsafeFlushMempool{}, nil
}
