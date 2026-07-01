package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user-links", getUserLinks)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUserLinks(w http.ResponseWriter, _ *http.Request) {
	//var (
	//	ctx    context.Context
	//	cancel context.CancelFunc
	//)
	//
	//timeout, err := time.ParseDuration(r.FormValue("timeout"))
	//
	//if err == nil {
	//	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	//} else {
	//	ctx, cancel = context.WithCancel(context.Background())
	//}
	//defer cancel()

	fmt.Fprint(w, "Hello World")
}
