# Steps to create app using docker

### 1º Up containers
> docker-compose up -d --build

### 2º Create table
> docker exec -it db.golang-api.dev bash

> mysql -uroot -p dev_database

> create table courses(id varchar(255), name varchar(255), description varchar(255), status varchar(255));

Here we can see "Hello world" cli