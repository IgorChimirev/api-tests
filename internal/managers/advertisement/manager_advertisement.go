package advertisement

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	advertisementClient "api-tests-template/internal/client/http/advertisement"
	photoCheckClient "api-tests-template/internal/client/http/photoCheck"
	advertisementModels "api-tests-template/internal/managers/advertisement/models"
)

// CreateAdvertisement создаёт объявление и возвращает тело ответа в виде строки
func CreateAdvertisement(t *testing.T, token string, request advertisementModels.CreateAdvertisementRequest, expectedStatusCode int) string {
	resp := advertisementClient.HttpPostCreateAdvertisement(t, token, request)
	defer resp.Body.Close()

	require.Equalf(t, expectedStatusCode, resp.StatusCode,
		"HTTP status code должен быть %d, получили %d", expectedStatusCode, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return string(body)
}

// CreateAdvertisementExpectError создаёт объявление и проверяет, что статус-код входит в список допустимых.
func CreateAdvertisementExpectError(t *testing.T, token string, request advertisementModels.CreateAdvertisementRequest, acceptableCodes ...int) string {
	resp := advertisementClient.HttpPostCreateAdvertisement(t, token, request)
	defer resp.Body.Close()

	require.Containsf(t, acceptableCodes, resp.StatusCode,
		"HTTP status code должен быть одним из %v, получили %d", acceptableCodes, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return string(body)
}

// GetAdvertisement получает объявление по id и возвращает тело ответа
func GetAdvertisement(t *testing.T, id string, expectedStatusCode int) string {
	resp := advertisementClient.HttpGetAdvertisement(t, id)
	defer resp.Body.Close()
	require.Equalf(t, expectedStatusCode, resp.StatusCode,
		"HTTP status code должен быть %d, получили %d", expectedStatusCode, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(body)
}

// GetAdvertisements получает список объявлений с поиском и возвращает тело ответа
func GetAdvertisements(t *testing.T, search string, expectedStatusCode int) string {
	resp := advertisementClient.HttpGetAdvertisements(t, search)
	defer resp.Body.Close()
	require.Equalf(t, expectedStatusCode, resp.StatusCode,
		"HTTP status code должен быть %d, получили %d", expectedStatusCode, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(body)
}

// GetAdvertisementPhotos получает фотографии объявления и возвращает тело ответа
func GetAdvertisementPhotos(t *testing.T, token string, id string, expectedStatusCode int) string {
	resp := advertisementClient.HttpGetAdvertisementPhotos(t, token, id)
	defer resp.Body.Close()
	require.Equalf(t, expectedStatusCode, resp.StatusCode,
		"HTTP status code должен быть %d, получили %d", expectedStatusCode, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(body)
}

// CheckPhotoAvailableByURL проверяет что фотография доступна по URL через HEAD-запрос к серверу хранилища
func CheckPhotoAvailableByURL(t *testing.T, url string) {
	resp := photoCheckClient.HttpHeadPhotoURL(t, url)
	defer resp.Body.Close()
	require.Equalf(t, http.StatusOK, resp.StatusCode,
		"Фото по URL %s должно быть доступно на сервере, получили статус %d", url, resp.StatusCode)
}

// DeleteAdvertisement удаляет объявление по id

func DeleteAdvertisement(t *testing.T, token string, id string, expectedStatusCode int) string {
	resp := advertisementClient.HttpDeleteAdvertisement(t, token, id)
	defer resp.Body.Close()
	require.Equal(t, expectedStatusCode, resp.StatusCode,
		"HTTP status code должен быть %d, получили %d", expectedStatusCode, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(body)
}
