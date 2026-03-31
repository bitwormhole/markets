package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Examples struct {

	// base
	Base

	Items []*dto.Example `json:"examples"`
}
