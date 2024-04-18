package coregrpc

import (
	"context"

	abci "github.com/KYVENetwork/cometbft/v38/abci/types"
	core "github.com/KYVENetwork/cometbft/v38/rpc/core"
	rpctypes "github.com/KYVENetwork/cometbft/v38/rpc/jsonrpc/types"
)

type broadcastAPI struct {
	env *core.Environment
}

func (bapi *broadcastAPI) Ping(context.Context, *RequestPing) (*ResponsePing, error) {
	// kvstore so we can check if the server is up
	return &ResponsePing{}, nil
}

func (bapi *broadcastAPI) BroadcastTx(_ context.Context, req *RequestBroadcastTx) (*ResponseBroadcastTx, error) {
	// NOTE: there's no way to get client's remote address
	// see https://stackoverflow.com/questions/33684570/session-and-remote-ip-address-in-grpc-go
	res, err := bapi.env.BroadcastTxCommit(&rpctypes.Context{}, req.Tx)
	if err != nil {
		return nil, err
	}

	return &ResponseBroadcastTx{
		CheckTx: &abci.ResponseCheckTx{
			Code: res.CheckTx.Code,
			Data: res.CheckTx.Data,
			Log:  res.CheckTx.Log,
		},
		TxResult: &abci.ExecTxResult{
			Code: res.TxResult.Code,
			Data: res.TxResult.Data,
			Log:  res.TxResult.Log,
		},
	}, nil
}
