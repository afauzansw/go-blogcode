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

func Count(isToday bool) int {
	var totalPosts int
	var query string

	if isToday {
		query = `SELECT COUNT(*) AS total_posts FROM posts WHERE DATE(created_at) = CURRENT_DATE()`
	} else {
		query = `SELECT COUNT(*) AS total_posts FROM posts`
	}
	

	err := config.DB.QueryRow(query).Scan(&totalPosts)
	if err != nil {
		panic(err)
	}

	return totalPosts
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

func FindById(id int) entities.Post {
	query := config.DB.QueryRow(`SELECT * FROM posts WHERE id=?`, id)

	var post entities.Post

	err := query.Scan(&post.Id, &post.Title, &post.Description, &post.Tags, &post.Status, &post.Slug, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		panic(err)
	}

	return post
}

func Update(id int, post entities.Post) bool {
	post.UpdatedAt = time.Now()

	query, err := config.DB.Exec(`UPDATE posts SET title = ?, tags = ?, status = ?, description = ?, updated_at = ? WHERE id = ?`,
		post.Title, post.Tags, post.Status, post.Description, post.UpdatedAt, id)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0

}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM posts WHERE id=?`, id)

	return err
}
