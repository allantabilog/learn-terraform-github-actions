package sample

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/url"
)

func main() {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:4532",
		User:   url.UserPassword("user", "password"),
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("sql.Open failed", err)
		return
	}
	//defer db.Close()

	defer func() {
		_ = db.Close()
		fmt.Println("Database closed")
	}()

}
