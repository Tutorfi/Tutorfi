package app

import (
	storage "app/internal/storage/postgres"
	"log"
)

func BuildTestDB() (*storage.PostgresStorage, error) {
	db, err := storage.ConnectPgsqlTest()
	if err != nil {
		log.Println(err.Error())
	}
	poststore := storage.NewPostgresStorage(db)
	err = poststore.BuildDevDB()
	if err != nil {
		return nil, err
	}
	return poststore, nil
}
