package icompany

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/companies"
)

type CompanyServiceImpl struct {

	//starter:component

	_as func(companies.Service) //starter:as("#")

	Dao companies.DAO //starter:inject("#")

}

// Find implements companies.Service.
func (inst *CompanyServiceImpl) Find(ctx context.Context, id companies.ID) (*companies.DTO, error) {
	panic("unimplemented")
}

// Insert implements companies.Service.
func (inst *CompanyServiceImpl) Insert(ctx context.Context, o1 *companies.DTO) (*companies.DTO, error) {

	o2 := new(companies.Entity)
	o4 := new(companies.DTO)

	err := companies.ConvertD2E(o1, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	err = companies.ConvertE2D(o3, o4)
	if err != nil {
		return nil, err
	}

	return o4, nil

}

// Query implements companies.Service.
func (inst *CompanyServiceImpl) Query(ctx context.Context, q *companies.Query) ([]*companies.DTO, error) {

	list1, err := inst.Dao.Query(nil, q)

	if err != nil {
		return nil, err
	}

	list2 := make([]*companies.DTO, 0)

	return companies.ConvertListE2D(list1, list2)
}

// Remove implements companies.Service.
func (inst *CompanyServiceImpl) Remove(ctx context.Context, id companies.ID) error {
	panic("unimplemented")
}

// Update implements companies.Service.
func (inst *CompanyServiceImpl) Update(ctx context.Context, id companies.ID, item *companies.DTO) (*companies.DTO, error) {
	panic("unimplemented")
}

func (inst *CompanyServiceImpl) _impl() companies.Service {
	return inst
}
