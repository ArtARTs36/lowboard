package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	tablePageComponents = "page_components"
)

type PageComponentRepository struct {
	db *sqlx.DB
}

func NewPageComponentRepository(db *sqlx.DB) *PageComponentRepository {
	return &PageComponentRepository{
		db: db,
	}
}

func (p *PageComponentRepository) List(ctx context.Context) ([]*model.PageComponent, error) {
	var items []*model.PageComponent

	q, _, err := goqu.Select().From(tablePageComponents).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.SelectContext(ctx, &items, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return items, nil
}

func (p *PageComponentRepository) Get(ctx context.Context, id string) (*model.PageComponent, error) {
	var item model.PageComponent

	q, _, err := goqu.Select().From(tablePageComponents).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &item, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &item, nil
}

func (p *PageComponentRepository) Delete(ctx context.Context, id string) error {
	q, _, err := goqu.Delete(tablePageComponents).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = p.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (p *PageComponentRepository) Create(ctx context.Context, pc *model.PageComponent) (*model.PageComponent, error) {
	var newPage model.PageComponent

	pc.ID = uuid.NewString()
	pc.CreatedAt = time.Now()

	q, _, err := goqu.Insert(tablePageComponents).Rows(pc).Returning("*").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &newPage, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &newPage, nil
}

func (p *PageComponentRepository) Update(ctx context.Context, page *model.PageComponent) (*model.PageComponent, error) {
	var newPage model.PageComponent

	page.UpdatedAt = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	q, _, err := goqu.Update(tablePageComponents).Set(page).Where(goqu.C("id").Eq(page.ID)).Returning("*").ToSQL()
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
