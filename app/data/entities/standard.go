package entities

import "github.com/bitwormhole/markets/app/data/dxo"

type BaseStandard struct {

	// id
	// ID dxo.StandardID

	Base

	// fields

	Code dxo.StandardCode

	UniqueCode dxo.StandardCode `gorm:"unique"`

	Reference dxo.URL // 作为参考 (数据来源) 的 web 页面

	URI dxo.URI `gorm:"unique"`
}

type Standard struct {

	// id
	ID dxo.StandardID

	BaseStandard
}

// type UserStandard struct {
// 	// id
// 	ID dxo.UserStandardID
// 	BaseStandard
// }
