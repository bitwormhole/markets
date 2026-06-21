package dto

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/base/lang"
)

type Company struct {

	// id

	ID dxo.CompanyID `json:"id"`

	//  base

	Base

	// ext

	Code dxo.CompanyCode `json:"code"` //  统一社会信用代码 (pure)

	Name dxo.CompanyName `json:"name"` // 企业名称

	Representative string `json:"representative"` //  法定代表人

	CompanyType string `json:"company_type"` // 类型

	Address string `json:"address"` // 住所

	FoundedAt lang.Time `json:"founded_at"` // 成立日期

	ApprovedAt lang.Time `json:"approved_at"` // 核准日期

	Capital string `json:"capital"` // 注册资本

	Registry string `json:"registry"` // 登记机关

	State string `json:"state"` // 登记状态

	OperationCategory string `json:"operation_category"` // 经营类别

	OperationRange string `json:"operation_range"` // 经营范围

	Reference dxo.URL `json:"reference"` // 参考网址

	URI dxo.URI `json:"uri"`

	Web dxo.CompanyURL `json:"web"` // 公司官网地址

	Remarks string `json:"remarks"` // 备注

}
