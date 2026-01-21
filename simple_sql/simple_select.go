package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	SELECT *
	FROM pets
	ORDER BY id asc;
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			return err
		}
		PrintRows(id, name, age)

	}
	return nil
}

func PrintRows(id int, name string, age int) {
	fmt.Println("-------------------------------------")
	fmt.Println("Id: ", id)
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
}
