package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type BaseProduct struct {

	// id
	// ID dxo.ProductID

	Base

	// fields

	Name dxo.ProductName

	URL dxo.ProductURL

	Code dxo.ProductCode

	UniqueCode dxo.ProductCode `gorm:"unique"`

	Label       string
	Description string

	Standard dxo.StandardID

	Reference dxo.URL // 作为参考 (数据来源) 的 web 页面

	URI dxo.URI `gorm:"unique"`
}

type Product struct {

	// id
	ID dxo.ProductID

	BaseProduct
}

// type UserProduct struct {
// 	// id
// 	ID dxo.UserProductID
// 	BaseProduct
// }
