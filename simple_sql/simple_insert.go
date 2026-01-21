package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	id int,
	name string,
	age int,

) error {
	sqlQuery := `
	INSERT INTO pets(id,name,age)
	VALUES($1,$2,$3)
	`
	_, err := conn.Exec(ctx, sqlQuery, id, name, age)
	return err
}
