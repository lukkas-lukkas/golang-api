package main

import (
	"log"
	"database/sql"
	database "github.com/lukkas-lukkas/go-api-rest/src/infrastructure/database"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/dev_database")
	if err != nil {
		log.Fatalln(err)
	}

	repository := database.MysqlCourseRepository{
		Db: db
	}

	service :=
}
