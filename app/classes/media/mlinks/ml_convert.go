package mlinks

import (
	"github.com/bitwormhole/markets/app/classes/utils"
)

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	utils.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Name = src.Name

	dst.TargetID = src.TargetID
	dst.TargetUUID = src.TargetUUID

	dst.ContentLength = src.ContentLength
	dst.ContentSum = src.ContentSum
	dst.ContentType = src.ContentType

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	utils.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Name = src.Name

	dst.TargetID = src.TargetID
	dst.TargetUUID = src.TargetUUID

	dst.ContentLength = src.ContentLength
	dst.ContentSum = src.ContentSum
	dst.ContentType = src.ContentType

	return nil
}
