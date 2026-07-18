package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type BaseProduct struct {

	// id
	// ID dxo.ProductID

	Base

	// core

	Name dxo.ProductName
	Code dxo.ProductCode

	Label       string
	Remark      string
	Description string

	// Standard
	StandardID   dxo.StandardID
	StandardCode dxo.StandardCode

	// tm
	TrademarkID   dxo.TrademarkID
	TrademarkCode dxo.TrademarkCode
	TrademarkName dxo.TrademarkName

	//refs
	URL       dxo.ProductURL
	URI       dxo.URI `gorm:"unique"`
	Reference dxo.URL // 作为参考 (数据来源) 的 web 页面

}

type Product struct {

	// id
	ID dxo.ProductID

	BaseProduct
}
