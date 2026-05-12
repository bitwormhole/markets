package licences

import (
	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.ID = src.ID

	return nil

}

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Code = src.Code
	dst.Type = src.Type
	dst.Reference = src.Reference
	dst.URI = src.URI
	dst.Remarks = src.Remarks

	dst.NotBefore = src.NotBefore.Time()
	dst.NotAfter = src.NotAfter.Time()

	dst.IssuerAddress = src.IssuerAddress
	dst.IssuerName = src.IssuerName

	dst.SubjectAddress = src.SubjectAddress
	dst.SubjectCode = src.SubjectCode
	dst.SubjectName = src.SubjectName

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
