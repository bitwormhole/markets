package dxo

import "github.com/bitwormhole/markets/app/data/normalizers"

// 表示企业名称
type CompanyName string

// 表示企业的统一社会信用代码
type CompanyCode string

// 表示企业的官网主页地址
type CompanyURL URL

////////////////////////////////////////////////////////////////////////////////

func (name CompanyName) Normalize() CompanyName {

	str := name.String()
	str = normalizers.NormalizeRegularName(str)
	return CompanyName(str)
}

func (name CompanyName) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////

func (code CompanyCode) Pure() CompanyCode {
	str1 := code.String()
	str2 := normalizers.PurifyCode(str1)
	return CompanyCode(str2)
}

func (code CompanyCode) ForDomain(dn DomainName) CompanyCode {
	str1 := code.Pure().String()
	str2 := dn.String()
	str3 := str1 + "@" + str2
	return CompanyCode(str3)
}

func (code CompanyCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////
