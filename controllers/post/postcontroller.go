package postcontroller

import (
	"go-web-crud/entities"
	postmodel "go-web-crud/models/post"
	"net/http"
	"strconv"
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

	idString := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	data := map[string]any{
		"post": postmodel.FindById(id),
	}

	files.Execute(w, data)

	if r.Method == "POST" {
		idString := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idString)

		var post = entities.Post{}
		post.Title = r.FormValue("title")
		post.Tags = r.FormValue("tags")
		post.Status = r.FormValue("status")
		post.Description = r.FormValue("desc")

		if success := postmodel.Update(id, post); !success {
			files, _ := template.ParseFiles("views/pages/post/edit.html")
			files.Execute(w, nil)
		}

		http.Redirect(w, r, "/post", http.StatusSeeOther)

	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idString)

	if err := postmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/post", http.StatusSeeOther)
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
