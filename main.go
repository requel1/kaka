package main

import (
	"context"
	"fmt"
	"postgres/simple_connection"
	"postgres/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	simple_sql.SelectRows(ctx, conn)
	fmt.Println("succeed!")
}
