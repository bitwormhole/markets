package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Products struct {

	// base
	Base

	Items []*dto.Product `json:"products"`
}
