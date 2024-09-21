package models

type App struct {
	AppID   int    `json:"app_id"`
	AppName string `json:"app_name"`
	AppURL  string `json:"app_url"`
	Image   *Image `json:"image"` // Pointer to allow nil
}
