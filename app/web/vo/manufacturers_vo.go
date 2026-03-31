package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Manufacturers struct {

	// base
	Base

	Items []*dto.Manufacturer `json:"manufacturers"`
}
