package dxo

import "github.com/bitwormhole/markets/app/data/normalizers"

// 表示商品名称
type ProductName string

// 表示商品条码
type ProductCode string

// 表示商品官网页面
type ProductURL URL

////////////////////////////////////////////////////////////////////////////////

func (code ProductCode) Pure() ProductCode {
	str1 := code.String()
	str2 := normalizers.PurifyCode(str1)
	return ProductCode(str2)
}

func (code ProductCode) ForDomain(dn DomainName) ProductCode {
	str1 := code.Pure().String()
	str2 := dn.String()
	str3 := str1 + "@" + str2
	return ProductCode(str3)
}

func (code ProductCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////
