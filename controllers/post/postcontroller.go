package postcontroller

import (
	"go-web-crud/entities"
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

	if r.Method == "POST" {
		var post = entities.Post{}
		post.Title = r.FormValue("title")
		post.Status = r.FormValue("status")
		post.Tags = r.FormValue("tags")
		post.Description = r.FormValue("desc")

		if success := postmodel.Create(post); !success {
			files, _ := template.ParseFiles("views/pages/post/create.html")
			files.Execute(w, nil)
		}

		http.Redirect(w, r, "/post", http.StatusSeeOther)

	}
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

func Show(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("views/pages/public/post.html")

	if err != nil {
		panic(err)
	}

	err = files.Execute(w, nil)
	if err != nil {
	}
}
