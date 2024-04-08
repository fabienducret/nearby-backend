package news

import (
	"nearby/models"
)

func NewsRepositoryFactory() func(city string) ([]models.News, error) {
	return func(city string) ([]models.News, error) {
		return []models.News{
			{Title: "titre 1", Description: "desc 1"},
			{Title: "titre 2", Description: "desc 2"},
		}, nil
	}
}
