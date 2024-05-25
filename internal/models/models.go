package models

import (
	"time"
)

type Post struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"createdAt"`
	AuthorID        string    `json:"authorId"`
	CommentsAllowed bool      `json:"commentsAllowed"`
}

type Comment struct {
	ID        string    `json:"id"`
	PostID    string    `json:"postId"`
	ParentID  *string   `json:"parentId"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"authorId"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
