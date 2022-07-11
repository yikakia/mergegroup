package mergegroup

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type mergeGroup[T any] struct {
	errG *errgroup.Group
}

func New[T any]() *mergeGroup[T] {
	return &mergeGroup[T]{errG: &errgroup.Group{}}
}

func WithContext[T any](ctx context.Context) (*mergeGroup[T], context.Context) {
	errG, ctx := errgroup.WithContext(ctx)
	return &mergeGroup[T]{errG: errG}, ctx
}

func (m *mergeGroup[T]) Wait() error {
	return m.errG.Wait()
}

func (m *mergeGroup[T]) Go(dst []T, target int, fn func() (T, error)) {
	m.errG.Go(func() error {
		t, err := fn()
		if err != nil {
			return err
		}
		dst[target] = t
		return nil

	})
}
