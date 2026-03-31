package istandard

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/standards"
)

type StandardServiceImpl struct {

	//starter:component

	_as func(standards.Service) //starter:as("#")

	Dao standards.DAO //starter:inject("#")

}

// Find implements standards.Service.
func (inst *StandardServiceImpl) Find(ctx context.Context, id standards.ID) (*standards.DTO, error) {
	panic("unimplemented")
}

// Insert implements standards.Service.
func (inst *StandardServiceImpl) Insert(ctx context.Context, o1 *standards.DTO) (*standards.DTO, error) {

	o2 := new(standards.Entity)
	o4 := new(standards.DTO)

	err := standards.ConvertD2E(o1, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	err = standards.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil

}

// Query implements standards.Service.
func (inst *StandardServiceImpl) Query(ctx context.Context, q *standards.Query) ([]*standards.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	list2 := make([]*standards.DTO, 0)

	if err != nil {
		return nil, err
	}

	return standards.ConvertListE2D(list1, list2)

}

// Remove implements standards.Service.
func (inst *StandardServiceImpl) Remove(ctx context.Context, id standards.ID) error {
	panic("unimplemented")
}

// Update implements standards.Service.
func (inst *StandardServiceImpl) Update(ctx context.Context, id standards.ID, item *standards.DTO) (*standards.DTO, error) {
	panic("unimplemented")
}

func (inst *StandardServiceImpl) _impl() standards.Service {
	return inst
}
