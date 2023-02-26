package main

import (
	"log"
	"database/sql"
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lukkas-lukkas/go-api-rest/src/application"
	database "github.com/lukkas-lukkas/go-api-rest/src/infrastructure/database"
	kafka "github.com/lukkas-lukkas/go-api-rest/src/infrastructure/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/dev_database")
	if err != nil {
		log.Fatalln(err)
	}

	repository := database.MysqlCourseRepository{Db: db}
	service := application.CreateCourse{Repository: &repository}

	messageChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":"broker:29092",
		"group.id": "golang-api",
	}

	topics := []string{"golang-api_create-course"}
	consumer := kafka.NewConsumer(configMap, topics)

	go consumer.Consume(messageChan)

	for message := range messageChan {
		var dto application.CourseDto
		json.Unmarshal(message.Value, &dto)

		result, err := service.Execute(dto)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Created:")
			fmt.Println(result)
		}
	}
}
