package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type BaseShop struct {

	// id
	// ID dxo.ShopID

	Base

	// fields

	Name dxo.ShopName

	Description string

	OperatorID   dxo.CompanyID
	OperatorCode dxo.CompanyCode

	Web dxo.URL
}

type Shop struct {

	// id
	ID dxo.ShopID

	BaseShop
}

// type UserShop struct {
// 	// id
// 	ID dxo.UserShopID
// 	BaseShop
// }
