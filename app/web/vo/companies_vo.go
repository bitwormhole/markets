package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Companies struct {

	// base
	Base

	Items []*dto.Company `json:"companies"`
}
