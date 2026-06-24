package dxo

// 表示商品名称
type ProductName string

// 表示商品条码
type ProductCode string

// 表示商品官网页面
type ProductURL URL

////////////////////////////////////////////////////////////////////////////////

func (code ProductCode) Normalize() ProductCode {
	return normalizeProductCode(code)
}

func (code ProductCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////
