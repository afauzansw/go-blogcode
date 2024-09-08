package tagcontroller

import (
	"go-web-crud/entities"
	tagmodel "go-web-crud/models/tag"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tags := tagmodel.GetAll()
	data := map[string]any{
		"tags": tags,
	}

	files, err := template.ParseFiles("views/pages/tag/index.html")

	err = files.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("views/pages/tag/create.html")

		if err != nil {
			panic(err)
		}

		err = files.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == "POST" {
		var tag = entities.Tag{}
		tag.Name = r.FormValue("name")

		if success := tagmodel.Create(tag); !success {
			files, _ := template.ParseFiles("views/pages/tag/create.html")
			files.Execute(w, nil)
		}

		http.Redirect(w, r, "/tag", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("views/pages/tag/edit.html")

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
