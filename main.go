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
	personTpl.Execute(w, defaultPerson)
}

func changeNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		changeNameTpl.Execute(w, defaultPerson)
		return
	}

	r.ParseForm()

	newName, err := checkAndRefactorName(defaultPerson.Name, r.Form["name"][0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("400 bad request - %s", err.Error())))

		return
	}

	defaultPerson.Name = newName
	http.Redirect(w, r, "/", http.StatusFound)
}

func checkAndRefactorName(oldName string, newName string) (string, error) {
	newNameCleaned := strings.TrimSpace(strings.ToLower(newName))

	if newNameCleaned == "" {
		return "", errors.New("missing name")
	}

	if newNameCleaned == strings.TrimSpace(strings.ToLower(oldName)) {
		return "", errors.New("no changes made")
	}

	return capitalize(newNameCleaned), nil
}

func capitalize(str string) string {
	if str == "" {
		return ""
	}

	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
