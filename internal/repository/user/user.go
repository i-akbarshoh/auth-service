package user

import (
	"context"

	u "github.com/i-akbarshoh/auth-service/internal/service/user"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(ctx context.Context, create u.Create) error {
	if _, err := r.db.NewInsert().Model(&create).ModelTableExpr("users").Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(ctx context.Context, email string) (get u.Get, err error) {
	_, err = r.db.NewSelect().Table("users").Where("email=?", email).Exec(ctx, &get)
	return
}
