package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Product struct {

	// id
	ID dxo.ProductID `json:"id"`

	Base

	Name dxo.ProductName `json:"name"`
	Code dxo.ProductCode `json:"code"`
	URL  dxo.ProductURL  `json:"url"`

	Label       string `json:"label"`
	Description string `json:"description"`
}
