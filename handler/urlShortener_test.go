package handler

import (
	"URL-shortener/service"
	mock_service "URL-shortener/service/mocks"
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestGetByAlias(t *testing.T) {

	type mockBehaviour func(s *mock_service.MockURLService, alias string)

	type tc struct {
		name               string
		inputBody          string
		mockBehaviour      mockBehaviour
		expectedStatusCode int
	}

	tests := []tc{
		{
			name:      "Valid",
			inputBody: "oSmFKu0G",
			mockBehaviour: func(s *mock_service.MockURLService, alias string) {
				s.EXPECT().GetByAlias(alias).Return(`http://www.google.com/?q=golang`, nil)
			},
			expectedStatusCode: 302,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Инициализация зависимостей
			c := gomock.NewController(t)
			defer c.Finish()

			receiveURL := mock_service.NewMockURLService(c)
			test.mockBehaviour(receiveURL, test.inputBody)

			services := &service.Service{URLService: receiveURL}
			handler := NewHandler(services)

			//Тестовый сервер
			r := gin.New()
			r.GET("/s/:alias", handler.getUrl)

			//Тестовый запрос
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/s/oSmFKu0G", bytes.NewBufferString(test.inputBody))

			//Запрос
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
		})
	}
}

func TestSaveByURL(t *testing.T) {

	type mockBehaviour func(s *mock_service.MockURLService)

	type tc struct {
		name               string
		mockBehaviour      mockBehaviour
		expectedStatusCode int
		expectedBody       string
	}

	tests := []tc{
		{
			name: "Valid",
			mockBehaviour: func(s *mock_service.MockURLService) {
				s.EXPECT().SaveURL("http://www.google.com/?q=golang").Return(`oSmFKu0G`, nil)
			},
			expectedStatusCode: 200,
			expectedBody:       `{"alias":"oSmFKu0G"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Инициализация зависимостей
			c := gomock.NewController(t)
			defer c.Finish()

			receiveAlias := mock_service.NewMockURLService(c)
			test.mockBehaviour(receiveAlias)

			services := &service.Service{URLService: receiveAlias}
			handler := NewHandler(services)

			//Тестовый сервер
			r := gin.New()
			r.POST("/a/", handler.createUrl)

			//Тестовый запрос
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/a/?url=http://www.google.com/?q=golang", bytes.NewBufferString(""))

			//Запрос
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedBody, w.Body.String())
		})
	}
}
