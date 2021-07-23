package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const message = "Welcome to Server"

type NewsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "Some news for now"}
	t, _ := template.ParseFiles("basicTemplating.html")
	t.Execute(w, p)
}

func func1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
	fmt.Fprintf(w, "<h1>Welcome to HTTP server HTML Tag!</h1>")
}

func main() {
	fmt.Println("Server started")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func1)
	mux.HandleFunc("/agg", newsAggHandler)
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}
