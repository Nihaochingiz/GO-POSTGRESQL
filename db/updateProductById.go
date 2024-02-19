package db

import (
	"database/sql"
	"fmt"
)

func UpdateProductById(db *sql.DB, id int, newModel string, newCompany string, newPrice float64) {
	_, err := db.Exec("UPDATE Products SET model = $1, company = $2, price = $3 WHERE id = $4", newModel, newCompany, newPrice, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product updated successfully!")
}