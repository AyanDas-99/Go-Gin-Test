package post

import (
	"database/sql"
	"fmt"
	"time"
)

func PutPost(post Post) (sql.Result, error) {
	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	fmt.Println("Putting post with:", post)
	result, err := DB.Exec(`INSERT INTO posts (author, content, created_at) VALUES (?, ?, ?)`, post.Author, post.Content, time.Now())
	return result, err
}
