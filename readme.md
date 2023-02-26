# GOLANG API
Consume kafka and persist at mysql

### 1º Build and run application
> make build-run

### 2º Public on topic
- Open kafka ui: http://localhost:7002/topics/golang-api_create-course

Publish message with format
```
{
    "name": "Course name",
    "description": "Course description"
}
```