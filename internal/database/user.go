package database

import (
	"errors"
	"time"

	"github.com/definev/gosocial/internal/model"
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

func (c Client) UpdateUser(email string, password, name model.Maybe[string], age model.Maybe[uint]) (User, error) {
	user, err := c.GetUser(email)
	if err != nil {
		return user, err
	}

	if !name.Nil {
		user.Name = name.Value
	}
	if !age.Nil {
		user.Age = age.Value
	}
	if !password.Nil {
		user.Password = password.Value
	}

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

func (c Client) GetUsers() ([]User, error) {
	db, err := c.readDB()
	if err != nil {
		return ([]User)(nil), err
	}
	users := make([]User, 0, len(db.Users))
	for _, user := range db.Users {
		users = append(users, user)
	}
	return users, nil
}
