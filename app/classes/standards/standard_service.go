package standards

import "context"

type Service interface {

	// query

	Find(ctx context.Context, id ID) (*DTO, error)

	Query(ctx context.Context, q *Query) ([]*DTO, error)

	// edit

	Insert(ctx context.Context, item *DTO) (*DTO, error)

	Update(ctx context.Context, id ID, item *DTO) (*DTO, error)

	Remove(ctx context.Context, id ID) error
}
