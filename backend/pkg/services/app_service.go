package services

import (
	"database/sql"
	"fmt"

	"github.com/EmilSunden/electron-app/backend/pkg/models"
)

type AppService struct {
	db *sql.DB
}

func NewAppService(db *sql.DB) *AppService {
	return &AppService{
		db: db,
	}
}

func (s *AppService) GetApps() ([]models.App, error) {
	query := `SELECT a.app_id, a.app_name, a.app_url, i.image_id, i.image_url, i.image_name
	FROM 
		electron_app_schema.application_table a 
	INNER JOIN 
		electron_app_schema.image_table i ON a.image_id = i.image_id
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []models.App
	for rows.Next() {
		var app models.App
		var image models.Image

		err := rows.Scan(
			&app.AppID, &app.AppName, &app.AppURL,
			&image.ImageID, &image.ImageURL, &image.ImageName,
		)

		if err != nil {
			return nil, err
		}

		app.Image = &image

		apps = append(apps, app)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return apps, nil
}

func (s *AppService) CreateApp(app *models.App) error {
	// Attempt to find the image_id based on the app_name
	var imageID sql.NullInt64
	imageQuery := `SELECT image_id FROM image_table WHERE LOWER(image_name) = LOWER($1)`
	err := s.db.QueryRow(imageQuery, app.AppName).Scan(&imageID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to query image: %v", err)
	}

	// Insert the app with the found image_id (which may be NULL)
	appQuery := `
		INSERT INTO application_table (app_name, app_url, image_id)
		VALUES ($1, $2, $3) RETURNING app_id
	`

	var imageIDValue interface{}
	if imageID.Valid {
		imageIDValue = imageID.Int64
	} else {
		imageIDValue = nil
	}

	err = s.db.QueryRow(appQuery, app.AppName, app.AppURL, imageIDValue).Scan(&app.AppID)
	if err != nil {
		return fmt.Errorf("failed to insert app: %v", err)
	}

	return nil
}
