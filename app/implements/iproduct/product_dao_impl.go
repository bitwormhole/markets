package iproduct

import (
	"github.com/bitwormhole/markets/app/classes/products"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type ProductDaoImpl struct {

	//starter:component

	_as func(products.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *ProductDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("products.Entity")
	return b.Generate()
}

func (inst *ProductDaoImpl) innerGetModel() *products.Entity {

	return new(products.Entity)
}

func (inst *ProductDaoImpl) innerGetModelList() []*products.Entity {

	return make([]*products.Entity, 0)
}

// Find implements products.DAO.
func (inst *ProductDaoImpl) Find(db *gorm.DB, id products.ID) (*products.Entity, error) {
	panic("unimplemented")
}

// Insert implements products.DAO.
func (inst *ProductDaoImpl) Insert(db *gorm.DB, o1 *products.Entity) (*products.Entity, error) {

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

// Query implements products.DAO.
func (inst *ProductDaoImpl) Query(db *gorm.DB, q *products.Query) ([]*products.Entity, error) {

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

// Remove implements products.DAO.
func (inst *ProductDaoImpl) Remove(db *gorm.DB, id products.ID) error {
	panic("unimplemented")
}

// Update implements products.DAO.
func (inst *ProductDaoImpl) Update(db *gorm.DB, id products.ID, callback func(item *products.Entity) error) (*products.Entity, error) {
	panic("unimplemented")
}

func (inst *ProductDaoImpl) _impl() products.DAO {
	return inst
}
