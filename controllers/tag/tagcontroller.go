package tagcontroller

import (
	"go-web-crud/entities"
	tagmodel "go-web-crud/models/tag"
	"net/http"
	"strconv"
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
	if r.Method == "GET" {
		files, err := template.ParseFiles("views/pages/tag/edit.html")

		idString := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		data := map[string]any{
			"tag": tagmodel.FindById(id),
		}

		files.Execute(w, data)
	}

	if r.Method == "POST" {
		idString := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idString)

		var tag = entities.Tag{}
		tag.Name = r.FormValue("name")

		if success := tagmodel.Update(id, tag); !success {
			files, _ := template.ParseFiles("views/pages/tag/edit.html")
			files.Execute(w, nil)
		}

		http.Redirect(w, r, "/tag", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idString)

	if err := tagmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/tag", http.StatusSeeOther)
}
