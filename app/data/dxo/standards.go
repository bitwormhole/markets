package dxo

import "github.com/bitwormhole/markets/app/data/normalizers"

// 表示标准代码, 例如: GB25190
type StandardCode string

////////////////////////////////////////////////////////////////////////////////

func (code StandardCode) Pure() StandardCode {
	str1 := code.String()
	str2 := normalizers.PurifyCode(str1)
	return StandardCode(str2)
}

func (code StandardCode) ForDomain(dn DomainName) StandardCode {
	str1 := code.Pure().String()
	str2 := dn.String()
	str3 := str1 + "@" + str2
	return StandardCode(str3)
}

func (code StandardCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////
