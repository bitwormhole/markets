package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type Domain struct {

	// id
	ID dxo.DomainID

	Base

	// fields

	Name dxo.DomainName `gorm:"unique"`
}
