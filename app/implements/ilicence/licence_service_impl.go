package ilicence

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/licences"
	"github.com/bitwormhole/markets/app/classes/utils"
)

type LicenceServiceImpl struct {

	//starter:component

	_as func(licences.Service) //starter:as("#")

	Dao licences.DAO //starter:inject("#")

}

// Find implements licences.Service.
func (inst *LicenceServiceImpl) Find(ctx context.Context, id licences.ID) (*licences.DTO, error) {
	panic("unimplemented")
}

// Insert implements licences.Service.
func (inst *LicenceServiceImpl) Insert(ctx context.Context, o1 *licences.DTO) (*licences.DTO, error) {

	// subject

	subHolder := utils.NewSubjectHolder(ctx).UseChecker().UseGetter()
	uid := subHolder.UID()
	checker := subHolder.Checker()
	o1.Owner = uid
	o1.Creator = uid
	o1.Updater = uid
	o1.URI = licences.ComputeUri(o1)

	o2 := new(licences.Entity)
	o4 := new(licences.DTO)

	err := licences.ConvertD2E(o1, o2)
	if err != nil {
		return nil, err
	}

	// check

	err = checker.Check()
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	err = licences.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil

}

// Query implements licences.Service.
func (inst *LicenceServiceImpl) Query(ctx context.Context, q *licences.Query) ([]*licences.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	list2 := make([]*licences.DTO, 0)

	if err != nil {
		return nil, err
	}

	return licences.ConvertListE2D(list1, list2)

}

// Remove implements licences.Service.
func (inst *LicenceServiceImpl) Remove(ctx context.Context, id licences.ID) error {
	panic("unimplemented")
}

// Update implements licences.Service.
func (inst *LicenceServiceImpl) Update(ctx context.Context, id licences.ID, item *licences.DTO) (*licences.DTO, error) {
	panic("unimplemented")
}

func (inst *LicenceServiceImpl) _impl() licences.Service {
	return inst
}
