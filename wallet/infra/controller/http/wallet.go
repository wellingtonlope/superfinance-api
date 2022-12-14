package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/app/usecase"
	httpGo "net/http"
)

type WalletResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewWalletResponse(w usecase.WalletOutput) WalletResponse {
	return WalletResponse{
		ID:          w.ID,
		Name:        w.Name,
		Description: w.Description,
	}
}

type Wallet interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
}

type wallet struct {
	createWalletUC  usecase.CreateWallet
	getByIDWalletUC usecase.GeByIDWallet
	updateWallet    usecase.UpdateWallet
}

func NewWallet(createWalletUC usecase.CreateWallet, getByIDWalletUC usecase.GeByIDWallet,
	updateWallet usecase.UpdateWallet) Wallet {
	return &wallet{createWalletUC, getByIDWalletUC, updateWallet}
}

func (h *wallet) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *wallet) GetByID(c *gin.Context) {
	walletOutput, err := h.getByIDWalletUC.Handle(c.Param("id"))
	if err != nil {
		var status int
		switch err {
		case repository.ErrWalletNotFound:
			status = httpGo.StatusNotFound
		default:
			status = httpGo.StatusInternalServerError
		}
		c.JSON(status, NewErrorResponse(err))
		return
	}
	c.JSON(200, NewWalletResponse(walletOutput))
}

func (h *wallet) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
