package handlers

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/services"
	"api-money-management/pkg/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WalletHandler struct {
	walletservice *services.WalletService
}

func NewWalletHandler(walletservice *services.WalletService) *WalletHandler {
	return &WalletHandler{
		walletservice: walletservice,
	}
}
func (h *WalletHandler) GetWalletUser(c echo.Context) error {
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}
	response, err := h.walletservice.GetWalletByUserID(int(userID))
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "success to get data",
		Code:    http.StatusOK,
		Data:    response,
	})
}

func (h *WalletHandler) CreateWallet(c echo.Context) error {
	// ambil terlebih dahulu data dari jwt token
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}

	req := new(dtos.RequestWallet)

	// Binding
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}
	req.UserID = int(userID)

	// // validasi request terlebih dahulu
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "validation failed",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}
	wallet := &models.Wallet{
		UserId:  uint(req.UserID),
		Name:    req.Name,
		Balance: req.Balance,
	}
	result, errordatabase := h.walletservice.CreateWallet(wallet)
	if errordatabase != nil {
		return c.JSON(errordatabase.Code, dtos.ErrorResponse{
			Message: errordatabase.Message,
			Code:    errordatabase.Code,
			Details: errordatabase.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.SuccessResponse{
		Message: "created successfully",
		Code:    http.StatusCreated,
		Data:    result,
	})

}
