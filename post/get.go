package post

import "github.com/AyanDas-99/blog/db"

func GetPosts() ([]Post, error) {

	rows, err := db.DB.Query(`SELECT id, author, content, created_at FROM posts`) // check err
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var u Post
		err := rows.Scan(&u.Id, &u.Author, &u.Content, &u.CreatedAt) // check err
		if err != nil {
			continue
		}
		posts = append(posts, u)
	}
	return posts, err
}

func GetPostById(id string) (Post, error) {
	var p Post
	err := db.DB.QueryRow(`SELECT id, author, content, created_at FROM posts WHERE id=?`, id).Scan(&p.Id, &p.Author, &p.Content, &p.CreatedAt)
	return p, err
}
