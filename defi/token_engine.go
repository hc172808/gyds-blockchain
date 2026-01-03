package defi

type Token struct {
	Name        string
	Symbol      string
	Decimals    int
	TotalSupply uint64
}

// TODO: implement token mint, burn, and native coin management
