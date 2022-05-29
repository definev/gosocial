package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	CreateAt  time.Time `json:"create_at"`
	Id        int64     `json:"id"`
	UserEmail string    `json:"user_email"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}

func (c Client) CreatePost(userEmail, title, body string) (Post, error) {
	post := Post{
		CreateAt:  time.Now().UTC(),
		Id:        int64(uuid.New().ID()),
		UserEmail: userEmail,
		Title:     title,
		Body:      body,
	}

	db, err := c.readDB()
	if err != nil {
		return post, err
	}

	db.Posts[post.Id] = post
	err = c.updateDB(db)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (c Client) GetPosts() ([]Post, error) {
	db, err := c.readDB()
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(db.Posts))
	for _, post := range db.Posts {
		posts = append(posts, post)
	}

	return posts, nil
}

func (c Client) GetPost(id int64) (Post, error) {
	db, err := c.readDB()
	if err != nil {
		return Post{}, err
	}

	post, ok := db.Posts[id]
	if !ok {
		return Post{}, errors.New("post not found")
	}

	return post, nil
}

func (c Client) DeletePost(id int64) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	delete(db.Posts, id)

	return c.updateDB(db)
}

func (c Client) UpdatePost(id int64, title, body string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	post, ok := db.Posts[id]
	if !ok {
		return errors.New("post not found")
	}

	if title != "" {
		post.Title = title
	}
	if body != "" {
		post.Body = body
	}
	
	db.Posts[id] = post

	return c.updateDB(db)
}
