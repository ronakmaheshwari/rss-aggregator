package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ronakmaheshwari/rss-aggregator/internal/database"
)

type Status struct {
	health string
	message string
	error bool
	ok bool
}

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port :=  os.Getenv("PORT");
	if port == "" {
		log.Fatal(`Port must be set`)
	}

	dbURL :=  os.Getenv("DB_URL");
	if dbURL == "" {
		log.Fatal(`Database URL must be set`)
	}

	conn, err := sql.Open("postgres", dbURL);
	if err != nil {
		log.Fatal(`Database URL issue couldnt connect`, err);
	}

	defer conn.Close();

	apiConfig := apiConfig {
		DB: database.New(conn),
	}

	router := chi.NewRouter();
	v1Router := chi.NewRouter();

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

	router.Mount("/api/v1", v1Router);

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.Get("/healthz", healthController);
	v1Router.Post("/create", apiConfig.handlerCreateUser);
	v1Router.Get("/users", apiConfig.getUsers);
	v1Router.Get("/users/{email}", apiConfig.getUserByEmail);
	v1Router.Get("/users/apikey", apiConfig.GetUserByApikey);
	v1Router.Patch("/", apiConfig.updateUser);
	v1Router.Delete("/", apiConfig.deleteUser);

	fmt.Printf(`RSS Aggregator is running on http://localhost:%v`, port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}