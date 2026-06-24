package icompany

import (
	"context"

	"github.com/bitwormhole/markets/app/classes/companies"
	"github.com/bitwormhole/markets/app/classes/utils"
)

type CompanyServiceImpl struct {

	//starter:component

	_as func(companies.Service) //starter:as("#")

	Dao companies.DAO //starter:inject("#")

}

// Find implements companies.Service.
func (inst *CompanyServiceImpl) Find(ctx context.Context, id companies.ID) (*companies.DTO, error) {

	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	o2 := new(companies.DTO)
	err = companies.ConvertE2D(o1, o2)
	if err != nil {
		return nil, err
	}

	return o2, nil
}

// Insert implements companies.Service.
func (inst *CompanyServiceImpl) Insert(ctx context.Context, o1 *companies.DTO) (*companies.DTO, error) {

	// subject

	subHolder := utils.NewSubjectHolder(ctx).UseChecker().UseGetter()
	uid := subHolder.UID()
	checker := subHolder.Checker()
	o1.Owner = uid
	o1.Creator = uid
	o1.Updater = uid
	o1.URI = companies.ComputeUri(o1)

	o2 := new(companies.Entity)
	o4 := new(companies.DTO)

	err := companies.ConvertD2E(o1, o2)
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

	hSub := utils.NewSubjectHolder(ctx).UseChecker().UseGetter()
	checker := hSub.Checker()

	o1 := new(companies.Entity)
	err := companies.ConvertD2E(item, o1)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Update(nil, id, func(o2 *companies.Entity) error {

		checker.CheckObject(companies.PreCheckEntity(o2))
		err := checker.Check()
		if err != nil {
			return err
		}

		return inst.handleUpdateItem(o1, o2)
	})

	if err != nil {
		return nil, err
	}

	o4 := new(companies.DTO)
	err = companies.ConvertE2D(o3, o4)
	return o4, err
}

func (inst *CompanyServiceImpl) handleUpdateItem(src, dst *companies.Entity) error {

	// address

	if src.Address != "" {
		dst.Address = src.Address
	}

	// app_at

	if !src.ApprovedAt.IsZero() {
		dst.ApprovedAt = src.ApprovedAt
	}

	// capital

	if src.Capital != "" {
		dst.Capital = src.Capital
	}

	// code

	if src.Code != "" {
		dst.Code = src.Code
	}

	// company_type

	if src.CompanyType != "" {
		dst.CompanyType = src.CompanyType
	}

	// domain

	if src.Domain != "" {
		dst.Domain = src.Domain
	}

	// founded_at

	if !src.FoundedAt.IsZero() {
		dst.FoundedAt = src.FoundedAt
	}

	// name

	if src.Name != "" {
		dst.Name = src.Name
	}

	// oper_category

	if src.OperationCategory != "" {
		dst.OperationCategory = src.OperationCategory
	}

	// oper_range

	if src.OperationRange != "" {
		dst.OperationRange = src.OperationRange
	}

	// reference

	if src.Reference != "" {
		dst.Reference = src.Reference
	}

	// registry

	if src.Registry != "" {
		dst.Registry = src.Registry
	}

	// remarks

	if src.Remarks != "" {
		dst.Remarks = src.Remarks
	}

	// representative

	if src.Representative != "" {
		dst.Representative = src.Representative
	}

	// state

	if src.State != "" {
		dst.State = src.State
	}

	// uri

	if src.URI != "" {
		dst.URI = src.URI
	}

	// web

	if src.Web != "" {
		dst.Web = src.Web
	}

	return nil
}

func (inst *CompanyServiceImpl) _impl() companies.Service {
	return inst
}
