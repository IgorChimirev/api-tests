package photoCheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// HttpHeadPhotoURL выполняет HEAD-запрос по URL фотографии для проверки её доступности на сервере
func HttpHeadPhotoURL(t *testing.T, url string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodHead, url, nil)
	require.NoError(t, err, "Не удалось создать HEAD-запрос для URL: "+url)

	resp, err := client.Do(req)
	require.NoError(t, err, "Не удалось выполнить HEAD-запрос к URL: "+url)

	return resp
}
