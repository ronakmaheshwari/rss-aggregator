package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port :=  os.Getenv("PORT");
	if port == "" {
		log.Fatal(`Port must be set`)
	}
	
	router := chi.NewRouter();
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
  	}))

	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Printf(`RSS Aggregator is running on http://localhost:%v`, port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}