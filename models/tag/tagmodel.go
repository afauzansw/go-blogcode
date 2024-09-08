package tagmodel

import (
	"go-web-crud/config"
	"go-web-crud/entities"
	"time"
)

func GetAll() []entities.Tag {
	query, err := config.DB.Query(`SELECT * FROM tags`)
	if err != nil {
		panic(err)
	}

	defer query.Close()

	var tags []entities.Tag

	for query.Next() {
		var tag entities.Tag

		err := query.Scan(&tag.Id, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			panic(err)
		}

		tags = append(tags, tag)
	}

	return tags
}

func Create(tag entities.Tag) bool {
	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()
	
	query, err := config.DB.Exec(`
INSERT INTO tags (name, created_at, updated_at)
VALUE (?, ?, ?)`, tag.Name, tag.CreatedAt, tag.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := query.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}
