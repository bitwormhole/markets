package imedia

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/media/mobjects"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects"
)

type MediaObjectService struct {

	//starter:component

	_as func(mobjects.Service) //starter:as("#")

	Dao mobjects.DAO //starter:inject("#")

}

// ContainsSum implements [mobjects.Service].
func (inst *MediaObjectService) ContainsSum(cc context.Context, sum []byte) (bool, error) {
	return inst.Dao.ContainsSum(nil, sum)
}

// FindBySum implements [mobjects.Service].
func (inst *MediaObjectService) FindBySum(cc context.Context, sum []byte) (*mobjects.DTO, error) {

	it1, err := inst.Dao.FindBySum(nil, sum)
	if err != nil {
		return nil, err
	}

	it2 := new(mobjects.DTO)
	err = mobjects.ConvertE2D(it1, it2)
	return it2, err
}

// Find implements [mobjects.Service].
func (inst *MediaObjectService) Find(cc context.Context, id mobjects.ID) (*mobjects.DTO, error) {

	it1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	it2 := new(mobjects.DTO)
	err = mobjects.ConvertE2D(it1, it2)
	return it2, err
}

// Insert implements [mobjects.Service].
func (inst *MediaObjectService) Insert(cc context.Context, item *mobjects.DTO) (*mobjects.DTO, error) {

	it2 := new(mobjects.Entity)
	it4 := new(mobjects.DTO)
	err := mobjects.ConvertD2E(item, it2)
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

	err = mobjects.ConvertE2D(it3, it4)
	return it4, err
}

func (inst *MediaObjectService) innerGetCurrentUserID(c context.Context) rbac.UserID {
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

func (inst *MediaObjectService) _impl() mobjects.Service {
	return inst
}
