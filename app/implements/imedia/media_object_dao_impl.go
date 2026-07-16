package imedia

import (
	"github.com/bitwormhole/markets/app/classes/media/mobjects"
	"gorm.io/gorm"
)

type MediaObjectDao struct {

	//starter:component

	_as func(mobjects.DAO) //starter:as("#")

}

// Find implements [mobjects.DAO].
func (inst *MediaObjectDao) Find(db *gorm.DB, id mobjects.ID) (*mobjects.Entity, error) {
	panic("unimplemented")
}

// Insert implements [mobjects.DAO].
func (inst *MediaObjectDao) Insert(db *gorm.DB, item *mobjects.Entity) (*mobjects.Entity, error) {
	panic("unimplemented")
}

func (inst *MediaObjectDao) _impl() mobjects.DAO {
	return inst
}
