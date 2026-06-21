package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Trademark struct {

	// id
	ID dxo.TrademarkID `json:"id"`

	Base

	Code dxo.TrademarkCode `json:"code"`

	URI dxo.URI `json:"uri"`
}
