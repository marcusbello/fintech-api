package http_test

import (
	"bytes"
	"encoding/json"
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
	gin.SetMode("test")
	t.Run("Test Login", func(t *testing.T) {

		payload := domain.LoginRequest{
			UserName: "marcus",
			Password: "password12",
		}
		// Encode the payload as JSON
		payloadBuf := new(bytes.Buffer)
		err := json.NewEncoder(payloadBuf).Encode(&payload)
		assert.NoError(t, err)
		//jsonPayload, _ := json.Marshal(payload)
		_, engine := gin.CreateTestContext(httptest.NewRecorder())
		mockFintechLoginUseCase := domain.NewMockFintechUseCase(t)
		mockFintechLoginUseCase.On("LoginUc", mock.Anything, payload.UserName, payload.Password).Return(nil)
		fintechhandler.NewFintechHandler(engine, mockFintechLoginUseCase)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/signin", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Test Login With Empty Data", func(t *testing.T) {
		var err error
		payload := domain.LoginRequest{
			UserName: "",
			Password: "",
		}
		// Encode the payload as JSON
		payloadBuf := new(bytes.Buffer)
		err = json.NewEncoder(payloadBuf).Encode(&payload)
		assert.NoError(t, err)
		//_, engine := gin.CreateTestContext(httptest.NewRecorder())
		//mockFintechLoginUseCase := domain.NewMockFintechUseCase(t)
		//mockFintechLoginUseCase.On("LoginUc", mock.Anything, payload.UserName, payload.Password).Return(err)
		//fintechhandler.NewFintechHandler(engine, mockFintechLoginUseCase)
		//req, err := http.NewRequest(http.MethodPost, "/api/v1/signin", payloadBuf)
		//req.Header.Set("Content-Type", "application/json")
		//assert.NoError(t, err)
		//w := httptest.NewRecorder()
		//engine.ServeHTTP(w, req)
		//assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
