package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type BaseSKU struct {

	// id
	// ID dxo.SkuID

	Base

	// fields

	Shop dxo.ShopID

	Web dxo.URL
}

type SKU struct {

	// id
	ID dxo.SkuID

	BaseSKU

	URI dxo.URI `gorm:"unique"`
}
