package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/arctiqteam/random-beer-go/randombeer"
)

func thisbeer(w http.ResponseWriter, r *http.Request) {
	joke := randombeer.Randombeerrequest{}
	joke.GetData()
	t, err := template.ParseFiles("templates/randombeer.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, joke)
}

func setupRoutes() {
	http.HandleFunc("/", thisbeer)
}

func main() {
	fmt.Println("Go Web App Started")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}
