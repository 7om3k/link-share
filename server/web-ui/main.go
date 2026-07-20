package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func homePageHandler(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("public/home.html")
	t.Execute(w, nil)
}

func signInPageHandler(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("public/sign-in.html")
	t.Execute(w, nil)
}

func signUpPageHandler(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("public/sign-up.html")
	t.Execute(w, nil)
}

func userLinksPageHandler(w http.ResponseWriter, _ *http.Request) {
	resp, err := http.Get("http://user-links:5000/api/user-links")

	if err != nil {
		http.Error(w, "An error occurred while fetching user links", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "User links fetch error: %v\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "An error occurred while fetching user links", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "User links fetch error - bad status: %v\n", err)
		return
	}

	type userLink struct {
		Id          int64
		Title       string
		Description *string
		Url         string
	}
	userLinkList := make([]userLink, 0)

	// TODO: switch to decoder v2, when no longer experimental
	dec := json.NewDecoder(resp.Body)

	// Endpoint returns array of user links, token bellow read open bracket "[" from json.
	_, err = dec.Token()
	if err != nil {
		http.Error(w, "An error occurred while fetching user links", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "User links decode error: %v\n", err)
		return
	}

	for dec.More() {
		var l userLink
		if err := dec.Decode(&l); err != nil {
			http.Error(w, "An error occurred while fetching user links", http.StatusInternalServerError)
			fmt.Fprintf(os.Stderr, "User links decode error: %v\n", err)
			return
		}
		userLinkList = append(userLinkList, l)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		http.Error(w, "An error occurred while fetching user links", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "User links decode error: %v\n", err)
		return
	}

	temp, _ := template.ParseFiles("public/user-links.html")
	temp.Execute(w, userLinkList)
}

func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/sign-in", signInPageHandler)
	http.HandleFunc("/sign-up", signUpPageHandler)
	http.HandleFunc("/user-links", userLinksPageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
