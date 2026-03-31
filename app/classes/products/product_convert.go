package products

import (
	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Name = src.Name
	dst.Code = src.Code
	dst.URL = src.URL
	dst.Label = src.Label
	dst.Description = src.Description

	return nil
}

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Name = src.Name
	dst.Code = src.Code
	dst.URL = src.URL
	dst.Label = src.Label
	dst.Description = src.Description

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
