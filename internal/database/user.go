package database

import (
	"errors"
	"time"
)

type User struct {
	CreateAt time.Time `json:"create_at"`
	Name     string    `json:"name"`
	Age      uint      `json:"age"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func (c Client) CreateUser(email, password, name string, age uint) (User, error) {
	user := User{
		Name:     name,
		Age:      age,
		Email:    email,
		Password: password,
		CreateAt: time.Now().UTC(),
	}

	db, err := c.readDB()
	if err != nil {
		return user, err
	}
	db.Users[user.Email] = user
	return user, c.updateDB(db)
}

func (c Client) UpdateUser(email, password, name string, age uint) (User, error) {
	user, err := c.GetUser(email)
	if err != nil {
		return user, err
	}

	user.Name = name
	user.Age = age
	user.Password = password

	db, err := c.readDB()
	if err != nil {
		return user, err
	}

	db.Users[user.Email] = user
	return user, c.updateDB(db)
}

func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	delete(db.Users, email)

	return c.updateDB(db)
}

func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	user, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}
