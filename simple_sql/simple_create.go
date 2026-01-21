package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE pets(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	age INTEGER NOT NULL

	);
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
