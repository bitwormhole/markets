package mlinks

import "context"

type Service interface {
	Find(cc context.Context, id ID) (*DTO, error)

	Insert(cc context.Context, item *DTO) (*DTO, error)
}
