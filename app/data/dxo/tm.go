package dxo

// 表示商标注册号
type TrademarkCode string

// 表示商标的正式名称
type TrademarkName string

////////////////////////////////////////////////////////////////////////////////

func (code TrademarkCode) Normalize() TrademarkCode {
	return normalizeTrademarkCode(code)
}

func (code TrademarkCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////

func (name TrademarkName) Normalize() TrademarkName {
	return normalizeTrademarkName(name)
}

func (name TrademarkName) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////
