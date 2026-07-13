package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, nil)
}

func signInPageHandler(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("sign-in.html")
	t.Execute(w, nil)
}

func signUpPageHandler(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("sign-up.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/sign-in", signInPageHandler)
	http.HandleFunc("/sign-up", signUpPageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
