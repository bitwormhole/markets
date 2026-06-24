package dxo

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

func (code LicenceCode) Normalize() LicenceCode {
	return normalizeLicenceCode(code)
}

func (code LicenceCode) String() string {
	return string(code)
}

////////////////////////////////////////////////////////////////////////////////
