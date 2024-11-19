package publishermodel

import (
	"go-web-crud/config"
	"go-web-crud/entities"
	"time"
)

func GetAll() []entities.Publisher {
	query, err := config.DB.Query(`SELECT * FROM publishers`)

	if err != nil {
		panic(err)
	}
	defer query.Close()

	var publishers []entities.Publisher

	for query.Next() {
		var publisher entities.Publisher

		err = query.Scan(&publisher.Id, &publisher.Name, &publisher.Email, &publisher.JobTitle, &publisher.CreatedAt, &publisher.UpdatedAt)
		if err != nil {
			panic(err)
		}

		publishers = append(publishers, publisher)
	}

	return publishers
}

func Create(publisher entities.Publisher) bool {
	publisher.CreatedAt = time.Now()
	publisher.UpdatedAt = time.Now()

	query, err := config.DB.Exec(`INSERT INTO publishers (name, email, job_title, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`, publisher.Name, publisher.Email, publisher.JobTitle, publisher.CreatedAt, publisher.UpdatedAt)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := query.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func FindById(id int) entities.Publisher {
	query := config.DB.QueryRow(`SELECT * FROM publishers WHERE id=?`, id)

	var publisher entities.Publisher

	err := query.Scan(&publisher.Id, &publisher.Name, &publisher.Email, &publisher.JobTitle, &publisher.CreatedAt, &publisher.UpdatedAt)
	if err != nil {
		panic(err)
	}

	return publisher
}

func Update(publisher entities.Publisher) bool {
	return true
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM publishers WHERE id=?", id)

	return err
}
