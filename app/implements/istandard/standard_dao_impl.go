package istandard

import (
	"github.com/bitwormhole/markets/app/classes/standards"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type StandardDaoImpl struct {

	//starter:component

	_as func(standards.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *StandardDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("standards.Entity")
	return b.Generate()
}

func (inst *StandardDaoImpl) innerGetModel() *standards.Entity {

	return new(standards.Entity)
}

func (inst *StandardDaoImpl) innerGetModelList() []*standards.Entity {

	return make([]*standards.Entity, 0)
}

// Find implements standards.DAO.
func (inst *StandardDaoImpl) Find(db *gorm.DB, id standards.ID) (*standards.Entity, error) {
	panic("unimplemented")
}

// Insert implements standards.DAO.
func (inst *StandardDaoImpl) Insert(db *gorm.DB, o1 *standards.Entity) (*standards.Entity, error) {

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

// Query implements standards.DAO.
func (inst *StandardDaoImpl) Query(db *gorm.DB, q *standards.Query) ([]*standards.Entity, error) {

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

// Remove implements standards.DAO.
func (inst *StandardDaoImpl) Remove(db *gorm.DB, id standards.ID) error {
	panic("unimplemented")
}

// Update implements standards.DAO.
func (inst *StandardDaoImpl) Update(db *gorm.DB, id standards.ID, callback func(item *standards.Entity) error) (*standards.Entity, error) {
	panic("unimplemented")
}

func (inst *StandardDaoImpl) _impl() standards.DAO {
	return inst
}
