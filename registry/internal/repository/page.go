package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/doug-martin/goqu/v9"

	"github.com/jmoiron/sqlx"
)

const tablePages = "pages"

type PageRepository struct {
	db *sqlx.DB
}

func NewPageRepository(db *sqlx.DB) *PageRepository {
	return &PageRepository{
		db: db,
	}
}

func (p *PageRepository) List(ctx context.Context) ([]*model.Page, error) {
	var items []*model.Page

	q, _, err := goqu.Select().From(tablePages).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.SelectContext(ctx, &items, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return items, nil
}

func (p *PageRepository) Get(ctx context.Context, name string) (*model.Page, error) {
	var item model.Page

	q, _, err := goqu.Select().From(tablePages).Where(goqu.C("name").Eq(name)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &item, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &item, nil
}

func (p *PageRepository) Delete(ctx context.Context, name string) error {
	q, _, err := goqu.Delete(tablePages).Where(goqu.C("name").Eq(name)).ToSQL()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = p.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (p *PageRepository) Create(ctx context.Context, page *model.Page) (*model.Page, error) {
	var newPage model.Page

	page.CreatedAt = time.Now()

	q, _, err := goqu.Insert(tablePages).Rows(page).Returning("*").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &newPage, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &newPage, nil
}

func (p *PageRepository) Update(ctx context.Context, page *model.Page, pageName string) (*model.Page, error) {
	var newPage model.Page

	page.UpdatedAt = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	q, _, err := goqu.Update(tablePages).Set(page).Where(goqu.C("name").Eq(pageName)).Returning("*").ToSQL()
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
