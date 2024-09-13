package services

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/EmilSunden/electron-app/backend/pkg/models"
)

type AppService struct {
	filePath string // TODO: change property to db
	Apps     []models.App
}

func NewAppService(filePath string) *AppService {
	service := &AppService{
		filePath: filePath, // TODO: Change value to db
		Apps:     []models.App{},
	}

	service.loadAppsFromFile()
	return service
}

// loadAppsFromFile reads the apps from the JSON file
func (s *AppService) loadAppsFromFile() error {
	file, err := ioutil.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No file exists yet, this is fine
		}
		return err
	}

	err = json.Unmarshal(file, &s.Apps)
	if err != nil {
		return err
	}

	return nil
}

func (s *AppService) saveAppsToFile() error {
	// Save apps to file
	data, err := json.MarshalIndent(s.Apps, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(s.filePath, data, 0644)
}

func (s *AppService) GetApps() []models.App {
	return s.Apps
}

func (s *AppService) CreateApp(app models.App) error {
	s.Apps = append(s.Apps, app)
	return s.saveAppsToFile()
}
