package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Shop struct {

	// id
	ID dxo.ShopID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
