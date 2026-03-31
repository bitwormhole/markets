package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type Shop struct {

	// id
	ID dxo.ShopID

	Base

	// fields

	Name dxo.ShopName

	Description string

	URL dxo.ShopURL

	OperatorID   dxo.CompanyID
	OperatorCode dxo.CompanyCode
}
