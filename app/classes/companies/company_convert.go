package companies

import (
	"github.com/bitwormhole/markets/app/classes/utils"
	"github.com/starter-go/base/lang"
)

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	utils.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	app_at := src.ApprovedAt
	fon_at := src.FoundedAt

	dst.Address = src.Address
	dst.ApprovedAt = lang.NewTime(app_at)
	dst.Capital = src.Capital
	dst.Code = src.Code
	dst.CompanyType = src.CompanyType
	dst.FoundedAt = lang.NewTime(fon_at)
	dst.Name = src.Name
	dst.OperationCategory = src.OperationCategory
	dst.OperationRange = src.OperationRange
	dst.Reference = src.Reference
	dst.Registry = src.Registry
	dst.Remarks = src.Remarks
	dst.Representative = src.Representative
	dst.State = src.State
	dst.URI = src.URI
	dst.Web = src.Web

	return nil

}

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	utils.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Address = src.Address
	dst.ApprovedAt = src.ApprovedAt.Time()
	dst.Capital = src.Capital
	dst.Code = src.Code
	dst.CompanyType = src.CompanyType
	dst.FoundedAt = src.FoundedAt.Time()
	dst.Name = src.Name
	dst.OperationCategory = src.OperationCategory
	dst.OperationRange = src.OperationRange
	dst.Reference = src.Reference
	dst.Registry = src.Registry
	dst.Remarks = src.Remarks
	dst.Representative = src.Representative
	dst.State = src.State
	dst.URI = src.URI
	dst.Web = src.Web

	return nil

}

func ConvertListE2D(src []*Entity, dst []*DTO) ([]*DTO, error) {

	for _, it1 := range src {
		it2 := new(DTO)
		err := ConvertE2D(it1, it2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, it2)
	}

	return dst, nil
}
