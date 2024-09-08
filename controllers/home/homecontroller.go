package homecontroller

import (
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("views/pages/home/home.html")

	if err != nil {
		panic(err)
	}

	err = files.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
