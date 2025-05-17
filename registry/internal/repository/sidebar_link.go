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
	tableSidebarLinks = "sidebar_links"
)

type SidebarLinkRepository struct {
	db *sqlx.DB
}

func NewSidebarLinkRepository(db *sqlx.DB) *SidebarLinkRepository {
	return &SidebarLinkRepository{
		db: db,
	}
}

func (p *SidebarLinkRepository) List(ctx context.Context) ([]*model.SidebarLink, error) {
	var items []*model.SidebarLink

	q, _, err := goqu.Select().From(tableSidebarLinks).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.SelectContext(ctx, &items, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return items, nil
}

func (p *SidebarLinkRepository) Tree(ctx context.Context) ([]*model.SidebarLink, error) {
	var items []*model.SidebarLink

	q, _, err := goqu.Select().From(tableSidebarLinks).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.SelectContext(ctx, &items, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	itemByParentMap := map[string][]*model.SidebarLink{}
	roots := []*model.SidebarLink{}

	for _, item := range items {
		if item.ParentID.Valid {
			if _, exists := itemByParentMap[item.ParentID.String]; !exists {
				itemByParentMap[item.ParentID.String] = []*model.SidebarLink{}
			}

			itemByParentMap[item.ParentID.String] = append(itemByParentMap[item.ParentID.String], item)
		} else {
			roots = append(roots, item)
		}
	}

	for _, item := range items {
		item.Children = itemByParentMap[item.ID]
	}

	return roots, nil
}

func (p *SidebarLinkRepository) Get(ctx context.Context, id string) (*model.PageComponent, error) {
	var item model.PageComponent

	q, _, err := goqu.Select().From(tableSidebarLinks).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &item, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &item, nil
}

func (p *SidebarLinkRepository) Delete(ctx context.Context, id string) error {
	q, _, err := goqu.Delete(tableSidebarLinks).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = p.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (p *SidebarLinkRepository) Create(ctx context.Context, pc *model.SidebarLink) (*model.SidebarLink, error) {
	var newPage model.SidebarLink

	pc.ID = uuid.NewString()
	pc.CreatedAt = time.Now()

	q, _, err := goqu.Insert(tableSidebarLinks).Rows(pc).Returning("*").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	err = p.db.GetContext(ctx, &newPage, q)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return &newPage, nil
}

func (p *SidebarLinkRepository) Update(ctx context.Context, page *model.PageComponent) (*model.PageComponent, error) {
	var newPage model.PageComponent

	page.UpdatedAt = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	q, _, err := goqu.Update(tableSidebarLinks).Set(page).Where(goqu.C("id").Eq(page.ID)).Returning("*").ToSQL()
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
