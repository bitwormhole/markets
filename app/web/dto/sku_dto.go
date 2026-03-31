package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type SKU struct {

	// id
	ID dxo.SkuID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
