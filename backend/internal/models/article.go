package models

import "time"

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	IsDeleted bool      `json:"isDeleted"`
}
