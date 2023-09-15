package types

type GetContractDetailsParams struct {
	Symbol string `url:"symbols,omitempty" validate:"omitempty"`
}

type GetContractDetailsResp struct {
	Response
	Data []*ContractDetail `json:"data"`
}

type ContractDetail struct {
	Symbol                    string   `json:"symbol"`
	DisplayName               string   `json:"displayName"`
	DisplayNameEn             string   `json:"displayNameEn"`
	PositionOpenType          int      `json:"positionOpenType"`
	BaseCoin                  string   `json:"baseCoin"`
	QuoteCoin                 string   `json:"quoteCoin"`
	SettleCoin                string   `json:"settleCoin"`
	ContractSize              float64  `json:"contractSize"`
	MinLeverage               float64  `json:"minLeverage"`
	MaxLeverage               float64  `json:"maxLeverage"`
	PriceScale                float64  `json:"priceScale"`
	VolScale                  float64  `json:"volScale"`
	AmountScale               float64  `json:"amountScale"`
	PriceUnit                 float64  `json:"priceUnit"`
	VolUnit                   float64  `json:"volUnit"`
	MinVol                    float64  `json:"minVol"`
	MaxVol                    float64  `json:"maxVol"`
	BidLimitPriceRate         float64  `json:"bidLimitPriceRate"`
	AskLimitPriceRate         float64  `json:"askLimitPriceRate"`
	TakerFeeRate              float64  `json:"takerFeeRate"`
	MakerFeeRate              float64  `json:"makerFeeRate"`
	MaintenanceMarginRate     float64  `json:"maintenanceMarginRate"`
	InitialMarginRate         float64  `json:"initialMarginRate"`
	RiskBaseVol               float64  `json:"riskBaseVol"`
	RiskIncrVol               float64  `json:"riskIncrVol"`
	RiskIncrMmr               float64  `json:"riskIncrMmr"`
	RiskIncrImr               float64  `json:"riskIncrImr"`
	RiskLevelLimit            float64  `json:"riskLevelLimit"`
	PriceCoefficientVariation float64  `json:"priceCoefficientVariation"`
	IndexOrigin               []string `json:"indexOrigin"`
	State                     int      `json:"state"`
	ApiAllowed                bool     `json:"apiAllowed"`
	ConceptPlate              []string `json:"conceptPlate"`
	RiskLimitType             string   `json:"riskLimitType"`
}
