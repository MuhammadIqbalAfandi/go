package database

import (
	"context"
	"fmt"
	"testing"

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
