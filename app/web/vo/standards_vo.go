package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Standards struct {

	// base
	Base

	Items []*dto.Standard `json:"standards"`
}
