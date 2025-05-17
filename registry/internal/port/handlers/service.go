package handlers

import (
	"context"
	"fmt"
	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
	"github.com/artarts36/lowboard/registry/internal/repository"
)

type Service struct {
	api.UnimplementedHandler

	repo *repository.Group
}

func NewService(repo *repository.Group) *Service {
	return &Service{repo: repo}
}

func listAndAdapt[M any, R any](ctx context.Context, list func(ctx context.Context) ([]M, error), adapt func(models []M) []R) ([]R, error) {
	items, err := list(ctx)
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}

	return adapt(items), nil
}

func getAndAdapt[M any, R any](
	ctx context.Context,
	get func(ctx context.Context) (M, error),
	adapt func(model M) R,
) (*R, error) {
	item, err := get(ctx)
	if err != nil {
		return new(R), fmt.Errorf("get: %w", err)
	}

	v := adapt(item)

	return &v, nil
}

func createAndAdapt[M any, R any](
	ctx context.Context,
	create func(ctx context.Context) (M, error),
	adapt func(model M) R,
) (*R, error) {
	item, err := create(ctx)
	if err != nil {
		return new(R), fmt.Errorf("create: %w", err)
	}

	v := adapt(item)

	return &v, nil
}

func updateAndAdapt[M any, R any](
	ctx context.Context,
	create func(ctx context.Context) (M, error),
	adapt func(model M) R,
) (*R, error) {
	item, err := create(ctx)
	if err != nil {
		return new(R), fmt.Errorf("update: %w", err)
	}

	v := adapt(item)

	return &v, nil
}
