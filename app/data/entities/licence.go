package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type Licence struct {

	// id
	ID dxo.LicenceID

	Base

	// fields

	Code dxo.LicenceCode `gorm:"unique"`

	Type dxo.LicenceType

	HolderID   dxo.CompanyID
	HolderCode dxo.CompanyCode
}
