package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const message = "Welcome to Boilerplate Website"

func firstFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logger := log.New(os.Stdout, "guck", log.LstdFlags|log.Lshortfile)
	fmt.Println("Connecting to Http server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", firstFunc)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	logger.Println("Server Starting")
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
	tc, _ := context.WithDeadline(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
