package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const message = "Welcome to MicroService Part 1"
const message1 = "Welcome to Service Section"

func function1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
	//reading data from the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
	}
	log.Printf("Data %s\n", d)
}

func function2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message1))
}

func main() {
	fmt.Println("Connecting to HTTP Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", function1)
	mux.HandleFunc("/service", function2)
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}
