package advertisementTest

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	advertisementHelpers "api-tests-template/internal/helpers/advertisement"
	managerAdv "api-tests-template/internal/managers/advertisement"
	advertisementModels "api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData   authModels.LoginOkResponse
	createdAdID string
	photos      [][]byte
}

type searchResponse struct {
	Items []advertisementModels.AdvertisementResponse `json:"items"`
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

// readTestPhoto читает файл фотографии из папки testdata по имени файла
func readTestPhoto(t *testing.T, filename string) []byte {
	data, err := os.ReadFile("testdata/" + filename)
	require.NoError(t, err, "Не удалось прочитать файл: "+filename)
	return data
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация пользователя с кредами из переменных окружения")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	base.Precondition("Чтение тестовых фотографий из папки testdata")
	entries, err := os.ReadDir("testdata")
	require.NoError(s.T(), err, "Не удалось прочитать папку testdata")

	for _, entry := range entries {
		if !entry.IsDir() {
			photo := readTestPhoto(s.T(), entry.Name())
			s.photos = append(s.photos, photo)
		}
	}

	require.NotEmpty(s.T(), s.photos, "В папке testdata должна быть хотя бы одна фотография")
}

// TearDownSuite выполняется один раз после всех тестов — удаляем созданные данные
func (s *TestSuite) TearDownSuite() {
	if s.createdAdID != "" {
		s.Run("Удаляем созданное объявление после тестов", func() {
			managerAdv.DeleteAdvertisement(s.T(), s.loginData.Token, s.createdAdID, http.StatusNoContent)
		})
	}
	base.TearDownSuite()
}

func (s *TestSuite) TestCreateAdvertisementAllFields() {
	request := advertisementHelpers.DefaultCreateRequest(s.photos)

	var created advertisementModels.AdvertisementResponse

	s.Run("Шаг 1: создаём объявление со всеми полями", func() {
		body := managerAdv.CreateAdvertisement(s.T(), s.loginData.Token, request, http.StatusCreated)
		s.T().Log("Response body:", body)

		err := json.Unmarshal([]byte(body), &created)
		require.NoError(s.T(), err, "Ответ должен быть валидным JSON")
		require.NotEmpty(s.T(), created.ID, "ID созданного объявления не должен быть пустым")
		s.createdAdID = created.ID
	})

	s.Run("Шаг 2: все поля ответа соответствуют переданным значениям", func() {
		require.Equal(s.T(), request.Title, created.Title, "Title должен совпадать")
		require.Equal(s.T(), request.Description, created.Description, "Description должен совпадать")
		require.Equal(s.T(), request.Price, created.Price, "Price должен совпадать")
		require.Equal(s.T(), request.Quantity, created.Quantity, "Quantity должен совпадать")
		require.Lenf(s.T(), created.Photos, len(s.photos),
			"Количество фото должно совпадать: ожидали %d, получили %d", len(s.photos), len(created.Photos))
	})

	s.Run("Шаг 3: GET /advertisement?id={id} возвращает объект, совпадающий с созданным", func() {
		body := managerAdv.GetAdvertisement(s.T(), created.ID, http.StatusOK)

		var fetched advertisementModels.AdvertisementResponse
		err := json.Unmarshal([]byte(body), &fetched)
		require.NoError(s.T(), err, "GET-ответ должен быть валидным JSON")

		require.Equal(s.T(), created.ID, fetched.ID, "ID должен совпадать")
		require.Equal(s.T(), request.Title, fetched.Title, "Title должен совпадать")
		require.Equal(s.T(), request.Description, fetched.Description, "Description должен совпадать")
		require.Equal(s.T(), request.Price, fetched.Price, "Price должен совпадать")
		require.Equal(s.T(), request.Quantity, fetched.Quantity, "Quantity должен совпадать")
		require.Lenf(s.T(), fetched.Photos, len(s.photos),
			"Количество фото должно совпадать: ожидали %d, получили %d", len(s.photos), len(fetched.Photos))
	})

	s.Run("Шаг 4: GET /advertisements/{id}/photos возвращает все фото и они доступны на сервере", func() {
		body := managerAdv.GetAdvertisementPhotos(s.T(), s.loginData.Token, created.ID, http.StatusOK)

		var photos []advertisementModels.Photo
		err := json.Unmarshal([]byte(body), &photos)
		require.NoError(s.T(), err, "GET-ответ должен быть валидным JSON")
		require.Lenf(s.T(), photos, len(s.photos),
			"Количество фото должно совпадать: ожидали %d, получили %d", len(s.photos), len(photos))

		for _, photo := range photos {
			managerAdv.CheckPhotoAvailableByURL(s.T(), photo.URL)
		}
	})

	s.Run("Шаг 5: GET /advertisements?search=... — объявление находится в поиске по полному названию", func() {
		body := managerAdv.GetAdvertisements(s.T(), request.Title, http.StatusOK)

		var result searchResponse
		err := json.Unmarshal([]byte(body), &result)
		require.NoError(s.T(), err, "GET-ответ должен быть валидным JSON")

		var foundItem advertisementModels.AdvertisementResponse
		var found bool
		for _, ad := range result.Items {
			if ad.ID == created.ID {
				foundItem = ad
				found = true
				break
			}
		}

		require.Truef(s.T(), found,
			"Созданное объявление с id=%s должно быть найдено в поиске по названию '%s'",
			created.ID, request.Title)
		require.Equal(s.T(), request.Title, foundItem.Title, "Title найденного объявления должен совпадать")
		require.Equal(s.T(), request.Description, foundItem.Description, "Description найденного объявления должен совпадать")
		require.Equal(s.T(), request.Price, foundItem.Price, "Price найденного объявления должен совпадать")
		require.Equal(s.T(), request.Quantity, foundItem.Quantity, "Quantity найденного объявления должен совпадать")
	})
}

func (s *TestSuite) TestCreateAdvertisementWithoutToken() {
	request := advertisementHelpers.DefaultCreateRequest(s.photos)

	var body string

	s.Run("Создаём объявление без токена авторизации", func() {
		body = managerAdv.CreateAdvertisement(s.T(), "", request, http.StatusUnauthorized)
	})

	s.Run("Проверяем сообщение об ошибке", func() {
		require.Equal(s.T(), "unauthorized", gjson.Get(body, "error").String(),
			"Поле error должно содержать unauthorized")
		require.NotEmpty(s.T(), gjson.Get(body, "message").String(),
			"Поле message не должно быть пустым")
	})
}

func (s *TestSuite) TestCreateAdvertisementWithInvalidToken() {
	request := advertisementHelpers.DefaultCreateRequest(s.photos)

	var body string

	s.Run("Создаём объявление с невалидным токеном", func() {
		body = managerAdv.CreateAdvertisement(s.T(), "invalid_token", request, http.StatusUnauthorized)
	})

	s.Run("Проверяем сообщение об ошибке", func() {
		require.Equal(s.T(), "unauthorized", gjson.Get(body, "error").String(),
			"Поле error должно содержать unauthorized")
		require.Equal(s.T(), "Invalid or expired token", gjson.Get(body, "message").String(),
			"Поле message должно содержать корректное сообщение")
	})
}

func (s *TestSuite) TestCreateAdvertisementWithoutTitle() {
	request := advertisementHelpers.DefaultCreateRequest(s.photos)
	request.Title = ""

	var body string

	s.Run("Создаём объявление без title", func() {
		body = managerAdv.CreateAdvertisementExpectError(s.T(), s.loginData.Token, request,
			http.StatusBadRequest, http.StatusInternalServerError)
	})

	s.Run("Проверяем наличие ошибки в ответе", func() {
		require.NotEmpty(s.T(), gjson.Get(body, "error").String(),
			"Поле error не должно быть пустым")
		require.NotEmpty(s.T(), gjson.Get(body, "message").String(),
			"Поле message не должно быть пустым")
	})
}

func (s *TestSuite) TestCreateAdvertisementWithoutDescription() {
	request := advertisementHelpers.DefaultCreateRequest(s.photos)
	request.Description = ""

	var body string

	s.Run("Создаём объявление без description", func() {
		body = managerAdv.CreateAdvertisementExpectError(s.T(), s.loginData.Token, request,
			http.StatusBadRequest, http.StatusInternalServerError)
	})

	s.Run("Проверяем наличие ошибки в ответе", func() {
		require.NotEmpty(s.T(), gjson.Get(body, "error").String(),
			"Поле error не должно быть пустым")
		require.NotEmpty(s.T(), gjson.Get(body, "message").String(),
			"Поле message не должно быть пустым")
	})
}

func (s *TestSuite) TestCreateAdvertisementWithNegativePrice() {
	request := advertisementHelpers.DefaultCreateRequest(s.photos)
	request.Price = -1

	var body string

	s.Run("Создаём объявление с отрицательной ценой", func() {
		body = managerAdv.CreateAdvertisementExpectError(s.T(), s.loginData.Token, request,
			http.StatusBadRequest, http.StatusInternalServerError)
	})

	s.Run("Проверяем наличие ошибки в ответе", func() {
		require.NotEmpty(s.T(), gjson.Get(body, "error").String(),
			"Поле error не должно быть пустым")
	})
}

func (s *TestSuite) TestCreateAdvertisementWithNegativeQuantity() {
	request := advertisementHelpers.DefaultCreateRequest(s.photos)
	request.Quantity = -1

	var body string

	s.Run("Создаём объявление с отрицательным quantity", func() {
		body = managerAdv.CreateAdvertisementExpectError(s.T(), s.loginData.Token, request,
			http.StatusBadRequest, http.StatusInternalServerError)
	})

	s.Run("Проверяем наличие ошибки в ответе", func() {
		require.NotEmpty(s.T(), gjson.Get(body, "error").String(),
			"Поле error не должно быть пустым")
	})
}

func (s *TestSuite) TestCreateAdvertisementWithoutPhotos() {
	request := advertisementHelpers.DefaultCreateRequest(nil)

	var body string

	s.Run("Создаём объявление без фото", func() {
		body = managerAdv.CreateAdvertisementExpectError(s.T(), s.loginData.Token, request,
			http.StatusBadRequest, http.StatusInternalServerError)
	})

	s.Run("Проверяем наличие ошибки в ответе", func() {
		require.NotEmpty(s.T(), gjson.Get(body, "error").String(),
			"Поле error не должно быть пустым")
		require.NotEmpty(s.T(), gjson.Get(body, "message").String(),
			"Поле message не должно быть пустым")
	})
}

func (s *TestSuite) TestGetAdvertisementNotFound() {
	nonExistentID := "00000000-0000-0000-0000-000000000000"

	var body string

	s.Run("Запрашиваем объявление по несуществующему id", func() {
		body = managerAdv.GetAdvertisement(s.T(), nonExistentID, http.StatusNotFound)
	})

	s.Run("Проверяем сообщение об ошибке", func() {
		require.NotEmpty(s.T(), gjson.Get(body, "error").String(),
			"Поле error не должно быть пустым")
		require.NotEmpty(s.T(), gjson.Get(body, "message").String(),
			"Поле message не должно быть пустым")
	})
}
