package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handleValidateChirp(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}

	type returnVal struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanedBody := getCleanedBody(params.Body)

	respondWithJSON(w, http.StatusOK, returnVal{
		CleanedBody: cleanedBody,
	})

}

func getCleanedBody(body string) string {

	words := strings.Split(body, " ")

	for i, word := range words {
		wordLower := strings.ToLower(word)
		if wordLower == "kerfuffle" || wordLower == "sharbert" || wordLower == "fornax" {
			words[i] = "****"
		}

	}

	cleanedBody := strings.Join(words, " ")
	return cleanedBody
}
