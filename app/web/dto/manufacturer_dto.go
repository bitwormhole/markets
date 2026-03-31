package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Manufacturer struct {

	// id
	ID dxo.ManufacturerID `json:"id"`

	Base

	CompanyID dxo.CompanyID `json:"company_id"`
	ProductID dxo.ProductID `json:"product_id"`
	LicenceID dxo.LicenceID `json:"licence_id"`

	CompanyCode dxo.CompanyCode `json:"company_code"` // 企业统一信用代码
	ProductCode dxo.ProductCode `json:"product_code"` // 商品条码
	LicenceCode dxo.LicenceCode `json:"licence_code"` // 生产许可证代码
}
