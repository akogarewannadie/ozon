package db

import (
	"sync"
)

type InMemoryDB struct {
	posts        map[string]Post
	comments     map[string][]Comment
	postMutex    sync.RWMutex
	commentMutex sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		posts:    make(map[string]Post),
		comments: make(map[string][]Comment),
	}
}
