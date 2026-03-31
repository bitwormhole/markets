package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type Standard struct {

	// id
	ID dxo.StandardID

	Base

	// fields

	Code dxo.StandardCode
}
