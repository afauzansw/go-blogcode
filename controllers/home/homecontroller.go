package homecontroller

import (
	postmodel "go-web-crud/models/post"
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	posts := postmodel.GetAll()
	data := map[string]any{
		"posts": posts,
		"totalPostToday": 3,
		"totalPost": 15,
		"totalPublisher": 5,
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
