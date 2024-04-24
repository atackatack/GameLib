package entities

import (
	"time"
)

type Game struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Done     bool   `json:"done"`
	Favorite bool   `json:"favorite"`

	ImageURL *string `json:"image_url"`

	HowLongToBeatID       int `json:"hltb_id"`
	HowLongToBeatMainTime int `json:"hltb_main_time"`
	HowLongToBeatFullTime int `json:"hltb_full_time"`

	CreateDt time.Time `json:"create_dt"`
	UpdateDt time.Time `json:"update_dt"`
}

type CreateGame struct {
	Name  string  `json:"name"`
	Done  bool    `json:"done"`
	Image *string `json:"image"`

	HowLongToBeatID       int `json:"hltb_id"`
	HowLongToBeatMainTime int `json:"hltb_main_time"`
	HowLongToBeatFullTime int `json:"hltb_full_time"`

	FindGrid       bool `json:"find_grid"`
	ClearPathImage bool `json:"clear_path_image"`
}

type UpdateGame struct {
	Name  string  `json:"name"`
	Done  *bool   `json:"done"`
	Image *string `json:"image"`

	HowLongToBeatID       int `json:"hltb_id"`
	HowLongToBeatMainTime int `json:"hltb_main_time"`
	HowLongToBeatFullTime int `json:"hltb_full_time"`

	FindGrid       bool `json:"find_grid"`
	ClearPathImage bool `json:"clear_path_image"`
}
