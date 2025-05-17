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
	tableSidebars = "sidebars"
)

type SidebarRepository struct {
	db *sqlx.DB
}

func NewSidebarRepository(db *sqlx.DB) *SidebarRepository {
	return &SidebarRepository{
		db: db,
	}
}

func (p *SidebarRepository) List(ctx context.Context) ([]*model.Sidebar, error) {
	var items []*model.Sidebar

	q, _, err := goqu.Select().From(tableSidebars).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.SelectContext(ctx, &items, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return items, nil
}

func (p *SidebarRepository) Get(ctx context.Context, id string) (*model.PageComponent, error) {
	var item model.PageComponent

	q, _, err := goqu.Select().From(tableSidebars).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &item, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &item, nil
}

func (p *SidebarRepository) Delete(ctx context.Context, id string) error {
	q, _, err := goqu.Delete(tableSidebars).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = p.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (p *SidebarRepository) Create(ctx context.Context, pc *model.PageComponent) (*model.PageComponent, error) {
	var newPage model.PageComponent

	pc.ID = uuid.NewString()
	pc.CreatedAt = time.Now()

	q, _, err := goqu.Insert(tableSidebars).Rows(pc).Returning("*").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &newPage, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &newPage, nil
}

func (p *SidebarRepository) Update(ctx context.Context, page *model.PageComponent) (*model.PageComponent, error) {
	var newPage model.PageComponent

	page.UpdatedAt = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	q, _, err := goqu.Update(tableSidebars).Set(page).Where(goqu.C("id").Eq(page.ID)).Returning("*").ToSQL()
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
