package imanufacturer

import (
	"github.com/bitwormhole/markets/app/classes/manufacturers"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type ManufacturerDaoImpl struct {

	//starter:component

	_as func(manufacturers.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *ManufacturerDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("manufacturers.Entity")
	return b.Generate()
}

func (inst *ManufacturerDaoImpl) innerGetModel() *manufacturers.Entity {

	return new(manufacturers.Entity)
}

func (inst *ManufacturerDaoImpl) innerGetModelList() []*manufacturers.Entity {

	return make([]*manufacturers.Entity, 0)
}

// Find implements manufacturers.DAO.
func (inst *ManufacturerDaoImpl) Find(db *gorm.DB, id manufacturers.ID) (*manufacturers.Entity, error) {
	panic("unimplemented")
}

// Insert implements manufacturers.DAO.
func (inst *ManufacturerDaoImpl) Insert(db *gorm.DB, o1 *manufacturers.Entity) (*manufacturers.Entity, error) {

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

// Query implements manufacturers.DAO.
func (inst *ManufacturerDaoImpl) Query(db *gorm.DB, q *manufacturers.Query) ([]*manufacturers.Entity, error) {

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

// Remove implements manufacturers.DAO.
func (inst *ManufacturerDaoImpl) Remove(db *gorm.DB, id manufacturers.ID) error {
	panic("unimplemented")
}

// Update implements manufacturers.DAO.
func (inst *ManufacturerDaoImpl) Update(db *gorm.DB, id manufacturers.ID, callback func(item *manufacturers.Entity) error) (*manufacturers.Entity, error) {
	panic("unimplemented")
}

func (inst *ManufacturerDaoImpl) _impl() manufacturers.DAO {
	return inst
}
