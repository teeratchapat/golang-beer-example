package services

import (
	"golang-beer-example/database"
	"golang-beer-example/models"
	"golang-beer-example/modules/logic"
	"golang-beer-example/modules/repository"
	"golang-beer-example/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func ConnectionDB() (db *gorm.DB, mongo *mongo.Database) {
	mongoDB := database.GetMongoDbPool()
	mariaDB := database.ConnectMariaDB()

	return mariaDB, mongoDB
}

func BeerGetService(pagination logic.InputPagination) (*models.APIResponse, error) {
	mongoDB, mariaDB := ConnectionDB()
	connectRepo := repository.NewBeerRepository(mongoDB, mariaDB)
	beers, err := connectRepo.GetBeerRepository()

	result := logic.NumberBasePaginate(pagination, logic.MapDataListToBeerList(beers))
	if err != nil {
		return utils.ResponseMessageSetup(500, false, nil), err
	}

	return utils.ResponseMessageSetup(200, true, result), nil
}

func BeerPostService(data *models.Beer) (*models.APIResponse, error) {
	mongoDB, mariaDB := ConnectionDB()
	connectRepo := repository.NewBeerRepository(mongoDB, mariaDB)
	err := connectRepo.CreateBeerRepository(data)
	if err != nil {
		return utils.ResponseMessageSetup(500, false, nil), err
	}

	return utils.ResponseMessageSetup(200, true, nil), nil
}

func BeerPutService(id int, data *models.Beer) (*models.APIResponse, error) {
	mongoDB, mariaDB := ConnectionDB()
	connectRepo := repository.NewBeerRepository(mongoDB, mariaDB)
	err := connectRepo.UpdateBeerRepository(id, data)
	if err != nil {
		return utils.ResponseMessageSetup(500, false, nil), err
	}
	return utils.ResponseMessageSetup(200, true, nil), nil
}

func BeerDeleteService(id int) (*models.APIResponse, error) {
	mongoDB, mariaDB := ConnectionDB()
	connectRepo := repository.NewBeerRepository(mongoDB, mariaDB)
	err := connectRepo.DeleteBeerRepository(id)
	if err != nil {
		return utils.ResponseMessageSetup(500, false, nil), err
	}

	return utils.ResponseMessageSetup(200, true, nil), nil
}
