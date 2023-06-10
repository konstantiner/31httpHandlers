package main

import (
	"31httpHandlers/internal/app"
)

func main() {
	app.Run()
}

/*
curl -X POST -d "{\"name\": \"Vasiliy\", \"age\": 20}" http://localhost:8080/users
curl -X POST -d "{\"name\": \"Ivan\", \"age\": 30}" http://localhost:8080/users
curl -X POST -d "{\"name\": \"Boss\", \"age\": 30}" http://localhost:8080/users
curl -X POST -d "{\"sourceId\": 1, \"targetId\": 2}" http://localhost:8080/friends
curl -X DELETE -d "{\"targetId\": 3}" http://localhost:8080/users
curl -X PUT -d "{\"age\": 25}" http://localhost:8080/users/2/age

http://localhost:8080/users
http://localhost:8080/users/1/friends
*/