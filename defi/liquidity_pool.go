package defi

type LiquidityPool struct {
	AssetA  string
	AssetB  string
	Balance map[string]uint64
}

// TODO: Add functions to add/remove liquidity and calculate swaps
