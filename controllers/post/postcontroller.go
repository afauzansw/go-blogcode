package postcontroller

import (
	postmodel "go-web-crud/models/post"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	posts := postmodel.GetAll()
	data := map[string]any{
		"posts": posts,
	}

	files, err := template.ParseFiles("views/pages/post/index.html")

	err = files.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("views/pages/post/create.html")

		if err != nil {
			panic(err)
		}

		err = files.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}

	//if r.Method == "POST" {
	//	var post = entities.Post{}
	//}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("views/pages/post/edit.html")

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
