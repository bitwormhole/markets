package iexamples

import (
	"github.com/bitwormhole/markets/app/classes/examples"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type ExampleDaoImpl struct {

	//starter:component

	_as func(examples.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *ExampleDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("examples.Entity")
	return b.Generate()
}

// Find implements examples.DAO.
func (inst *ExampleDaoImpl) Find(db *gorm.DB, id examples.ID) (*examples.Entity, error) {
	panic("unimplemented")
}

// Insert implements examples.DAO.
func (inst *ExampleDaoImpl) Insert(db *gorm.DB, o1 *examples.Entity) (*examples.Entity, error) {

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

func (inst *ExampleDaoImpl) innerGetModel() *examples.Entity {

	return new(examples.Entity)
}

func (inst *ExampleDaoImpl) innerGetModelList() []*examples.Entity {

	return make([]*examples.Entity, 0)
}

// Query implements examples.DAO.
func (inst *ExampleDaoImpl) Query(db *gorm.DB, q *examples.Query) ([]*examples.Entity, error) {

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

// Remove implements examples.DAO.
func (inst *ExampleDaoImpl) Remove(db *gorm.DB, id examples.ID) error {
	panic("unimplemented")
}

// Update implements examples.DAO.
func (inst *ExampleDaoImpl) Update(db *gorm.DB, id examples.ID, callback func(item *examples.Entity) error) (*examples.Entity, error) {
	panic("unimplemented")
}

func (inst *ExampleDaoImpl) _impl() examples.DAO {
	return inst
}
