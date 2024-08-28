package models

type App struct {
	ID  int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}