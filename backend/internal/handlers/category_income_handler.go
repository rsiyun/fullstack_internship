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

type CategoryIncomeHandler struct {
	categoryIncomeService *services.CategoryIncomeService
}

func NewCategoryIncomeHandler(categoryIncomeService *services.CategoryIncomeService) *CategoryIncomeHandler {
	return &CategoryIncomeHandler{
		categoryIncomeService: categoryIncomeService,
	}
}

func (h *CategoryIncomeHandler) GetCategoryIncome(c echo.Context) error {
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}

	response, err := h.categoryIncomeService.GetCategoryIncomeByUserId(int(userID))
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Categories retrieved successfully",
		Code:    http.StatusOK,
		Data:    response,
	})
}

func (h *CategoryIncomeHandler) ShowCategoryIncome(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid category ID",
			Code:    http.StatusBadRequest,
			Details: "Category ID must be a number",
		})
	}

	category, errdatabase := h.categoryIncomeService.GetCategoryIncomeById(categoryID)
	if errdatabase != nil {
		return c.JSON(errdatabase.Code, err)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Category retrieved successfully",
		Code:    http.StatusOK,
		Data:    category,
	})
}

func (h *CategoryIncomeHandler) CreateCategoryIncome(c echo.Context) error {
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}

	req := new(dtos.RequestCategoryIncome)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}

	req.UserID = int(userID)
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewValidationError(err))
	}

	category := &models.IncomeCategory{
		UserId: uint(req.UserID),
		Name:   req.Name,
		Icon:   req.Icon,
		Color:  req.Color,
	}

	result, err := h.categoryIncomeService.CreateCategoryIncome(category)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusCreated, dtos.SuccessResponse{
		Message: "Category created successfully",
		Code:    http.StatusCreated,
		Data:    result,
	})
}

func (h *CategoryIncomeHandler) UpdateCategoryIncome(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid category ID",
			Code:    http.StatusBadRequest,
			Details: "Category ID must be a number",
		})
	}

	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}

	req := new(dtos.RequestCategoryIncome)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}

	req.UserID = int(userID)
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewValidationError(err))
	}

	category := &models.IncomeCategory{
		ID:     uint(categoryID),
		UserId: uint(req.UserID),
		Name:   req.Name,
		Icon:   req.Icon,
		Color:  req.Color,
	}

	result, errdatabase := h.categoryIncomeService.UpdateCategoryIncome(category)
	if errdatabase != nil {
		return c.JSON(errdatabase.Code, err)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Category updated successfully",
		Code:    http.StatusOK,
		Data:    result,
	})
}
