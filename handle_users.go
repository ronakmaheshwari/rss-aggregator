package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ronakmaheshwari/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("Error took place at decoder: %v", err)))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Email:     params.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	})

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("DB error: %v", err)))
		return
	}

	respondWithJson(w, 200, user)
}

func (apiCfg *apiConfig) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetUsers(r.Context())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("DB error: %v", err)))
		return
	}

	respondWithJson(w, 200, users)
}

func (apiCfg *apiConfig) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email");
	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	user, err := apiCfg.DB.GetUserByEmail(r.Context(), email)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("DB error: %v", err)))
		return
	}
	respondWithJson(w, 200, user)
}

func (apiCfg *apiConfig) updateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("Error took place at decoder: %v", err)))
		return
	}
	user, err := apiCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		Name:  params.Name,
		Email: params.Email,
		UpdatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	})

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("DB error: %v", err)))
		return
	}
	respondWithJson(w, 200, user)
}

func (apiCfg *apiConfig) deleteUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("Error took place at decoder: %v", err)))
		return
	}
	err = apiCfg.DB.DeleteUser(r.Context(), params.Email)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("DB error: %v", err)))
		return
	}
	respondWithJson(w, 200, map[string]string{"message": "Success"})
}