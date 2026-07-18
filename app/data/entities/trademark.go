package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type BaseTrademark struct {

	// id
	ID dxo.TrademarkID

	Base

	// fields

	Name dxo.TrademarkName

	Code dxo.TrademarkCode // 商标注册码

	Type string

	Category string

	// holder

	HolderID   dxo.CompanyID
	HolderCode dxo.CompanyCode
	HolderName dxo.CompanyName

	// tm

	//refs

	Reference dxo.URL // 作为参考 (数据来源) 的 web 页面

	URI dxo.URI `gorm:"unique"`
}

type Trademark struct {

	// id
	ID dxo.TrademarkID

	BaseTrademark
}
