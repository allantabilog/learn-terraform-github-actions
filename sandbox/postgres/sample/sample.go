package sample

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func SampleMain() {
	const (
		DatabaseUrl    = "postgres://postgres:postgres@localhost:5438"
		findPersonById = "select first_name, last_name from person where id=$1"
	)

	conn, err := pgx.Connect(context.Background(), DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var firstName string
	var lastName string
	err = conn.QueryRow(context.Background(), findPersonById, 1).Scan(&firstName, &lastName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(firstName, lastName)
}
