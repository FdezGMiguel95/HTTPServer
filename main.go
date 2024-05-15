package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func Index(w http.ResponseWriter, r *http.Request) {
	template, error := template.ParseFiles("templates/index.html")

	u := User{Name: "Miguel", Age: 28}

	if error != nil {
		panic(error)
	}

	template.Execute(w, u)
}
func main() {
	//
	http.HandleFunc("/", Index)
	//crea el server
	fmt.Println("Servidor en Go en 8080")
	http.ListenAndServe("localhost:8080", nil)

}
