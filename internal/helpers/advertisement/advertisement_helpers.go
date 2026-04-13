package advertisementHelpers

import (
	advertisementModels "api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/utils"
)

// DefaultCreateRequest возвращает стандартный запрос на создание объявления с рандомными данными.
func DefaultCreateRequest(photos [][]byte) advertisementModels.CreateAdvertisementRequest {
	return advertisementModels.CreateAdvertisementRequest{
		Title:       "Test " + utils.RandomString(8),
		Description: "Description " + utils.RandomString(16),
		Price:       1000,
		Quantity:    5,
		PhotoFiles:  photos,
	}
}
