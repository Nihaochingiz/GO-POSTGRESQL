package db

import (
	"database/sql"
	"fmt"
)

func GetProductById(db *sql.DB, id int) {
	rows, err := db.Query("SELECT * FROM Products WHERE id = $1", id)
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