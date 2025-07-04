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

type CategoryExpenseHandler struct {
	categoryExpenseService *services.CategoryExpenseService
}

func NewCategoryExpenseHandler(categoryExpenseService *services.CategoryExpenseService) *CategoryExpenseHandler {
	return &CategoryExpenseHandler{
		categoryExpenseService: categoryExpenseService,
	}
}

func (h *CategoryExpenseHandler) GetCategoryExpense(c echo.Context) error {
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}

	response, err := h.categoryExpenseService.GetCategoryExpenseByUserId(int(userID))
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Categories retrieved successfully",
		Code:    http.StatusOK,
		Data:    response,
	})
}

func (h *CategoryExpenseHandler) ShowCategoryExpense(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid category ID",
			Code:    http.StatusBadRequest,
			Details: "Category ID must be a number",
		})
	}

	category, errdatabase := h.categoryExpenseService.GetCategoryExpenseById(categoryID)
	if errdatabase != nil {
		return c.JSON(errdatabase.Code, err)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Category retrieved successfully",
		Code:    http.StatusOK,
		Data:    category,
	})
}

func (h *CategoryExpenseHandler) CreateCategoryExpense(c echo.Context) error {
	userID, errorIDToken := auth.GetUserIDFromToken(c)
	if errorIDToken != nil {
		return c.JSON(errorIDToken.Code, errorIDToken)
	}

	req := new(dtos.RequestCategory)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewValidationError(err))
	}

	category := &models.ExpenseCategory{
		UserId: userID,
		Name:   req.Name,
		Icon:   req.Icon,
		Color:  req.Color,
	}

	result, err := h.categoryExpenseService.CreateCategoryExpense(category)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusCreated, dtos.SuccessResponse{
		Message: "Category created successfully",
		Code:    http.StatusCreated,
		Data:    result,
	})
}

func (h *CategoryExpenseHandler) UpdateCategoryExpense(c echo.Context) error {
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

	req := new(dtos.RequestCategory)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.NewValidationError(err))
	}

	category := &models.ExpenseCategory{
		ID:     uint(categoryID),
		UserId: userID,
		Name:   req.Name,
		Icon:   req.Icon,
		Color:  req.Color,
	}

	result, errdatabase := h.categoryExpenseService.UpdateCategoryExpense(category)
	if errdatabase != nil {
		return c.JSON(errdatabase.Code, err)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResponse{
		Message: "Category updated successfully",
		Code:    http.StatusOK,
		Data:    result,
	})
}
