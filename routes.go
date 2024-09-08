package main

import (
	homecontroller "go-web-crud/controllers/home"
	postcontroller "go-web-crud/controllers/post"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/posts", postcontroller.Index)
	http.HandleFunc("/posts/create", postcontroller.Create)
	http.HandleFunc("/posts/edit", postcontroller.Edit)
	http.HandleFunc("/posts/delete", postcontroller.Delete)
}
