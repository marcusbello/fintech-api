package usecase_test

import (
	"errors"
	"fintech-api/pkg/domain"
	"fintech-api/pkg/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

func TestLoginUseCase(t *testing.T) {
	gin.SetMode("test")
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	var err error
	t.Run("Test Empty Login", func(t *testing.T) {
		mockFintechRepo := domain.NewMockFintechRepository(t)
		mockFintechRepo.On("LoginRepository", c, mock.Anything, mock.Anything).Return(fmt.Errorf("error: empty username and password %v", err)).Once()
		u := usecase.NewFintechUseCase(mockFintechRepo)
		err := u.LoginUc(c, "", "")
		assert.Error(t, err)
	})
	t.Run("Test Correct Password", func(t *testing.T) {
		mockFintechRepo := domain.NewMockFintechRepository(t)
		mockFintechRepo.On("LoginRepository", c, mock.Anything, mock.Anything).Return(nil).Once()
		u := usecase.NewFintechUseCase(mockFintechRepo)
		err := u.LoginUc(c, "juwon", "password123")
		assert.Nil(t, err)
	})
	t.Run("Test InCorrect Username", func(t *testing.T) {
		mockFintechRepo := domain.NewMockFintechRepository(t)
		mockFintechRepo.On("LoginRepository", c, mock.Anything, mock.Anything).Return(fmt.Errorf("error : %v", err)).Once()
		u := usecase.NewFintechUseCase(mockFintechRepo)
		err := u.LoginUc(c, "juwn", "password123")
		assert.NotNil(t, err)
	})
}

func TestRemoveMoneyUseCase(t *testing.T) {
	gin.SetMode("test")
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	var err error
	t.Run("Test Empty Details", func(t *testing.T) {
		mockFintechRepo := domain.NewMockFintechRepository(t)
		mockFintechRepo.On("RemoveMoneyRepository", c, mock.Anything, mock.Anything).Return(domain.Account{}, errors.New("error details empty")).Once()
		u := usecase.NewFintechUseCase(mockFintechRepo)
		uc, err := u.RemoveMoneyUc(c, "", 0)
		assert.NotNil(t, err)
		assert.NotNil(t, uc)
	})
	t.Run("Test Remove Money", func(t *testing.T) {
		mockFintechRepo := domain.NewMockFintechRepository(t)
		mockFintechRepo.On("RemoveMoneyRepository", c, mock.Anything, mock.Anything).Return(domain.Account{}, err).Once()
		u := usecase.NewFintechUseCase(mockFintechRepo)
		uc, err := u.RemoveMoneyUc(c, "marcus", 200)
		assert.Nil(t, err)
		assert.NotNil(t, uc)
	})
}
