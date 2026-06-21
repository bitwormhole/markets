package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Standard struct {

	// id
	ID dxo.StandardID `json:"id"`

	Base

	Code      dxo.StandardCode `json:"code"`
	URI       dxo.URI          `json:"uri"`
	Reference dxo.URL          `json:"ref"`
	Title     string           `json:"title"`
}
