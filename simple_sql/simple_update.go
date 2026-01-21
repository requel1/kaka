package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRows(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	UPDATE pets 
	SET name = 'alex'
	WHERE id % 2 = 0;
	
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
