package post

import "database/sql"

func DeletePost(id string) (sql.Result, error) {
	result, err := DB.Exec(`DELETE FROM posts WHERE id=?`, id)
	return result, err
}
