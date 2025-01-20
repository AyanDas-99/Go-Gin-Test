package post

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB sql.DB

type Post struct {
	Id        int    `json:"id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func InitDb() {

	db, err := sql.Open("mysql", "root:mysqlpassword@(127.0.0.1:3306)/mysql?parseTime=true")

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

	query := `
    CREATE TABLE posts (
        id INT AUTO_INCREMENT,
        author TEXT NOT NULL,
        content TEXT NOT NULL,
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

	DB = *db

}
