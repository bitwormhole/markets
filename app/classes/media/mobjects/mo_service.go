package mobjects

import "context"

type Service interface {
	Find(cc context.Context, id ID) (*DTO, error)

	FindBySum(cc context.Context, sum []byte) (*DTO, error)

	ContainsSum(cc context.Context, sum []byte) (bool, error)

	Insert(cc context.Context, item *DTO) (*DTO, error)
}
