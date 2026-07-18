package imedia

import (
	"fmt"

	"github.com/bitwormhole/markets/app/classes/media/mobjects"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type MediaObjectDao struct {

	//starter:component

	_as func(mobjects.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

// ContainsSum implements [mobjects.DAO].
func (inst *MediaObjectDao) ContainsSum(db *gorm.DB, sum []byte) (bool, error) {

	hash, err := inst.innerPrepareSum(sum)
	if err != nil {
		return false, err
	}

	want := inst.innerMakeItem()
	have := inst.innerMakeItem()
	count := int64(0)
	want.ContentSum = hash
	db = inst.innerPrepareDB(db)

	db = db.Model(have).Where(want).Limit(1)
	res := db.Count(&count)

	err = res.Error
	return (count > 0), err
}

func (inst *MediaObjectDao) innerMakeItem() *mobjects.Entity {
	return new(mobjects.Entity)
}

func (inst *MediaObjectDao) innerMakeItemList() []*mobjects.Entity {
	return make([]*mobjects.Entity, 0)
}

func (inst *MediaObjectDao) innerPrepareDB(db *gorm.DB) *gorm.DB {
	return inst.Agent.DB(db)
}

func (inst *MediaObjectDao) innerPrepareSum(sum []byte) (lang.Hex, error) {
	const minSumLength = 6
	if len(sum) < minSumLength {
		return "", fmt.Errorf("bad content-sum")
	}
	hex := lang.HexFromBytes(sum)
	return hex, nil
}

func (inst *MediaObjectDao) innerGenUUID() lang.UUID {
	b := inst.UUIDGenSer.Build()
	b.Class("mobjects.Entity")
	return b.Generate()
}

// Find implements [mobjects.DAO].
func (inst *MediaObjectDao) Find(db *gorm.DB, id mobjects.ID) (*mobjects.Entity, error) {
	db = inst.innerPrepareDB(db)
	item := inst.innerMakeItem()
	res := db.First(item, id)
	err := res.Error
	return item, err
}

// FindBySum implements [mobjects.DAO].
func (inst *MediaObjectDao) FindBySum(db *gorm.DB, sum []byte) (*mobjects.Entity, error) {

	hash, err := inst.innerPrepareSum(sum)
	if err != nil {
		return nil, err
	}

	db = inst.innerPrepareDB(db)
	want := new(mobjects.Entity)
	have := new(mobjects.Entity)
	want.ContentSum = hash

	res := db.Model(have).Where(want).First(have)

	err = res.Error
	return have, err
}

// Insert implements [mobjects.DAO].
func (inst *MediaObjectDao) Insert(db *gorm.DB, item *mobjects.Entity) (*mobjects.Entity, error) {
	item.ID = 0
	item.UUID = inst.innerGenUUID()
	db = inst.innerPrepareDB(db)
	res := db.Create(item)
	err := res.Error
	return item, err
}

func (inst *MediaObjectDao) _impl() mobjects.DAO {
	return inst
}
