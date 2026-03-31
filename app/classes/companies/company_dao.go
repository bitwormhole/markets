package companies

import "gorm.io/gorm"

type DAO interface {

	// query

	Find(db *gorm.DB, id ID) (*Entity, error)

	Query(db *gorm.DB, q *Query) ([]*Entity, error)

	// edit

	Insert(db *gorm.DB, item *Entity) (*Entity, error)

	Update(db *gorm.DB, id ID, callback func(item *Entity) error) (*Entity, error)

	Remove(db *gorm.DB, id ID) error
}
