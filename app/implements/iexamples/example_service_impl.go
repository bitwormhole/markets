package iexamples

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/examples"
)

type ExampleServiceImpl struct {

	//starter:component

	_as func(examples.Service) //starter:as("#")

	Dao examples.DAO //starter:inject("#")

}

// Find implements examples.Service.
func (inst *ExampleServiceImpl) Find(ctx context.Context, id examples.ID) (*examples.DTO, error) {
	panic("unimplemented")
}

// Insert implements examples.Service.
func (inst *ExampleServiceImpl) Insert(ctx context.Context, o1 *examples.DTO) (*examples.DTO, error) {

	o2 := new(examples.Entity)
	o4 := new(examples.DTO)

	err := examples.ConvertD2E(o1, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	err = examples.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil

}

// Query implements examples.Service.
func (inst *ExampleServiceImpl) Query(ctx context.Context, q *examples.Query) ([]*examples.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	list2 := make([]*examples.DTO, 0)

	if err != nil {
		return nil, err
	}

	return examples.ConvertListE2D(list1, list2)

}

// Remove implements examples.Service.
func (inst *ExampleServiceImpl) Remove(ctx context.Context, id examples.ID) error {
	panic("unimplemented")
}

// Update implements examples.Service.
func (inst *ExampleServiceImpl) Update(ctx context.Context, id examples.ID, item *examples.DTO) (*examples.DTO, error) {
	panic("unimplemented")
}

func (inst *ExampleServiceImpl) _impl() examples.Service {
	return inst
}
