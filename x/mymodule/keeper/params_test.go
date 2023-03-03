package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mymodule/testutil/keeper"
	"mymodule/x/mymodule/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MymoduleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
