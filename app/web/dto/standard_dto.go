package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Standard struct {

	// id
	ID dxo.StandardID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
