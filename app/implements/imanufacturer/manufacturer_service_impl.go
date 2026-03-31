package imanufacturer

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/manufacturers"
)

type ManufacturerServiceImpl struct {

	//starter:component

	_as func(manufacturers.Service) //starter:as("#")

	Dao manufacturers.DAO //starter:inject("#")

}

// Find implements manufacturers.Service.
func (inst *ManufacturerServiceImpl) Find(ctx context.Context, id manufacturers.ID) (*manufacturers.DTO, error) {
	panic("unimplemented")
}

// Insert implements manufacturers.Service.
func (inst *ManufacturerServiceImpl) Insert(ctx context.Context, o1 *manufacturers.DTO) (*manufacturers.DTO, error) {

	o2 := new(manufacturers.Entity)
	o4 := new(manufacturers.DTO)

	err := manufacturers.ConvertD2E(o1, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	err = manufacturers.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil

}

// Query implements manufacturers.Service.
func (inst *ManufacturerServiceImpl) Query(ctx context.Context, q *manufacturers.Query) ([]*manufacturers.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	list2 := make([]*manufacturers.DTO, 0)

	if err != nil {
		return nil, err
	}

	return manufacturers.ConvertListE2D(list1, list2)

}

// Remove implements manufacturers.Service.
func (inst *ManufacturerServiceImpl) Remove(ctx context.Context, id manufacturers.ID) error {
	panic("unimplemented")
}

// Update implements manufacturers.Service.
func (inst *ManufacturerServiceImpl) Update(ctx context.Context, id manufacturers.ID, item *manufacturers.DTO) (*manufacturers.DTO, error) {
	panic("unimplemented")
}

func (inst *ManufacturerServiceImpl) _impl() manufacturers.Service {
	return inst
}
