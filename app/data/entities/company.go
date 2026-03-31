package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type Company struct {

	// id
	ID dxo.CompanyID

	Base

	// fields

	Name dxo.CompanyName

	Code dxo.CompanyCode `gorm:"unique"`

	Address string

	URL dxo.CompanyURL
}
