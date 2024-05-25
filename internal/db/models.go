package db

import "time"

type Post struct {
	ID            string
	Title         string
	Content       string
	AllowComments bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Comment struct {
	ID        string
	PostID    string
	ParentID  *string
	Content   string
	CreatedAt time.Time
}
