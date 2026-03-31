package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Company struct {

	// id
	ID dxo.CompanyID `json:"id"`

	Base

	Name dxo.CompanyName `json:"name"`
	Code dxo.CompanyCode `json:"code"`
}
