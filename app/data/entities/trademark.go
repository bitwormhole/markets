package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type Trademark struct {

	// id
	ID dxo.TrademarkID

	Base

	// fields

	Name string

	Type string

	Category string

	HolderID   dxo.CompanyID
	HolderCode dxo.CompanyCode
}
