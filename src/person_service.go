package src

import "net/http"
import (
	"html/template"
)

type person struct {
	Name string
	Age  int
}

var (
	personTpl     = template.Must(template.ParseFiles("person.gohtml"))
	changeNameTpl = template.Must(template.ParseFiles("change-name.gohtml"))
)

var defaultPerson = person{
	Name: "Daniel",
	Age:  22,
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	personTpl.Execute(w, defaultPerson)
}

func changeNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		changeNameTpl.Execute(w, defaultPerson)
		return
	}

	r.ParseForm()

	newName := r.Form["name"][0]
	defaultPerson.Name = newName

	http.Redirect(w, r, "/", http.StatusFound)
}
