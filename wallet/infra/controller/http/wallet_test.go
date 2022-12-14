package http_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/app/usecase"
	"github.com/wellingtonlope/superfinance-api/wallet/infra/controller/http"
	httpGo "net/http"
	"net/http/httptest"
	"testing"
)

func TestWallet_GetByID(t *testing.T) {
	tt := []struct {
		name               string
		expectedStatusCode int
		expectedBody       string
		outputUC           usecase.WalletOutput
		errorUC            error
	}{
		{
			name:               "should get wallet by id",
			expectedStatusCode: httpGo.StatusOK,
			expectedBody:       `{"id":"1","name":"name","description":"description"}`,
			outputUC:           usecase.WalletOutput{ID: "1", Name: "name", Description: "description"},
			errorUC:            nil,
		},
		{
			name:               "should fail with not found when wallet not exists",
			expectedStatusCode: httpGo.StatusNotFound,
			expectedBody:       `{"message":"wallet not found"}`,
			outputUC:           usecase.WalletOutput{},
			errorUC:            repository.ErrWalletNotFound,
		},
		{
			name:               "should fail with internal error when error not mapped",
			expectedStatusCode: httpGo.StatusInternalServerError,
			expectedBody:       `{"message":"i'm an error"}`,
			outputUC:           usecase.WalletOutput{},
			errorUC:            errors.New("i'm an error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			createWalletUC := usecase.NewCreateWalletMock()
			getByIDWalletUC := usecase.NewGetByIDWalletMock()
			updateWalletUC := usecase.NewUpdateWalletMock()
			getByIDWalletUC.
				On("Handle", "1").
				Return(tc.outputUC, tc.errorUC).Once()
			contr := http.NewWallet(createWalletUC, getByIDWalletUC, updateWalletUC)

			router := gin.Default()
			router.GET("/wallet/:id", contr.GetByID)

			res := httptest.NewRecorder()
			req, _ := httpGo.NewRequest(httpGo.MethodGet, "/wallet/1", nil)
			router.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatusCode, res.Code)
			assert.Equal(t, tc.expectedBody, res.Body.String())
			getByIDWalletUC.AssertExpectations(t)
		})
	}
}
