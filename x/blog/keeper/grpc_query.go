package keeper

import (
	"github.com/mmulji-ic/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
