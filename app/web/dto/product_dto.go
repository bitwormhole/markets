package dto

import (
	"github.com/bitwormhole/markets/app/data/dxo"
)

type Product struct {

	// id
	ID dxo.ProductID `json:"id"`

	Base

	// core

	Name dxo.ProductName `json:"name"`
	Code dxo.ProductCode `json:"code"`

	Label       string `json:"label"`
	Remark      string `json:"remark"`
	Description string `json:"description"`

	// tm
	TrademarkID   dxo.TrademarkID   `json:"tm_id"`
	TrademarkCode dxo.TrademarkCode `json:"tm_code"`
	TrademarkName dxo.TrademarkName `json:"tm_name"`

	// std
	StandardID   dxo.StandardID   `json:"standard_id"`
	StandardCode dxo.StandardCode `json:"standard_code"`

	//refs
	URI       dxo.URI        `json:"uri"`
	URL       dxo.ProductURL `json:"url"`
	Reference dxo.URL        `json:"reference"`
}
