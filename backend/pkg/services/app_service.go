package services

import "electron-app/backend/pkg/models"

func getAppByID(id int) models.App {
	return models.App{
		ID:  id,
		Name: "My App",
		URL:  "https://myapp.com",
	}
}