package handler

import (
	"dot/app/models/dto"
	"dot/app/services"
	"dot/helpers"
	"net/http"
	"strconv"

	_ "dot/docs"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryService services.CategoryServices
}

func NewCatgoryHandler(categoryService services.CategoryServices) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

// GetAllCategory godoc
// @Summary Get all categories
// @Description Retrieve a list of all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} dto.MultipleCategoryResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /admin/category [get]
func (ch *CategoryHandler) GetAllCategory(c echo.Context) error {
	resp, err := ch.categoryService.GetAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.MultipleCategoryResponse{
		Message: "Select All Category",
		Data:    resp,
	})
}

// ShowCategory godoc
// @Summary Get a category by ID
// @Description Retrieve a single category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} dto.SingleCategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /admin/category/{id} [get]
func (ch *CategoryHandler) ShowCategory(c echo.Context) error {
	id := c.Param("id")
	iduint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "id must number",
		})
	}
	resp, error := ch.categoryService.Show(uint(iduint))
	if error != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "cannot get this data",
		})
	}
	return c.JSON(http.StatusOK, dto.SingleCategoryResponse{
		Message: "Category Show",
		Data: dto.CategoryResponse{
			ID:   resp.ID,
			Name: resp.Name,
		},
	})
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Add a new category to the database
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CategoryCreate true "Category data"
// @Success 200 {object} dto.SingleCategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /admin/category [post]
func (ch *CategoryHandler) CreateCategory(c echo.Context) error {
	var req dto.CategoryCreate
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := ch.categoryService.Create(req.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SingleCategoryResponse{
		Message: "Category Create",
		Data: dto.CategoryResponse{
			ID:   resp.ID,
			Name: resp.Name,
		},
	})
}

// UpdateCategory godoc
// @Summary Update Category
// @Description Update Category to database
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CategoryCreate true "Category data"
// @Param id path int true "Category ID"
// @Success 200 {object} dto.SingleCategoryResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /admin/category/:id [put]
func (ch *CategoryHandler) UpdateCategory(c echo.Context) error {
	var req dto.CategoryCreate
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid ID format",
		})
	}
	category, err := ch.categoryService.Update(uint(id), req.Name)
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to update category",
		})
	}
	return c.JSON(http.StatusOK, dto.SingleCategoryResponse{
		Message: "Category updated successfully",
		Data: dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		},
	})
}

// DeleteCategory godoc
// @Summary Delete Category
// @Description Delete Category to database
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Success 200 {object} dto.SingleCategoryResponse
// @Security BearerAuth
// @Router /admin/category/:id [delete]
func (ch *CategoryHandler) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := ch.categoryService.Delete(uint(id))
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to delete category",
		})
	}

	return c.JSON(http.StatusOK, dto.SingleCategoryResponse{
		Message: "Category deleted successfully",
		Data: dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		},
	})
}
