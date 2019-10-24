package exchange

type RateLimiter struct {
	RateLimitType string
	Interval      string
	IntervalNum   int
	Limit         int
}

type SymbolFilter struct {
	FilterType  string
	MinPrice    string
	MaxPrice    string
	TickSize    string
	MultipierUp string
}

type ExchangeFilter struct {
	FilterType       string
	MaxNumOrders     int
	MaxNumAlgoOrders int
}

type SymbolInfo struct {
	Symbol string
	Status string
}

type ExchangeContainer struct {
	TimeZone        string
	ServerTime      int64
	RateLimits      []RateLimiter
	ExchangeFilters []ExchangeFilter
	Symbols         []SymbolInfo
}
