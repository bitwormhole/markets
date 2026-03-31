package itrademark

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/trademarks"
)

type TrademarkServiceImpl struct {

	//starter:component

	_as func(trademarks.Service) //starter:as("#")

	Dao trademarks.DAO //starter:inject("#")

}

// Find implements trademarks.Service.
func (inst *TrademarkServiceImpl) Find(ctx context.Context, id trademarks.ID) (*trademarks.DTO, error) {
	panic("unimplemented")
}

// Insert implements trademarks.Service.
func (inst *TrademarkServiceImpl) Insert(ctx context.Context, o1 *trademarks.DTO) (*trademarks.DTO, error) {

	o2 := new(trademarks.Entity)
	o4 := new(trademarks.DTO)

	err := trademarks.ConvertD2E(o1, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	err = trademarks.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil

}

// Query implements trademarks.Service.
func (inst *TrademarkServiceImpl) Query(ctx context.Context, q *trademarks.Query) ([]*trademarks.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)
	list2 := make([]*trademarks.DTO, 0)

	if err != nil {
		return nil, err
	}

	return trademarks.ConvertListE2D(list1, list2)

}

// Remove implements trademarks.Service.
func (inst *TrademarkServiceImpl) Remove(ctx context.Context, id trademarks.ID) error {
	panic("unimplemented")
}

// Update implements trademarks.Service.
func (inst *TrademarkServiceImpl) Update(ctx context.Context, id trademarks.ID, item *trademarks.DTO) (*trademarks.DTO, error) {
	panic("unimplemented")
}

func (inst *TrademarkServiceImpl) _impl() trademarks.Service {
	return inst
}
