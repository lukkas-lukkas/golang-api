# Steps to create app using docker

### 1º Up containers
> docker-compose up -d --build

### 2º Create table
> docker exec -it db.golang-api.dev bash

> mysql --password=root dev_database

> create table courses(id varchar(255), name varchar(255), description varchar(255), status varchar(255));

### 3º Create topic on kafka-ui
> docker exec -it broker bash
> kafka-topics --create --topic golang-api_create-course --bootstrap-server localhost:9092 --replication-factor 1

### 4º Run application
> docker exec -i api.golang-api.dev go run main.go

### 5º Public on topic
> docker exec -it broker bash
> kafka-console-producer --bootstrap-server localhost:9092 --topic golang-api_create-course
> {"name": "Course 1", "description": "Description course1"}