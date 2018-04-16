package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoPageData struct {
	PageTitle string
	Body      string
	ID        int
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	id := 0
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "Enron Mail tagger",
			Body:      "Body" + string(id),
			ID:        id,
		}
		id++
		tmpl.Execute(w, data)
	}).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received")
		bytes, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(bytes))
	}).Methods("POST")

	http.ListenAndServe(":8080", router)
}
