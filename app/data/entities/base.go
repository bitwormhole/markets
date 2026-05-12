package entities

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/security-gorm/rbacdb"
)

type Base struct {
	rbacdb.BaseEntity

	Domain dxo.DomainName
}
