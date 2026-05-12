package dxo

import "github.com/bitwormhole/markets/app/data/normalizers"

// 表示商标注册号
type TrademarkCode string

////////////////////////////////////////////////////////////////////////////////

func (code TrademarkCode) Pure() TrademarkCode {
	str1 := code.String()
	str2 := normalizers.PurifyCode(str1)
	return TrademarkCode(str2)
}

func (code TrademarkCode) ForDomain(dn DomainName) TrademarkCode {
	str1 := code.Pure().String()
	str2 := dn.String()
	str3 := str1 + "@" + str2
	return TrademarkCode(str3)
}

func (code TrademarkCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////
