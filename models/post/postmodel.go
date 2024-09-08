package postmodel

import (
	"go-web-crud/config"
	"go-web-crud/entities"
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

		err := query.Scan(&post.Id, &post.Title, &post.Desc, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	return posts
}
