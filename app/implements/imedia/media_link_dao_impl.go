package imedia

import (
	"github.com/bitwormhole/markets/app/classes/media/mlinks"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type MediaLinkDao struct {

	//starter:component

	_as func(mlinks.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *MediaLinkDao) innerMakeItem() *mlinks.Entity {
	return new(mlinks.Entity)
}

func (inst *MediaLinkDao) innerMakeItemList() []*mlinks.Entity {
	return make([]*mlinks.Entity, 0)
}

func (inst *MediaLinkDao) innerPrepareDB(db *gorm.DB) *gorm.DB {
	return inst.Agent.DB(db)
}

func (inst *MediaLinkDao) innerGenUUID() lang.UUID {
	b := inst.UUIDGenSer.Build()
	b.Class("mlinks.Entity")
	return b.Generate()
}

// Find implements [mlinks.DAO].
func (inst *MediaLinkDao) Find(db *gorm.DB, id mlinks.ID) (*mlinks.Entity, error) {
	db = inst.innerPrepareDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// Insert implements [mlinks.DAO].
func (inst *MediaLinkDao) Insert(db *gorm.DB, item *mlinks.Entity) (*mlinks.Entity, error) {
	item.ID = 0
	item.UUID = inst.innerGenUUID()
	db = inst.innerPrepareDB(db)
	res := db.Create(item)
	err := res.Error
	return item, err
}

func (inst *MediaLinkDao) _impl() mlinks.DAO {
	return inst
}
