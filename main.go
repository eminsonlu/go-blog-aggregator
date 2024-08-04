package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/eminsonlu/go-blog-aggregator/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	config := apiConfig{
		DB: dbQueries,
	}

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

	serveMux.HandleFunc("POST /v1/users", config.handlerUserCreate)
	serveMux.HandleFunc("GET /v1/users", config.handlerUserGet)

	serveMux.HandleFunc("POST /v1/feeds", config.middlewareAuth(config.handlerFeedCreate))
	serveMux.HandleFunc("GET /v1/feeds", config.handlerGetAllFeeds)

	log.Println("Server running on port", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
