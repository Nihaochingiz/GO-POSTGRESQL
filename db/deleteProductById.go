package db

import (
	"database/sql"
	"fmt"
)

func DeleteProductById(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM Products WHERE id = $1", id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product with ID", id, "deleted successfully!")
}