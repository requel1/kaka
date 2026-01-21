package main

import (
	"context"
	"postgres/simple_connection"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
}
