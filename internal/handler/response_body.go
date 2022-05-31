package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseBody struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage any    `json:"error_message"`
	Data         any    `json:"data"`
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	var errorMessage any

	json.Unmarshal([]byte(fmt.Sprintf(`{"error": "internal server error", "message": "%s"}`, err.Error())), &errorMessage)

	handlerResponse, _ := json.Marshal(responseBody{
		ErrorCode:    "99",
		ErrorMessage: errorMessage,
		Data:         nil,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(500)
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
