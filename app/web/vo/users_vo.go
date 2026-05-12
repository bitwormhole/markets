package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Users struct {

	// base
	Base

	Items []*dto.User `json:"users"`
}
