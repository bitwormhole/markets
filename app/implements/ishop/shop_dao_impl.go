package ishop

import (
	"github.com/bitwormhole/markets/app/classes/shops"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type ShopDaoImpl struct {

	//starter:component

	_as func(shops.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *ShopDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("shops.Entity")
	return b.Generate()
}

func (inst *ShopDaoImpl) innerGetModel() *shops.Entity {

	return new(shops.Entity)
}

func (inst *ShopDaoImpl) innerGetModelList() []*shops.Entity {

	return make([]*shops.Entity, 0)
}

// Find implements shops.DAO.
func (inst *ShopDaoImpl) Find(db *gorm.DB, id shops.ID) (*shops.Entity, error) {

	panic("unimplemented")
}

// Insert implements shops.DAO.
func (inst *ShopDaoImpl) Insert(db *gorm.DB, o1 *shops.Entity) (*shops.Entity, error) {

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

// Query implements shops.DAO.
func (inst *ShopDaoImpl) Query(db *gorm.DB, q *shops.Query) ([]*shops.Entity, error) {

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

// Remove implements shops.DAO.
func (inst *ShopDaoImpl) Remove(db *gorm.DB, id shops.ID) error {
	panic("unimplemented")
}

// Update implements shops.DAO.
func (inst *ShopDaoImpl) Update(db *gorm.DB, id shops.ID, callback func(item *shops.Entity) error) (*shops.Entity, error) {
	panic("unimplemented")
}

func (inst *ShopDaoImpl) _impl() shops.DAO {
	return inst
}
