package advertisementClient

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	advertisementModels "api-tests-template/internal/managers/advertisement/models"
)

// HttpPostCreateAdvertisement отправляет multipart/form-data запрос для создания объявления
func HttpPostCreateAdvertisement(t *testing.T, token string, request advertisementModels.CreateAdvertisementRequest) *http.Response {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("title", request.Title)
	_ = writer.WriteField("description", request.Description)
	_ = writer.WriteField("price", strconv.Itoa(request.Price))
	_ = writer.WriteField("quantity", strconv.Itoa(request.Quantity))

	for i, photoContent := range request.PhotoFiles {
		part, err := writer.CreateFormFile("photos", fmt.Sprintf("photo%d.jpg", i+1))
		if err != nil {
			t.Fatalf("Не удалось создать form file: %s", err)
		}
		_, err = part.Write(photoContent)
		if err != nil {
			t.Fatalf("Не удалось записать содержимое файла: %s", err)
		}
	}

	writer.Close()

	return apiRunner.GetRunner().Auth(token).Create().
		Post(path.CreateAdvertisementPath).
		ContentType(writer.FormDataContentType()).
		Body(body.String()).
		Expect(t).
		End().Response
}

// HttpGetAdvertisement выполняет GET-запрос для получения объявления по id
func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	return apiRunner.GetRunner().Create().
		Get(path.GetAdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}

// HttpGetAdvertisements выполняет GET-запрос для получения списка объявлений с поиском
func HttpGetAdvertisements(t *testing.T, search string) *http.Response {
	runner := apiRunner.GetRunner().Create().Get(path.GetAdvertisementsPath)
	if search != "" {
		runner = runner.Query("search", search)
	}
	return runner.Expect(t).End().Response
}

// HttpGetAdvertisementPhotos выполняет GET-запрос для получения фотографий объявления
func HttpGetAdvertisementPhotos(t *testing.T, token string, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(fmt.Sprintf(path.GetAdvertisementPhotos, id)).
		Expect(t).
		End().Response
}

// HttpDeleteAdvertisement выполняет DELETE-запрос для удаления объявления по id
func HttpDeleteAdvertisement(t *testing.T, token string, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Delete(path.DeleteAdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}
