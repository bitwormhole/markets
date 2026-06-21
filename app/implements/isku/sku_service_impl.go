package isku

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/skus"
	"github.com/bitwormhole/markets/app/classes/utils"
)

type SkuServiceImpl struct {

	//starter:component

	_as func(skus.Service) //starter:as("#")

	Dao skus.DAO //starter:inject("#")

}

// Find implements skus.Service.
func (inst *SkuServiceImpl) Find(ctx context.Context, id skus.ID) (*skus.DTO, error) {
	panic("unimplemented")
}

// Insert implements skus.Service.
func (inst *SkuServiceImpl) Insert(ctx context.Context, o1 *skus.DTO) (*skus.DTO, error) {

	// subject

	subHolder := utils.NewSubjectHolder(ctx).UseChecker().UseGetter()
	uid := subHolder.UID()
	checker := subHolder.Checker()
	o1.Owner = uid
	o1.Creator = uid
	o1.Updater = uid
	o1.URI = skus.ComputeUri(o1)

	o2 := new(skus.Entity)
	o4 := new(skus.DTO)

	err := skus.ConvertD2E(o1, o2)
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

	err = skus.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil

}

// Query implements skus.Service.
func (inst *SkuServiceImpl) Query(ctx context.Context, q *skus.Query) ([]*skus.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	list2 := make([]*skus.DTO, 0)

	if err != nil {
		return nil, err
	}

	return skus.ConvertListE2D(list1, list2)

}

// Remove implements skus.Service.
func (inst *SkuServiceImpl) Remove(ctx context.Context, id skus.ID) error {
	panic("unimplemented")
}

// Update implements skus.Service.
func (inst *SkuServiceImpl) Update(ctx context.Context, id skus.ID, item *skus.DTO) (*skus.DTO, error) {
	panic("unimplemented")
}

func (inst *SkuServiceImpl) _impl() skus.Service {
	return inst
}
