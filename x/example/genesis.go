package example

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v4/x/example/keeper"
	"github.com/umee-network/umee/v4/x/example/types"
)

// InitGenesis initializes the x/example module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, genState types.GenesisState) {
}

// ExportGenesis returns the x/example module's exported genesis.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{}
}
