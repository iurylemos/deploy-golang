package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("It's here within")
		rw.Write([]byte("Hello world"))
	})

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic("error to listen server" + err.Error())
	}
}
