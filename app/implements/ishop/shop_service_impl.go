package ishop

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/shops"
)

type ShopServiceImpl struct {

	//starter:component

	_as func(shops.Service) //starter:as("#")

	Dao shops.DAO //starter:inject("#")

}

// Find implements shops.Service.
func (inst *ShopServiceImpl) Find(ctx context.Context, id shops.ID) (*shops.DTO, error) {
	panic("unimplemented")
}

// Insert implements shops.Service.
func (inst *ShopServiceImpl) Insert(ctx context.Context, o1 *shops.DTO) (*shops.DTO, error) {

	o2 := new(shops.Entity)
	o4 := new(shops.DTO)

	err := shops.ConvertD2E(o1, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	err = shops.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil

}

// Query implements shops.Service.
func (inst *ShopServiceImpl) Query(ctx context.Context, q *shops.Query) ([]*shops.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	list2 := make([]*shops.DTO, 0)

	if err != nil {
		return nil, err
	}

	return shops.ConvertListE2D(list1, list2)

}

// Remove implements shops.Service.
func (inst *ShopServiceImpl) Remove(ctx context.Context, id shops.ID) error {
	panic("unimplemented")
}

// Update implements shops.Service.
func (inst *ShopServiceImpl) Update(ctx context.Context, id shops.ID, item *shops.DTO) (*shops.DTO, error) {
	panic("unimplemented")
}

func (inst *ShopServiceImpl) _impl() shops.Service {
	return inst
}
