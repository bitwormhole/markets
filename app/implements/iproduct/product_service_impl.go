package iproduct

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/products"
	"github.com/bitwormhole/markets/app/classes/utils"
)

type ProductServiceImpl struct {

	//starter:component

	_as func(products.Service) //starter:as("#")

	Dao products.DAO //starter:inject("#")
}

// Find implements products.Service.
func (inst *ProductServiceImpl) Find(ctx context.Context, id products.ID) (*products.DTO, error) {
	panic("unimplemented")
}

// Insert implements products.Service.
func (inst *ProductServiceImpl) Insert(ctx context.Context, o1 *products.DTO) (*products.DTO, error) {

	// subject

	subHolder := utils.NewSubjectHolder(ctx).UseChecker().UseGetter()
	uid := subHolder.UID()
	checker := subHolder.Checker()
	o1.Owner = uid
	o1.Creator = uid
	o1.Updater = uid
	o1.URI = products.ComputeUri(o1)

	o2 := new(products.Entity)
	o4 := new(products.DTO)

	err := products.ConvertD2E(o1, o2)
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

	err = products.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil
}

// Query implements products.Service.
func (inst *ProductServiceImpl) Query(ctx context.Context, q *products.Query) ([]*products.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	list2 := make([]*products.DTO, 0)

	if err != nil {
		return nil, err
	}

	return products.ConvertListE2D(list1, list2)

}

// Remove implements products.Service.
func (inst *ProductServiceImpl) Remove(ctx context.Context, id products.ID) error {
	panic("unimplemented")
}

// Update implements products.Service.
func (inst *ProductServiceImpl) Update(ctx context.Context, id products.ID, item *products.DTO) (*products.DTO, error) {
	panic("unimplemented")
}

func (inst *ProductServiceImpl) _impl() products.Service {
	return inst
}
