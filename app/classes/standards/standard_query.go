package standards

import (
	"github.com/starter-go/rbac"
)

type Query struct {
	All bool

	Pagination rbac.Pagination

	Want *Entity
}
