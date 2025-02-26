package repository

import "context"

type Filter interface{}
type Update interface{}

type ReadOption[F Filter] func(F) error
type UpdateOption[U Update] func(U) error

type Repository[D any, F Filter, U Update] interface {
	Create(ctx context.Context, domain D) (*D, error)
	Read(ctx context.Context, opts ...ReadOption[F]) ([]D, error)
	Update(ctx context.Context, opts ...UpdateOption[D]) error
	Delete(ctx context.Context, opts ...ReadOption[F]) error
}
