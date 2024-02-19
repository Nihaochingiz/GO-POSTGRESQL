package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Insert(db *sql.DB, model string, company string, price int) error {
	_, err := db.Exec("INSERT INTO Products (model, company, price) VALUES ($1, $2, $3)", model, company, price)
	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted into the database")
	return nil
}
