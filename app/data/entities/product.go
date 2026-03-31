package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type Product struct {

	// id
	ID dxo.ProductID

	Base

	// fields

	Name dxo.ProductName
	Code dxo.ProductCode `gorm:"unique"`
	URL  dxo.ProductURL

	Label       string
	Description string

	Standard dxo.StandardID
}
