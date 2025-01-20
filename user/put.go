package user

import (
	"database/sql"
	"fmt"
	"github.com/AyanDas-99/blog/db"
	"time"
)

func PutUsers(user User) (sql.Result, error) {
	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	fmt.Println("Putting user with:", user)
	result, err := db.DB.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, user.Username, user.Password, time.Now())
	return result, err
}
