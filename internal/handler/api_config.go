package handler

import (
	"net/http"

	"github.com/definev/gosocial/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

func NewApiConfig(dbClient database.Client) apiConfig {
	return apiConfig{dbClient: dbClient}
}

func (apiCfg apiConfig) EnsureDB() error {
	return apiCfg.dbClient.EnsureDB()
}

func (apiCfg apiConfig) RegisterEndpoint(m *http.ServeMux) {
	m.HandleFunc("/users", apiCfg.endpointUsersHandler)
	m.HandleFunc("/users/", apiCfg.endpointUsersHandler)
}
