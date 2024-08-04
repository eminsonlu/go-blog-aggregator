package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	serveMux.HandleFunc("/v1/healthz", handlerHealthz)
	serveMux.HandleFunc("/v1/err", handlerError)

	log.Println("Server running on port", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
