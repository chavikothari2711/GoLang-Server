package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/chavikothari2711/demo-golang-server/internal/auth"
	"github.com/chavikothari2711/demo-golang-server/internal/database"
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
		respondWithError(w, 400, "ERROR PARSING JSON")
		return
	}

	existingUser, err := apiCfg.DB.GetUsers(r.Context(), params.Email)
	if err == nil {
		respondWithError(w, 404, "User already present with same email")
		log.Println(existingUser)
		return
	}

	user, err := apiCfg.DB.CreateUsers(r.Context(), database.CreateUsersParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
	})

	if err != nil {
		respondWithError(w, 400, "ERROR IN CREATING USER")
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
		return
	}

	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "ERROR PARSING JSON")
		log.Println(err)
		return
	}

	existingUser, err := apiCfg.DB.GetUserByAPIKeys(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 404, "User is not present, check api_key")
		log.Println(existingUser)
		return
	}

	user, err := apiCfg.DB.UpdateUsers(r.Context(), database.UpdateUsersParams{
		ApiKey:    apiKey,
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
	})

	if err != nil {
		respondWithError(w, 400, "Error updating user")
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKeys(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get user")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}