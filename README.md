## README

### Introduction
This Go program connects to a PostgreSQL database and performs database operations using functions from the `db` package.

### Prerequisites
1. Go installed on your system.
2. PostgreSQL database running.
3. Required Go packages:
   - `github.com/lib/pq`
   - `github.com/lpernett/godotenv`

### Setup
1. Clone the repository.
2. Create a `.env` file with the following environment variables:
```
HOST=your_host
USER=your_user
PASSWORD=your_password
DATABASE=your_database
```
3. Install the required Go packages:
```
go get github.com/lib/pq
go get github.com/lpernett/godotenv
```
4. Run the program:
```
go run main.go
```

### Usage
- The program connects to the PostgreSQL database using the environment variables specified in the `.env` file.
- The program currently calls the function `DeleteProductById(dbconn, 2)` from the `db` package to delete a product with ID 2. You can comment out this line and uncomment other function calls from the `db` package to perform different operations.

### Note
- Ensure that the `.env` file is correctly configured with your database credentials before running the program.
- Modify the function calls in the `main` function to perform different database operations as needed.

### Author
Ilya Fofanov

### License
This project is licensed under the MIT License.
