package http_test

import (
	fintechhandler "fintech-api/pkg/delivery/http"
	"fintech-api/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFintechHandler_Login(t *testing.T) {
	t.Run("Test Login", func(t *testing.T) {
		_, engine := gin.CreateTestContext(httptest.NewRecorder())
		mockFintechLoginUseCase := domain.NewMockFintechUseCase(t)
		mockFintechLoginUseCase.On("LoginUc", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		fintechhandler.NewFintechHandler(engine, mockFintechLoginUseCase)
		req, err := http.NewRequest(http.MethodPost, "/api/signin", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Test Login With Empty Data", func(t *testing.T) {
		_, engine := gin.CreateTestContext(httptest.NewRecorder())
		mockFintechLoginUseCase := domain.NewMockFintechUseCase(t)
		mockFintechLoginUseCase.On("LoginUc", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		fintechhandler.NewFintechHandler(engine, mockFintechLoginUseCase)
		req, err := http.NewRequest(http.MethodPost, "/api/signin", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
