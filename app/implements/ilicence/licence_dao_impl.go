package ilicence

import (
	"github.com/bitwormhole/markets/app/classes/licences"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type LicenceDaoImpl struct {

	//starter:component

	_as func(licences.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *LicenceDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("licences.Entity")
	return b.Generate()
}

func (inst *LicenceDaoImpl) innerGetModel() *licences.Entity {

	return new(licences.Entity)
}

func (inst *LicenceDaoImpl) innerGetModelList() []*licences.Entity {

	return make([]*licences.Entity, 0)
}

// Find implements licences.DAO.
func (inst *LicenceDaoImpl) Find(db *gorm.DB, id licences.ID) (*licences.Entity, error) {
	panic("unimplemented")
}

// Insert implements licences.DAO.
func (inst *LicenceDaoImpl) Insert(db *gorm.DB, o1 *licences.Entity) (*licences.Entity, error) {

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

// Query implements licences.DAO.
func (inst *LicenceDaoImpl) Query(db *gorm.DB, q *licences.Query) ([]*licences.Entity, error) {

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

// Remove implements licences.DAO.
func (inst *LicenceDaoImpl) Remove(db *gorm.DB, id licences.ID) error {
	panic("unimplemented")
}

// Update implements licences.DAO.
func (inst *LicenceDaoImpl) Update(db *gorm.DB, id licences.ID, callback func(item *licences.Entity) error) (*licences.Entity, error) {
	panic("unimplemented")
}

func (inst *LicenceDaoImpl) _impl() licences.DAO {
	return inst
}
