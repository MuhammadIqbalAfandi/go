package category

import (
	"context"
	"database/sql"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, category Category) (Category, error)
	Update(ctx context.Context, tx *sql.Tx, category Category) (Category, error)
	Delete(ctx context.Context, tx *sql.Tx, category Category) error
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]Category, error)
}
