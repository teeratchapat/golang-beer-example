package repository

import (
	"context"
	"golang-beer-example/models"
	repoModel "golang-beer-example/modules/models"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type BeerRepository interface {
	GetBeerRepository() ([]models.Beer, error)
	CreateBeerRepository(data *models.Beer) error
	UpdateBeerRepository(id int, data *models.Beer) error
	DeleteBeerRepository(id int) error
}

type beerRepository struct {
	db        *gorm.DB
	mongoColl *mongo.Collection
}

func NewBeerRepository(db *gorm.DB, mongoDB *mongo.Database) BeerRepository {
	return &beerRepository{
		db:        db,
		mongoColl: mongoDB.Collection("logs"),
	}
}

func (r *beerRepository) GetBeerRepository() ([]models.Beer, error) {
	var data []models.Beer
	if err := r.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *beerRepository) CreateBeerRepository(data *models.Beer) error {
	if err := r.db.Create(data).Error; err != nil {
		return err
	}
	logging := repoModel.Beer{
		ID:        uint(data.ID),
		Name:      data.Name,
		Type:      "",
		Details:   "",
		ImageURL:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	logEntry := repoModel.LogEntry{
		Action:    "create",
		Timestamp: time.Now(),
		Data:      logging,
	}
	_, err := r.mongoColl.InsertOne(context.Background(), logEntry)
	return err
}

func (r *beerRepository) UpdateBeerRepository(id int, data *models.Beer) error {
	if err := r.db.Save(data).Error; err != nil {
		return err
	}

	logging := repoModel.Beer{
		ID:        uint(data.ID),
		Name:      data.Name,
		Type:      "",
		Details:   "",
		ImageURL:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	logEntry := repoModel.LogEntry{
		Action:    "update",
		Timestamp: time.Now(),
		Data:      logging,
	}
	_, err := r.mongoColl.InsertOne(context.Background(), logEntry)
	return err
}

func (r *beerRepository) DeleteBeerRepository(id int) error {
	if err := r.db.Delete(&models.Beer{}, id).Error; err != nil {
		return err
	}
	logging := repoModel.Beer{
		ID: uint(id),
	}
	logEntry := repoModel.LogEntry{
		Action:    "delete",
		Timestamp: time.Now(),
		Data:      logging,
	}
	_, err := r.mongoColl.InsertOne(context.Background(), logEntry)
	return err
}
