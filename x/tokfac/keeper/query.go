package keeper

import (
	"github.com/cosmos/cosmos-sdk/x/tokfac/types"
)

var _ types.QueryServer = Keeper{}
