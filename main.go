package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"unicode"
)

var (
	personTpl     = template.Must(template.ParseFiles("person.gohtml"))
	changeNameTpl = template.Must(template.ParseFiles("change-name.gohtml"))
)

type person struct {
	Name string
	Age  int
}

var defaultPerson = person{
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
	err := personTpl.Execute(w, defaultPerson)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
	}
}

func changeNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := changeNameTpl.Execute(w, defaultPerson)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 internal server error"))
		}
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
		return
	}

	newName, err := changeName(defaultPerson.Name, r.Form["name"][0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("400 bad request - %s", err.Error())))
	}

	defaultPerson.Name = newName

	/*	newName := r.Form["name"][0]
		if newName == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("400 bad request - missing name"))
			return
		}

		defaultPerson.Name = newName*/

	http.Redirect(w, r, "/", http.StatusFound)
}

func changeName(oldName string, newName string) (string, error) {
	if newName == "" {
		return "", errors.New("missing name")
	}

	newNameCleaned := strings.TrimSpace(strings.ToLower(newName))

	if newNameCleaned == strings.TrimSpace(strings.ToLower(oldName)) {
		return "", errors.New("no changes made")
	}

	return newNameCleaned, nil

}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
