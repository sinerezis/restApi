package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {

	// Создаем новый конфиг
	s := New(NewConfig())

	// Создаем рекордер и новый реквест
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	// Запускаем хэндлер и сравниваем его ответ
	// с тестовым значением
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello, rest api!")
}
