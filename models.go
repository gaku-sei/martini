package main

import (
	"time"
)

type Favorite struct {
	ID        int       `json:"id"`
	GameId    int       `json:"gameId"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt"`
}

type OwnedGame struct {
	ID        int       `json:"id"`
	GameId    int       `json:"gameId"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp" json:"updatedAt"`
}
