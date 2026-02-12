package category

import (
	"context"
	"database/sql"
	"errors"
)

type Postgres struct {
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (repository *Postgres) Save(ctx context.Context, tx *sql.Tx, category Category) (Category, error) {
	SQL := "insert into category(name) values($1) returning id"

	err := tx.QueryRowContext(ctx, SQL, category.Name).Scan(&category.Id)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (repository *Postgres) Update(ctx context.Context, tx *sql.Tx, category Category) (Category, error) {
	SQL := "update category set name=$1 where id=$2"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (repository *Postgres) Delete(ctx context.Context, tx *sql.Tx, category Category) error {
	SQL := "delete from category where id=$1"
	_, err := tx.ExecContext(ctx, SQL, category.Id)

	if err != nil {
		return err
	}

	return nil
}

func (repository *Postgres) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (Category, error) {
	SQL := "select id, name from category where id=$1"

	category := Category{}

	err := tx.QueryRowContext(ctx, SQL, categoryId).Scan(&category.Id, &category.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Category{}, err
		}
	}

	return category, nil
}

func (repository *Postgres) FindAll(ctx context.Context, tx *sql.Tx) ([]Category, error) {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return []Category{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var categories []Category
	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return []Category{}, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
