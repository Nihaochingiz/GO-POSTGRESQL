package db

import (
	"database/sql"
	"fmt"
)

type product struct {
	id     int
	model  string
	company string
	price  int
}

func GetProducts(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Println(p.id, p.model, p.company, p.price)
	}
}