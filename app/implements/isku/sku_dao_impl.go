package isku

import (
	"github.com/bitwormhole/markets/app/classes/skus"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type SkuDaoImpl struct {

	//starter:component

	_as func(skus.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *SkuDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("skus.Entity")
	return b.Generate()
}

func (inst *SkuDaoImpl) innerGetModel() *skus.Entity {

	return new(skus.Entity)
}

func (inst *SkuDaoImpl) innerGetModelList() []*skus.Entity {

	return make([]*skus.Entity, 0)
}

// Find implements skus.DAO.
func (inst *SkuDaoImpl) Find(db *gorm.DB, id skus.ID) (*skus.Entity, error) {
	panic("unimplemented")
}

// Insert implements skus.DAO.
func (inst *SkuDaoImpl) Insert(db *gorm.DB, o1 *skus.Entity) (*skus.Entity, error) {

	db = inst.Agent.DB(db)
	uuid := inst.innerGenUUID()

	o1.ID = 0
	o1.UUID = uuid

	res := db.Create(o1)
	err := res.Error
	if err != nil {
		return nil, err
	}

	return o1, nil

}

// Query implements skus.DAO.
func (inst *SkuDaoImpl) Query(db *gorm.DB, q *skus.Query) ([]*skus.Entity, error) {

	mod := inst.innerGetModel()
	list := inst.innerGetModelList()
	f := new(entities.Finder)

	db = inst.Agent.DB(db)

	f.SetDB(db)
	f.SetAll(q.All)
	f.SetWant(q.Want)
	f.SetPage(&q.Pagination)
	f.SetDest(&list)
	f.SetModel(mod)

	err := f.Find()

	return list, err

}

// Remove implements skus.DAO.
func (inst *SkuDaoImpl) Remove(db *gorm.DB, id skus.ID) error {
	panic("unimplemented")
}

// Update implements skus.DAO.
func (inst *SkuDaoImpl) Update(db *gorm.DB, id skus.ID, callback func(item *skus.Entity) error) (*skus.Entity, error) {
	panic("unimplemented")
}

func (inst *SkuDaoImpl) _impl() skus.DAO {
	return inst
}
