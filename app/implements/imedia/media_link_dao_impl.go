package imedia

import (
	"github.com/bitwormhole/markets/app/classes/media/mlinks"
	"gorm.io/gorm"
)

type MediaLinkDao struct {

	//starter:component

	_as func(mlinks.DAO) //starter:as("#")

}

// Find implements [mlinks.DAO].
func (inst *MediaLinkDao) Find(db *gorm.DB, id mlinks.ID) (*mlinks.Entity, error) {
	panic("unimplemented")
}

// Insert implements [mlinks.DAO].
func (inst *MediaLinkDao) Insert(db *gorm.DB, item *mlinks.Entity) (*mlinks.Entity, error) {
	panic("unimplemented")
}

func (inst *MediaLinkDao) _impl() mlinks.DAO {
	return inst
}
