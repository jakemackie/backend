package models

type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Length      int    `json:"length"`
}

type Videos []Video