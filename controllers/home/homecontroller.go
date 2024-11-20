package homecontroller

import (
	postmodel "go-web-crud/models/post"
	publishermodel "go-web-crud/models/publisher"
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"posts": postmodel.GetAll(),
		"totalPostToday": postmodel.Count(true),
		"totalPost": postmodel.Count(false),
		"totalPublisher": publishermodel.Count(),
	}
	
	files, err := template.ParseFiles("views/pages/home/home.html")

	if err != nil {
		panic(err)
	}

	err = files.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
