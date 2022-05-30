package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/definev/gosocial/internal/database"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/err", testErrHandler)

	http.ListenAndServe(":8080", m)
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, errors.New("test error"))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	// you can use any compatible type, but let's use our database package's User type for practice
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	var errorMessage any

	json.Unmarshal([]byte(fmt.Sprintf(`{"error": "internal server error", "message": "%s"}`, err.Error())), &errorMessage)

	w.WriteHeader(500)
	handlerResponse, _ := json.Marshal(responseBody{
		ErrorCode:    "99",
		ErrorMessage: errorMessage,
		Data:         nil,
	})
	w.Write(handlerResponse)
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if payload != nil {
		w.WriteHeader(code)
		_, err := json.Marshal(payload)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		handlerResponse, _ := json.Marshal(responseBody{
			ErrorCode:    "00",
			ErrorMessage: "Success",
			Data:         payload,
		})

		w.Write(handlerResponse)
	}
}

type responseBody struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage any    `json:"error_message"`
	Data         any    `json:"data"`
}
