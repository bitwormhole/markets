package dto

import "github.com/bitwormhole/markets/app/data/dxo"

type Licence struct {

	// id
	ID dxo.LicenceID `json:"id"`

	Base

	Code dxo.LicenceCode `json:"code"`
	Type dxo.LicenceType `json:"type"`
}
