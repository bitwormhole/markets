package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Trademark struct {

	// id
	ID dxo.TrademarkID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
