package utils

import "golang-beer-example/models"

func ResponseMessageSetup(code int, status bool, response []models.Beer) *models.APIResponse {
	return &models.APIResponse{
		Status: status,
		Code:   code,
		Data:   response,
	}
}
