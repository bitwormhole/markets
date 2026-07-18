package imedia

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/media/mlinks"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects"
)

type MediaLinkService struct {

	//starter:component

	_as func(mlinks.Service) //starter:as("#")

	Dao mlinks.DAO //starter:inject("#")

}

// Find implements [mlinks.Service].
func (inst *MediaLinkService) Find(cc context.Context, id mlinks.ID) (*mlinks.DTO, error) {

	it1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	it2 := new(mlinks.DTO)
	err = mlinks.ConvertE2D(it1, it2)
	return it2, err
}

// Insert implements [mlinks.Service].
func (inst *MediaLinkService) Insert(cc context.Context, item *mlinks.DTO) (*mlinks.DTO, error) {

	it2 := new(mlinks.Entity)
	it4 := new(mlinks.DTO)
	err := mlinks.ConvertD2E(item, it2)
	if err != nil {
		return nil, err
	}

	uid := inst.innerGetCurrentUserID(cc)
	it2.Owner = uid
	it2.Creator = uid
	it2.Updater = uid

	it3, err := inst.Dao.Insert(nil, it2)
	if err != nil {
		return nil, err
	}

	err = mlinks.ConvertE2D(it3, it4)
	return it4, err
}

func (inst *MediaLinkService) innerGetCurrentUserID(c context.Context) rbac.UserID {
	sub, err := subjects.GetCurrent(c)
	if err != nil {
		return 0
	}
	gett, err := sub.DoGet()
	if err != nil {
		return 0
	}
	return gett.GetUserID()
}

func (inst *MediaLinkService) _impl() mlinks.Service {
	return inst
}
