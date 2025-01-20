package post

import (
	"database/sql"
	"github.com/AyanDas-99/blog/db"
)

func DeletePost(id string) (sql.Result, error) {
	result, err := db.DB.Exec(`DELETE FROM posts WHERE id=?`, id)
	return result, err
}
