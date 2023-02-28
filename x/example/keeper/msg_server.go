package keeper

import (
	"context"

	"github.com/umee-network/umee/v4/x/example/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the example MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// Example any example msg.
func (ms msgServer) Example(
	_ context.Context,
	_ *types.MsgExample,
) (*types.MsgExampleResponse, error) {
	return &types.MsgExampleResponse{}, nil
}
