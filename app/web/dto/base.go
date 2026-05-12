package dto

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/rbac"
)

type Base struct {
	rbac.BaseDTO

	Domain dxo.DomainName `json:"domain"`
}
