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

func Count() int {
	var totalPublishers int

	query := `SELECT COUNT(*) AS total_publishers FROM publishers`
	

	err := config.DB.QueryRow(query).Scan(&totalPublishers)
	if err != nil {
		panic(err)
	}

	return totalPublishers
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

func Update(id int, publisher entities.Publisher) bool {
	publisher.UpdatedAt = time.Now()

	query, err := config.DB.Exec(`UPDATE publishers SET name = ?, email = ?, job_title = ?, updated_at = ? WHERE id = ?`, publisher.Name, publisher.Email, publisher.JobTitle, publisher.UpdatedAt, id)
	
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
	_, err := config.DB.Exec("DELETE FROM publishers WHERE id=?", id)

	return err
}
