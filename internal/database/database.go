package database

import (
	"encoding/json"
	"errors"
	"os"
)

type Client struct {
	dbPath string
}

func NewClient(dbPath string) Client {
	return Client{
		dbPath,
	}
}

func (c Client) createDB() error {
	data, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[int64]Post),
	})

	if err != nil {
		return err
	}

	err = os.WriteFile(c.dbPath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) updateDB(db databaseSchema) error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.dbPath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	var db databaseSchema

	data, err := os.ReadFile(c.dbPath)
	if err != nil {
		return db, err
	}

	err = json.Unmarshal(data, &db)
	if err != nil {
		return db, err
	}

	return db, nil
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.dbPath)

	if errors.Is(err, os.ErrNotExist) {
		return c.createDB()
	}
	return nil
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[int64]Post  `json:"posts"`
}
