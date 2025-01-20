package user

import "database/sql"
import "github.com/AyanDas-99/blog/db"

func DeleteUser(id string) (sql.Result, error) {
	result, err := db.DB.Exec(`DELETE FROM users WHERE id=?`, id)
	return result, err
}
