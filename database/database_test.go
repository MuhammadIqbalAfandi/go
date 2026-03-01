package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestExecuteContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(name) VALUES('Mark')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id:", id)
		fmt.Println("name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, balance int32
		var name string
		var email sql.NullString
		var rating float64
		var birth_date, created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			fmt.Println("Scan error:", err)
			continue
		}

		fmt.Println("id:", id)
		fmt.Println("name:", name)
		fmt.Println("email:", email)
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		fmt.Println("birth date:", birth_date)
		fmt.Println("married:", married)
		fmt.Println("created:", created_at)
		fmt.Println("---")
	}
}
