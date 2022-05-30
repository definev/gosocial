package handler

import (
	"errors"
	"net/http"

	"github.com/definev/gosocial/internal/database"
)

func TestErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, errors.New("test error"))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// you can use any compatible type, but let's use our database package's User type for practice
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}
