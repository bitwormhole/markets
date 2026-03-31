package vo

import "github.com/bitwormhole/markets/app/web/dto"

type SKUs struct {

	//  base
	Base

	Items []*dto.SKU `json:"skus"`
}
