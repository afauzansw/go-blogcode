package postmodel

import (
	"go-web-crud/config"
	"go-web-crud/entities"
	"go-web-crud/utils"
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

		err := query.Scan(&post.Id, &post.Title, &post.Description, &post.Tags, &post.Status, &post.Slug, &post.CreatedAt, &post.UpdatedAt)
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
	post.Slug = utils.GenerateSlug(post.Title)

	query, err := config.DB.Exec(`
INSERT INTO posts (title, description, tags, status, slug, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)`, post.Title, post.Description, post.Tags, post.Status, post.Slug, post.CreatedAt, post.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := query.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func findById(id int) entities.Post {
	query := config.DB.QueryRow(`SELECT * FROM posts WHERE id=?`, id)

	var post entities.Post

	err := query.Scan(&post.Id, &post.Title, &post.Slug, &post.Tags, &post.Status, &post.Description, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		panic(err)
	}

	return post
}

func Update() {
	//
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM posts WHERE id=?`, id)

	return err
}
