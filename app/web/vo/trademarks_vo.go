package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Trademarks struct {

	// base
	Base

	Items []*dto.Trademark `json:"trademarks"`
}
