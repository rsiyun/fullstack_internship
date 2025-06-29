package handlers

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/services"
	"api-money-management/pkg/auth"
	"net/http"
	"strconv"

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

func (h *WalletHandler) ShowWallet(c echo.Context) error {
	walletId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid wallet ID",
			Code:    http.StatusBadRequest,
			Details: "Wallet ID must be a number",
		})
	}
	wallet, errDatabase := h.walletservice.GetWalletByID(walletId)
	if errDatabase != nil {
		return c.JSON(errDatabase.Code, errDatabase)
	}
	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Wallet retrieved successfully",
		Code:    http.StatusOK,
		Data:    wallet,
	})
}

func (h *WalletHandler) UpdateWallet(c echo.Context) error {
	walletId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid wallet ID",
			Code:    http.StatusBadRequest,
			Details: "Wallet ID must be a number",
		})
	}
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}
	req := new(dtos.RequestWallet)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}
	req.UserID = int(userID)
	errors := c.Validate(req)
	if errors != nil {
		err := dtos.NewValidationError(errors)
		return c.JSON(err.Code, err)
	}
	wallet := &models.Wallet{
		ID:      uint(walletId),
		UserId:  uint(req.UserID),
		Name:    req.Name,
		Balance: req.Balance,
	}
	result, errDatabase := h.walletservice.UpdateWallet(wallet)
	if errDatabase != nil {
		return c.JSON(errDatabase.Code, err)
	}
	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Wallet updated successfully",
		Code:    http.StatusOK,
		Data:    result,
	})
	// result, errordatabase := h.walletservice.(wallet)
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
	errvalidation := c.Validate(req)
	if errvalidation != nil {
		err := dtos.NewValidationError(errvalidation)
		return c.JSON(err.Code, err)
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
