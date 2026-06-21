package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type BaseTrademark struct {

	// id
	ID dxo.TrademarkID

	Base

	// fields

	Name string

	Type string

	Category string

	Code dxo.TrademarkCode // 商标注册码

	HolderID   dxo.CompanyID
	HolderCode dxo.CompanyCode
	HolderName dxo.CompanyName

	Reference dxo.URL // 作为参考 (数据来源) 的 web 页面

	URI dxo.URI `gorm:"unique"`
}

type Trademark struct {

	// id
	ID dxo.TrademarkID

	BaseTrademark
}

// type UserTrademark struct {
// 	// id
// 	ID dxo.UserTrademarkID
// 	BaseTrademark
// }
