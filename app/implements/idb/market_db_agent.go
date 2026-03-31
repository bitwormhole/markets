package idb

import (
	"fmt"

	"github.com/bitwormhole/markets/app/data/marketdb"
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

type MarketDBAgent struct {

	//starter:component

	_as func(marketdb.Agent) //starter:as("#")

	DSM libgorm.DataSourceManager //starter:inject("#")

	Alias string //starter:inject("${datagroup.markets.datasource}")

	cached *gorm.DB
}

func (inst *MarketDBAgent) innerGetDB() (*gorm.DB, error) {

	db := inst.cached
	if db != nil {
		return db, nil
	}

	db, err := inst.innerLoadDB()
	if err != nil {
		return nil, err
	}

	inst.cached = db
	return db, nil
}

func (inst *MarketDBAgent) innerLoadDB() (*gorm.DB, error) {

	name := inst.Alias

	ds, err := inst.DSM.GetDataSource(name)
	if err != nil {
		return nil, err
	}

	db, err := ds.DB()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (inst *MarketDBAgent) innerStart() error {

	db, err := inst.innerGetDB()

	if err != nil {
		return err
	}

	if db == nil {
		return fmt.Errorf("MarketDBAgent.innerStart: db is nil")
	}

	return nil
}

// DB implements marketdb.Agent.
func (inst *MarketDBAgent) DB(db *gorm.DB) *gorm.DB {

	if db != nil {
		return db
	}

	db, err := inst.innerGetDB()

	if err != nil {
		panic(err)
	} else if db == nil {
		panic("MarketDBAgent: db is nil")
	}

	return db
}

func (inst *MarketDBAgent) Life() *application.Life {
	l := new(application.Life)
	l.OnStart = inst.innerStart
	return l
}

func (inst *MarketDBAgent) _impl() marketdb.Agent {
	return inst
}
