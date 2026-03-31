package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Shops struct {

	// base
	Base

	Items []*dto.Shop `json:"shops"`
}
