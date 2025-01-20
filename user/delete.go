package user

import "database/sql"

func DeleteUser(id string) (sql.Result, error) {
	result, err := DB.Exec(`DELETE FROM users WHERE id=?`, id)
	return result, err
}
