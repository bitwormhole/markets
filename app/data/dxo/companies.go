package dxo

// 表示企业名称
type CompanyName string

// 表示企业的统一社会信用代码
type CompanyCode string

// 表示企业的官网主页地址
type CompanyURL URL

////////////////////////////////////////////////////////////////////////////////

func (name CompanyName) Normalize() CompanyName {

	// todo : 'noimpl'

	return name

}

func (name CompanyName) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////

func (code CompanyCode) Normalize() CompanyCode {
	return normalizeCompanyCode(code)
}

func (code CompanyCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////
