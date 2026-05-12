package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Domain struct {

	// id
	ID dxo.DomainID `json:"id"`

	Base

	Name dxo.DomainName `json:"name"`
}
