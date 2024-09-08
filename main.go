package main

import (
	"go-web-crud/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	Routes()

	log.Println("Server running")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
