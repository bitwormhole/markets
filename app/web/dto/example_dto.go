package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Example struct {

	// id
	ID dxo.ExampleID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
