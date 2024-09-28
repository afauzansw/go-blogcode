package main

import (
	homecontroller "go-web-crud/controllers/home"
	postcontroller "go-web-crud/controllers/post"
	"go-web-crud/controllers/publisher"
	tagcontroller "go-web-crud/controllers/tag"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/post", postcontroller.Index)
	http.HandleFunc("/post/create", postcontroller.Create)
	http.HandleFunc("/post/edit", postcontroller.Edit)
	http.HandleFunc("/post/delete", postcontroller.Delete)
	http.HandleFunc("/post/show", postcontroller.Show)

	http.HandleFunc("/tag", tagcontroller.Index)
	http.HandleFunc("/tag/create", tagcontroller.Create)
	http.HandleFunc("/tag/edit", tagcontroller.Edit)
	http.HandleFunc("/tag/delete", tagcontroller.Delete)

	http.HandleFunc("/publisher", publisher.Index)
	http.HandleFunc("/publisher/create", publisher.Create)
	http.HandleFunc("/publisher/edit", publisher.Edit)
	http.HandleFunc("/publisher/delete", publisher.Delete)
}
