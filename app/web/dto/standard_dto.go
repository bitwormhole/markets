package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Standard struct {

	// id
	ID dxo.StandardID `json:"id"`

	Base

	Code        dxo.StandardCode `json:"code"`
	URI         dxo.URI          `json:"uri"`
	Refs        dxo.URLList      `json:"refs"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
}
