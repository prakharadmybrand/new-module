package keeper

import (
	"mymodule/x/mymodule/types"
)

var _ types.QueryServer = Keeper{}
