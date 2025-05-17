package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	tableAPIMethods = "api_methods"
)

type APIMethodRepository struct {
	db *sqlx.DB
}

func NewAPIMethodRepository(db *sqlx.DB) *APIMethodRepository {
	return &APIMethodRepository{
		db: db,
	}
}

func (p *APIMethodRepository) List(ctx context.Context) ([]*model.APIMethod, error) {
	var items []*model.APIMethod

	q, _, err := goqu.Select().From(tableAPIMethods).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.SelectContext(ctx, &items, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return items, nil
}

func (p *APIMethodRepository) Get(ctx context.Context, id string) (*model.APIMethod, error) {
	var item model.APIMethod

	q, _, err := goqu.Select().From(tableAPIs).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &item, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &item, nil
}

func (p *APIMethodRepository) Delete(ctx context.Context, id string) error {
	q, _, err := goqu.Delete(tableAPIs).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = p.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (p *APIMethodRepository) Create(ctx context.Context, pc *model.APIMethod) (*model.APIMethod, error) {
	var newPage model.APIMethod

	pc.CreatedAt = time.Now()

	q, _, err := goqu.Insert(tableAPIs).Rows(pc).Returning("*").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &newPage, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &newPage, nil
}

func (p *APIMethodRepository) Update(ctx context.Context, page *model.APIMethod) (*model.APIMethod, error) {
	var newPage model.APIMethod

	page.UpdatedAt = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	q, _, err := goqu.Update(tableAPIs).Set(page).Where(goqu.C("name").Eq(page.Name)).Returning("*").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &newPage, q)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEntityNotFound
		}

		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &newPage, nil
}
