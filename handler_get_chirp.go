package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {

	requestedChirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid chirp ID", err)
	}

	dbChirp, err := cfg.db.GetChirp(r.Context(), requestedChirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't create chirp", err)
	}

	responseChirp := Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	}

	respondWithJSON(w, http.StatusOK, responseChirp)

}
