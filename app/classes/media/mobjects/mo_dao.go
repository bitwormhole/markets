package mobjects

import (
	"gorm.io/gorm"
)

type DAO interface {
	Find(db *gorm.DB, id ID) (*Entity, error)

	FindBySum(db *gorm.DB, sum []byte) (*Entity, error)

	ContainsSum(db *gorm.DB, sum []byte) (bool, error)

	Insert(db *gorm.DB, item *Entity) (*Entity, error)
}
