package entities

import (
	"time"

	"github.com/bitwormhole/markets/app/data/dxo"
)

type BaseCompany struct {

	// id
	// ID dxo.CompanyID

	Base

	// fields

	Address string

	ApprovedAt time.Time // 核准日期

	Capital string // 注册资本

	Code dxo.CompanyCode

	CompanyType string // 类型

	FoundedAt time.Time // 成立日期

	Name dxo.CompanyName

	OperationCategory string // 经营类别

	OperationRange string // 经营范围

	Reference dxo.URL // 作为参考 (数据来源) 的 web 页面

	Registry string // 登记机关

	Remarks string // 备注

	Representative string //  法定代表人

	State string // 登记状态

	URI dxo.URI `gorm:"unique"`

	Web dxo.CompanyURL
}

type Company struct {

	// id
	ID dxo.CompanyID

	BaseCompany
}

// type UserCompany struct {
// 	// id
// 	ID dxo.UserCompanyID
// 	BaseCompany
// }
