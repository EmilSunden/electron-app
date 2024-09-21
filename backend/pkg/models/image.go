package models

type Image struct {
	ImageID   int    `json:"image_id"`
	ImageURL  string `json:"image_url"`
	ImageName string `json:"image_name"`
}
