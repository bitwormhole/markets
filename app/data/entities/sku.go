package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type SKU struct {

	// id
	ID dxo.SkuID

	Base

	// fields

	Shop dxo.ShopID
}
