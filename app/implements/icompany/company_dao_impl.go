package icompany

import (
	"github.com/bitwormhole/markets/app/classes/companies"
	"github.com/bitwormhole/markets/app/data/entities"
	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type CompanyDaoImpl struct {

	//starter:component

	_as func(companies.DAO) //starter:as("#")

	Agent      marketdb.Agent     //starter:inject("#")
	UUIDGenSer random.UUIDService //starter:inject("#")

}

func (inst *CompanyDaoImpl) innerGenUUID() lang.UUID {

	b := inst.UUIDGenSer.Build()
	b.Class("companies.Entity")
	return b.Generate()
}

// Find implements companies.DAO.
func (inst *CompanyDaoImpl) Find(db *gorm.DB, id companies.ID) (*companies.Entity, error) {
	panic("unimplemented")
}

// Insert implements companies.DAO.
func (inst *CompanyDaoImpl) Insert(db *gorm.DB, o1 *companies.Entity) (*companies.Entity, error) {

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

func (inst *CompanyDaoImpl) innerGetModel() *companies.Entity {

	return new(entities.Company)
}

func (inst *CompanyDaoImpl) innerGetModelList() []*companies.Entity {

	return make([]*companies.Entity, 0)
}

// Query implements companies.DAO.
func (inst *CompanyDaoImpl) Query(db *gorm.DB, q *companies.Query) ([]*companies.Entity, error) {

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

// Remove implements companies.DAO.
func (inst *CompanyDaoImpl) Remove(db *gorm.DB, id companies.ID) error {
	panic("unimplemented")
}

// Update implements companies.DAO.
func (inst *CompanyDaoImpl) Update(db *gorm.DB, id companies.ID, callback func(item *companies.Entity) error) (*companies.Entity, error) {
	panic("unimplemented")
}

func (inst *CompanyDaoImpl) _impl() companies.DAO {
	return inst
}
