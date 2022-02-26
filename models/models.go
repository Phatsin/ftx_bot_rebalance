package models

type Assign struct {
	Account    string  `json:"account,omitempty"`
	Instrument string  `json:"instrument,omitempty"`
	Fvalue     int64   `json:"fvalue,omitempty"`
	Bdist      float64 `json:"bdist,omitempty"`
	Sdist      float64 `json:"sdist,omitempty"`
	Prest      float64 `json:"prest,omitempty"`
}

type Sender_Target struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type resData struct {
	OrderId                int         `json:"orderId,omitempty"`
	TradeId                int         `json:"tradeId,omitempty"`
	Id                     int         `json:"id,omitempty"`
	ClientId               int         `json:"clientId,omitempty"`
	PriceIncrement         float32     `json:"priceIncrement,omitempty"`
	SizeIncrement          float32     `json:"sizeIncrement,omitempty"`
	Free                   float32     `json:"free,omitempty"`
	SpotBorrow             float32     `json:"spotBorrow,omitempty"`
	Total                  float32     `json:"total,omitempty"`
	UsdValue               float32     `json:"usdValue,omitempty"`
	AvailableWithoutBorrow float32     `json:"availableWithoutBorrow,omitempty"`
	Size                   float32     `json:"size,omitempty"`
	Price                  float32     `json:"price,omitempty"`
	Bid                    float32     `json:"bid,omitempty"`
	Ask                    float32     `json:"ask,omitempty"`
	BidSize                float32     `json:"bidSize,omitempty"`
	AskSize                float32     `json:"askSize,omitempty"`
	Last                   float32     `json:"last,omitempty"`
	FilledSize             float32     `json:"filledSize,omitempty"`
	RemainingSize          float32     `json:"remainingSize,omitempty"`
	AvgFillPrice           float32     `json:"avgFillPrice,omitempty"`
	Fee                    float32     `json:"fee,omitempty"`
	FeeRate                float32     `json:"feeRate,omitempty"`
	Name                   string      `json:"name,omitempty"`
	BaseCurrency           string      `json:"baseCurrency,omitempty"`
	QuoteCurrency          string      `json:"quoteCurrency,omitempty"`
	Underlying             string      `json:"underlying,omitempty"`
	Coin                   string      `json:"coin,omitempty"`
	Type                   string      `json:"type,omitempty"`
	Market                 string      `json:"market,omitempty"`
	Future                 string      `json:"future,omitempty"`
	Side                   string      `json:"side,omitempty"`
	Status                 string      `json:"status,omitempty"`
	Enabled                bool        `json:"enabled,omitempty"`
	Restricted             bool        `json:"restricted,omitempty"`
	Liquidity              bool        `json:"liquidity,omitempty"`
	ReduceOnly             bool        `json:"reduceOnly,omitempty"`
	Ioc                    bool        `json:"ioc,omitempty"`
	PostOnly               bool        `json:"postOnly,omitempty"`
	CreatedAt              interface{} `json:"createdAt,omitempty"`
	Time                   interface{} `json:"time,omitempty"`
}

type Req struct {
	Type    string   `json:"type,omitempty"`
	Channel string   `json:"channel,omitempty"`
	Market  string   `json:"market,omitempty"`
	Data    *resData `json:"data,omitempty"`
}

type GetData struct {
	Result  []resData `json:"result,omitempty"`
	Error   string    `json:"error,omitempty"`
	Success bool      `json:"success"`
}

type GetData_ struct {
	Result  resData `json:"result,omitempty"`
	Error   string  `json:"error,omitempty"`
	Success bool    `json:"success"`
}
