package products

import (
	"github.com/bitwormhole/markets/app/classes/utils"
)

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	utils.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Name = src.Name
	dst.Code = src.Code
	dst.Label = src.Label
	dst.Remark = src.Remark
	dst.Description = src.Description

	dst.TrademarkID = src.TrademarkID
	dst.TrademarkCode = src.TrademarkCode
	dst.TrademarkName = src.TrademarkName

	dst.StandardID = src.StandardID
	dst.StandardCode = src.StandardCode

	dst.URL = src.URL
	dst.URI = src.URI
	dst.Reference = src.Reference

	return nil
}

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	utils.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Name = src.Name
	dst.Code = src.Code
	dst.Label = src.Label
	dst.Remark = src.Remark
	dst.Description = src.Description

	dst.TrademarkID = src.TrademarkID
	dst.TrademarkCode = src.TrademarkCode
	dst.TrademarkName = src.TrademarkName

	dst.StandardID = src.StandardID
	dst.StandardCode = src.StandardCode

	dst.URL = src.URL
	dst.URI = src.URI
	dst.Reference = src.Reference

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
