package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
	"GO-POSTGRESQL/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DATABASE"))

	dbconn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()
   
	//db.GetProducts(dbconn) // Вызов функции GetProducts из пакета db
	//db.GetProductById(dbconn,2) // Вызов функции GetProductById из пакета db
	//db.UpdateProductById(dbconn, 2, "test", "test", 22)
	db.DeleteProductById(dbconn,2)
	if err != nil {
		panic(err)
	}
}
