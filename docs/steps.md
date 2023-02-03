# Steps to create app using docker

### 1º create main.go
### 2° create go module with command
> docker run -it -v ${PWD}:/app -w /app golang:1.19 go mod init github.com/lukkas-lukkas/go-api-rest
### 3º create dockerfile
### 4º create docker-compose.yml
### 5º run
> docker-compose up --build

Here we can see "Hello world" cli