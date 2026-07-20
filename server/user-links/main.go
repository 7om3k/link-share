package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	http.HandleFunc("/api/user-links", getUserLinks)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

type UserLink struct {
	Id          int64
	Title       string
	Description *string
	Url         string
}

func getUserLinks(w http.ResponseWriter, _ *http.Request) {
	dbUserPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbUserPassword == "" {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: password is empty")
		return
	}

	databaseUrl := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", dbUserPassword),
		Host:   "user-links-db:5432",
		Path:   "postgres",
	}

	conn, err := pgx.Connect(context.Background(), databaseUrl.String())
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, title, description, url FROM link")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Query error: %v\n", err)
		return
	}
	defer rows.Close()

	var links []UserLink
	for rows.Next() {
		var link UserLink

		if err := rows.Scan(&link.Id, &link.Title, &link.Description, &link.Url); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			fmt.Fprintf(os.Stderr, "Row scan error: %v\n", err)
			return
		}

		links = append(links, link)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(links); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "JSON encoding error: %v\n", err)
		return
	}
}
