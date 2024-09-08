package postcontroller

import (
	postmodel "go-web-crud/models"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	posts := postmodel.GetAll()
	data := map[string]any{
		"post": posts,
	}

	files, err := template.ParseFiles("views/post/index.html")
	if err != nil {
		panic(err)
	}

	err = files.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("views/post/create.html")

	if err != nil {
		panic(err)
	}

	err = files.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("views/post/edit.html")

	if err != nil {
		panic(err)
	}

	err = files.Execute(w, nil)
	if err != nil {
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	//
}
