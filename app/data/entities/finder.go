package entities

import (
	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

////////////////////////////////////////////////////////////////////////////////

type Finder struct {
	db       *gorm.DB
	item     any
	itemList any
	want     any
	page     *rbac.Pagination
	all      bool
	count    int64
}

func (inst *Finder) SetDB(db *gorm.DB) *Finder {

	inst.db = db
	return inst
}

func (inst *Finder) SetDest(d any) *Finder {

	inst.itemList = d
	return inst
}

func (inst *Finder) SetModel(m any) *Finder {

	inst.item = m
	return inst
}

func (inst *Finder) SetWant(w any) *Finder {

	inst.want = w
	return inst
}

func (inst *Finder) SetAll(all bool) *Finder {

	inst.all = all
	return inst
}

func (inst *Finder) SetPage(p *rbac.Pagination) *Finder {

	inst.page = p
	return inst
}

func (inst *Finder) Find() error {

	db := inst.db
	dest := inst.itemList
	item := inst.item
	page := inst.page
	want := inst.want
	all := inst.all

	// model

	if item != nil {
		db = db.Model(item)
	}

	// want

	if want != nil {
		db = db.Where(want)
	}

	// page

	if page == nil {
		if all {
			// NOP
		} else {
			page = &rbac.Pagination{
				Page: 1,
				Size: 5,
			}
		}
	}

	if page != nil {
		var count int64
		db.Count(&count)
		db = db.Limit(int(page.Limit())).Offset(int(page.Offset()))
		page.Total = count
		inst.count = count
	}

	// find

	res := db.Find(dest)

	return res.Error
}

////////////////////////////////////////////////////////////////////////////////
