package vo

import "github.com/bitwormhole/markets/app/web/dto"

type Licences struct {

	// base
	Base

	Items []*dto.Licence `json:"licences"`
}
