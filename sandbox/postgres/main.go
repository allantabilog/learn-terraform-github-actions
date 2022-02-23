package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"time"
)

const DatabaseUrl = "postgres://postgres:postgres@localhost:5438"

func main() {
	//basicQuery()
	executeFunction()

}

func executeFunction() {
	dbPool, err := pgxpool.Connect(context.Background(), DatabaseUrl)
	defer dbPool.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	}

	const id = 1
	rows, err := dbPool.Query(context.Background(), "select * from public.get_person_details($1)", id)
	if err != nil {
		log.Fatal("Error while executing query")
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("Error while iterating over dataset")
		}

		firstName := values[0].(string)
		lastName := values[1].(string)
		dateOfBirth := values[2].(time.Time)

		log.Printf("firstName: %v, lastName: %v, dateOfBirth: %v", firstName, lastName, dateOfBirth)
	}

}

func basicQuery() {
	dbPool, err := pgxpool.Connect(context.Background(), DatabaseUrl)
	defer dbPool.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Database pool: %v", dbPool.Config())

	rows, err := dbPool.Query(context.Background(), "select * from public.person")
	if err != nil {
		log.Fatal("Error while executing query")
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("Error while iterating dataset")
		}
		//   convert DB types to go types
		id := values[0].(int32)
		firstName := values[1].(string)
		lastName := values[2].(string)
		dateOfBirth := values[3].(time.Time)

		log.Printf("id: %v firstName: %v lastName: %v dateOfBirth: %v\n", id, firstName, lastName, dateOfBirth)
	}
}

func poolTest1() {
	dbPool, err := pgxpool.Connect(context.Background(), DatabaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Database pool: %v", dbPool.Config())
	defer dbPool.Close()

}
