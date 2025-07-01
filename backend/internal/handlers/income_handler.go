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

type IncomeHandler struct {
	incomeService *services.IncomeService
}

func NewIncomeHandler(incomeService *services.IncomeService) *IncomeHandler {
	return &IncomeHandler{
		incomeService: incomeService,
	}
}

func (h *IncomeHandler) GetIncomes(c echo.Context) error {
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}

	response, err := h.incomeService.GetIncomeByUserId(int(userID))
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Income list retrieved successfully",
		Code:    http.StatusOK,
		Data:    response,
	})
}

func (h *IncomeHandler) ShowIncome(c echo.Context) error {
	// get the id from routes
	incomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid wallet ID",
			Code:    http.StatusBadRequest,
			Details: "Wallet ID must be a number",
		})
	}
	result, errdatabase := h.incomeService.GetIncomeIncomeById(incomeID)
	if errdatabase != nil {
		return c.JSON(http.StatusBadRequest, errdatabase)
	}
	// call the services
	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Income retrieved successfully",
		Code:    http.StatusOK,
		Data:    result,
	})

}

func (h *IncomeHandler) CreateIncome(c echo.Context) error {
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}
	req := new(dtos.RequestIncome)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}
	req.UserId = int(userID)
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewValidationError(err))
	}
	income := &models.Income{
		UserId:           uint(req.UserId),
		WalletId:         uint(req.WalletId),
		IncomeCategoryId: uint(req.IncomeCategoryId),
		Amount:           req.Amount,
		Description:      req.Description,
	}
	result, err := h.incomeService.CreateIncome(income)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusCreated, dtos.SuccessResponse{
		Message: "Income created successfully",
		Code:    http.StatusCreated,
		Data:    result,
	})
}

func (h *IncomeHandler) UpdateIncome(c echo.Context) error {
	incomeID, err := strconv.Atoi(c.Param("id"))
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
	req := new(dtos.RequestIncome)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}
	req.UserId = int(userID)
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewValidationError(err))
	}
	income := &models.Income{
		ID:               uint(incomeID),
		WalletId:         uint(req.WalletId),
		IncomeCategoryId: uint(req.IncomeCategoryId),
		UserId:           userID,
		Amount:           req.Amount,
		Description:      req.Description,
	}
	result, errDatabase := h.incomeService.UpdateIncome(income)
	if errDatabase != nil {
		return c.JSON(errDatabase.Code, errDatabase)
	}
	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Income updated successfully",
		Code:    http.StatusOK,
		Data:    result,
	})
}

// when delete
