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
	tableAPIActions = "api_actions"
)

type APIActionRepository struct {
	db *sqlx.DB
}

func NewAPIActionRepository(db *sqlx.DB) *APIActionRepository {
	return &APIActionRepository{
		db: db,
	}
}

func (p *APIActionRepository) List(ctx context.Context) ([]*model.APIAction, error) {
	var items []*model.APIAction

	q, _, err := goqu.Select().From(tableAPIActions).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.SelectContext(ctx, &items, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return items, nil
}

func (p *APIActionRepository) Get(ctx context.Context, id string) (*model.APIAction, error) {
	var item model.APIAction

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

func (p *APIActionRepository) Delete(ctx context.Context, id string) error {
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

func (p *APIActionRepository) Create(ctx context.Context, pc *model.APIAction) (*model.APIAction, error) {
	var newPage model.APIAction

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

func (p *APIActionRepository) Update(ctx context.Context, page *model.APIAction) (*model.APIAction, error) {
	var newPage model.APIAction

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
