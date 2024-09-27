package postmodel

import (
	"go-web-crud/config"
	"go-web-crud/entities"
	"time"
)

func GetAll() []entities.Post {
	query, err := config.DB.Query(`SELECT * FROM posts`)
	if err != nil {
		panic(err)
	}

	defer query.Close()

	var posts []entities.Post

	for query.Next() {
		var post entities.Post

		err := query.Scan(&post.Id, &post.Title, &post.Desc, &post.Tags, &post.Status, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	return posts
}

func Create(post entities.Post) bool {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	query, err := config.DB.Exec(`
INSERT INTO posts (title, desc, tags, status, created_at, updated_at)
VALUE (?, ?, ?, ?, ?, ?)`, post.Title, post.Desc, post.Tags, post.Status, post.CreatedAt, post.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := query.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func findById() {
	//
}

func Update() {
	//
}

func Delete() {
	//
}
