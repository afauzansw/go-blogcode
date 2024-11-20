package publisher

import (
	"go-web-crud/entities"
	publishermodel "go-web-crud/models/publisher"
	"net/http"
	"strconv"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	publishers := publishermodel.GetAll()
	data := map[string]any{
		"publishers": publishers,
	}

	files, err := template.ParseFiles("views/pages/publisher/index.html")

	err = files.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("views/pages/publisher/create.html")

		err = files.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == "POST" {
		var publisher = entities.Publisher{}
		publisher.Name = r.FormValue("name")
		publisher.Email = r.FormValue("email")
		publisher.JobTitle = r.FormValue("job_title")

		if success := publishermodel.Create(publisher); !success {
			files, _ := template.ParseFiles("views/pages/publisher/create.html")
			files.Execute(w, nil)
		}

		http.Redirect(w, r, "/publisher", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idString)

	if r.Method == "GET" {
		files, err := template.ParseFiles("views/pages/publisher/edit.html")

		if err != nil {
			panic(err)
		}

		data := map[string]any{
			"publisher": publishermodel.FindById(id),
		}

		files.Execute(w, data)
	}

	if r.Method == "POST" {
		var publisher = entities.Publisher{}
		publisher.Name = r.FormValue("name")
		publisher.Email = r.FormValue("email")
		publisher.JobTitle = r.FormValue("job_title")

		if success := publishermodel.Update(id, publisher); !success {
			files, _ := template.ParseFiles("views/pages/publisher/edit.html")
			files.Execute(w, nil)
		}

		http.Redirect(w, r, "/publisher", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idString)

	if err := publishermodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/publisher", http.StatusSeeOther)
}
