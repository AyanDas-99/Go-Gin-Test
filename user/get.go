package user

import (
	"github.com/AyanDas-99/blog/db"
)

func GetUsers() ([]User, error) {

	rows, err := db.DB.Query(`SELECT id, username, password, created_at FROM users`) // check err
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt) // check err
		if err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, err
}

func GetUserById(id string) (User, error) {
	var u User
	err := db.DB.QueryRow(`SELECT id, username, password, created_at FROM users WHERE id=?`, id).Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
	return u, err
}
