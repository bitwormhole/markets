package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Shop struct {

	// id
	ID dxo.ShopID `json:"id"`

	Base

	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	URI         dxo.URI `json:"uri"`
}
