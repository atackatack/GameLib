package entities

import "time"

type SteamDBGame struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	ReleaseDate string    `json:"release_date"`
	Types       []string  `json:"types"`
	Verified    time.Time `json:"verified"`
}
