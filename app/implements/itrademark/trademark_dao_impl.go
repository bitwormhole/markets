package itrademark

import (
	"github.com/bitwormhole/markets/app/classes/trademarks"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type TrademarkDaoImpl struct {

	//starter:component

	_as func(trademarks.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *TrademarkDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("trademarks.Entity")
	return b.Generate()
}

func (inst *TrademarkDaoImpl) innerGetModel() *trademarks.Entity {

	return new(trademarks.Entity)
}

func (inst *TrademarkDaoImpl) innerGetModelList() []*trademarks.Entity {

	return make([]*trademarks.Entity, 0)
}

// Find implements Trademarks.DAO.
func (inst *TrademarkDaoImpl) Find(db *gorm.DB, id trademarks.ID) (*trademarks.Entity, error) {
	panic("unimplemented")
}

// Insert implements Trademarks.DAO.
func (inst *TrademarkDaoImpl) Insert(db *gorm.DB, o1 *trademarks.Entity) (*trademarks.Entity, error) {

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

// Query implements Trademarks.DAO.
func (inst *TrademarkDaoImpl) Query(db *gorm.DB, q *trademarks.Query) ([]*trademarks.Entity, error) {

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

// Remove implements Trademarks.DAO.
func (inst *TrademarkDaoImpl) Remove(db *gorm.DB, id trademarks.ID) error {
	panic("unimplemented")
}

// Update implements Trademarks.DAO.
func (inst *TrademarkDaoImpl) Update(db *gorm.DB, id trademarks.ID, callback func(item *trademarks.Entity) error) (*trademarks.Entity, error) {
	panic("unimplemented")
}

func (inst *TrademarkDaoImpl) _impl() trademarks.DAO {
	return inst
}
