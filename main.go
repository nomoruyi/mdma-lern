package main

import (
	"html/template"
	"net/http"
)

var (
	personTpl     = template.Must(template.ParseFiles("person.gohtml"))
	changeNameTpl = template.Must(template.ParseFiles("change-name.gohtml"))
)

type person struct {
	Name string
	Age  int
}

var daniel = person{
	Name: "Daniel",
	Age:  22,
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", personHandler)
	mux.HandleFunc("/change-name", changeNameHandler)

	srv := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	srv.ListenAndServe()
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	personTpl.Execute(w, daniel)
}

func changeNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		changeNameTpl.Execute(w, daniel)
		return
	}

	r.ParseForm()

	newName := r.Form["name"][0]
	daniel.Name = newName

	http.Redirect(w, r, "/", http.StatusFound)
}
