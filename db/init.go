package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB sql.DB

func InitDb() {

	// username := "admin"                                          // From AWS RDS configuration
	// password := "testdbpassword"                                 // From AWS RDS configuration
	// endpoint := "test.czggqauoi16w.ap-south-1.rds.amazonaws.com" // e.g., "example-instance.abcdefg12345.us-east-1.rds.amazonaws.com"
	// port := "3306"                                               // Default MySQL port
	// database := "test"                                           // The specific database you want to connect to

	// dsn := fmt.Sprintf("%s:%s(%s:%s)/%s", username, password, endpoint, port, database)

	db, err := sql.Open("mysql", "admin:testdbpassword@(test.czggqauoi16w.ap-south-1.rds.amazonaws.com:3306)/test?parseTime=true")

	// db, err := sql.Open("mysql", "root:mysqlpassword@(127.0.0.1:3306)/mysql?parseTime=true")
	if err != nil {
		fmt.Println("Error connecting:", err)
	} else {
		fmt.Println("Connected to db:", db.Driver())
	}

	er := db.Ping()
	if er != nil {
		fmt.Println("Ping err:", er)
		panic(err)
	} else {
		fmt.Println("Ping success")
	}
	DB = *db

	query := `
	   CREATE TABLE users (
	       id INT AUTO_INCREMENT,
	       username TEXT NOT NULL,
	       password TEXT NOT NULL,
	       created_at DATETIME,
	       PRIMARY KEY (id)
	   );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error executing query")
	} else {
		fmt.Println("Query result", res)
	}

	query = `
	   CREATE TABLE posts (
	       id INT AUTO_INCREMENT,
	       author TEXT NOT NULL,
	       content TEXT NOT NULL,
	       created_at DATETIME,
	       PRIMARY KEY (id)
	   );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	res, err = DB.Exec(query)

	if err != nil {
		fmt.Println("Error executing query")
	} else {
		fmt.Println("Query result", res)
	}

}
