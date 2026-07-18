package utils

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
)

func CopyBaseFieldsFromDtoToEntity(src *rbac.BaseDTO, dst *rbac.BaseEntity) {

	dst.UUID = src.UUID

	dst.CreatedAt = src.CreatedAt.Time()
	dst.UpdatedAt = src.UpdatedAt.Time()

	dst.Creator = src.Creator
	dst.Updater = src.Updater
	dst.Owner = src.Owner
	dst.Group = src.Group

}

func CopyBaseFieldsFromEntityToDTO(src *rbac.BaseEntity, dst *rbac.BaseDTO) {

	dst.UUID = src.UUID

	dst.CreatedAt = lang.NewTime(src.CreatedAt)
	dst.UpdatedAt = lang.NewTime(src.UpdatedAt)

	dst.Creator = src.Creator
	dst.Updater = src.Updater
	dst.Owner = src.Owner
	dst.Group = src.Group

}
