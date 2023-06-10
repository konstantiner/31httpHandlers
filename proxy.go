package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var counter = 0

type Handler struct {}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	firstHost := "http://localhost:8080" + r.URL.Path
	secondHost := "http://localhost:8081" + r.URL.Path

	if counter == 0 {
		_, err := http.Post(firstHost, "", bytes.NewBuffer([]byte(content)))
		if err != nil {
			fmt.Println("Сервер не доступен")
			//log.Fatal(err)
		}
		counter ++
	} else {
		_, err := http.Post(secondHost, "", bytes.NewBuffer([]byte(content)))
		if err != nil {
			fmt.Println("Сервер не доступен")
			//log.Fatal(err)
		}
		counter --
	}
}

func main() {
	log.Println(http.ListenAndServe("localhost:8877", &Handler{}))
}

//curl -X POST -d "{\"name\": \"Vasiliy\", \"age\": 20}" http://localhost:8877/test
//curl -X POST -d "{\"name\": \"Sveta\", \"age\": 20}" http://localhost:8877/users
//http://localhost:8877/users