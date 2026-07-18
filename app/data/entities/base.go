package entities

import (
	"github.com/bitwormhole/markets/app/data/dxo"
	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

type Base struct {

	// rbacdb.BaseEntity

	rbac.BaseEntity

	DeletedAt gorm.DeletedAt `gorm:"index"` // 这个字段需要在扩展结构中定义

	Domain dxo.DomainName
}

// GetTarget implements [rbac.EntityRef].
func (inst *Base) GetTarget() *rbac.Entity {
	return &inst.BaseEntity
}

func (inst *Base) _impl() rbac.EntityRef {
	return inst
}
