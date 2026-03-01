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

func TestQuerySqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username, password FROM public.user WHERE username = $1 AND password = $2"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("username:", username)
		fmt.Println("password:", password)
		fmt.Println("Login success")
	}
}

func TestIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	var insertId int64
	query := "INSERT INTO comments(email, comment) VALUES('mark@example.com', 'Hello World') RETURNING id"
	err := db.QueryRowContext(ctx, query).Scan(&insertId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with ID:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("user%d@example.com", i)
		comment := fmt.Sprintf("Comment number %d", i)

		var insertId int64
		err := stmt.QueryRowContext(ctx, email, comment).Scan(&insertId)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Success insert new comment with ID: %d\n", insertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id"

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("user%d@example.com", i)
		comment := fmt.Sprintf("Comment number %d", i)

		var insertId int64
		err := tx.QueryRowContext(ctx, query, email, comment).Scan(&insertId)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Success insert new comment with ID: %d\n", insertId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
