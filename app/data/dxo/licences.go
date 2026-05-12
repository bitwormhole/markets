package dxo

import "github.com/bitwormhole/markets/app/data/normalizers"

// 表示许可证类型
type LicenceType string

// 表示许可证编码
type LicenceCode string

// 表示许可证信息网页地址
type LicenceURL URL

////////////////////////////////////////////////////////////////////////////////

const (
	LicenceYYZZ LicenceType = "yyzz" // 营业执照
	LicenceSC   LicenceType = "sc"   // 食品生产许可证
	LicenceXK   LicenceType = "xk"   // 生产许可证
	LicenceWS   LicenceType = "ws"   // 卫生许可证
	LicenceTM   LicenceType = "tm"   // 商标注册证书
)

////////////////////////////////////////////////////////////////////////////////

func (code LicenceCode) Pure() LicenceCode {
	str1 := code.String()
	str2 := normalizers.PurifyCode(str1)
	return LicenceCode(str2)
}

func (code LicenceCode) ForDomain(dn DomainName) LicenceCode {
	str1 := code.Pure().String()
	str2 := dn.String()
	str3 := str1 + "@" + str2
	return LicenceCode(str3)
}

func (code LicenceCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////
