package main

import (
	"fmt"
	"net/http"

	"github.com/definev/gosocial/internal/database"
	"github.com/definev/gosocial/internal/handler"
)

func main() {
	m := http.NewServeMux()

	client := database.NewClient("db.json")
	apiClient := handler.NewApiConfig(client)
	apiClient.EnsureDB()
	apiClient.RegisterEndpoint(m)

	fmt.Printf("Listening on port 8080\n")
	http.ListenAndServe(":8080", m)
}
