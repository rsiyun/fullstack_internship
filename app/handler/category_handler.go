package handler

import (
	"dot/app/models/dto"
	"dot/app/services"
	"dot/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryService services.CategoryServices
}

func NewCatgoryHandler(categoryService services.CategoryServices) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (ch *CategoryHandler) GetAllCategory(c echo.Context) error {
	resp, err := ch.categoryService.GetAll()
	if err != nil {
		return c.JSON(404, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"message": "Select All Category",
		"data":    resp,
	})
}

func (ch *CategoryHandler) ShowCategory(c echo.Context) error {
	id := c.Param("id")
	iduint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "id must number",
		})
	}
	resp, error := ch.categoryService.Show(uint(iduint))
	if error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "cannot get this data",
		})
	}
	return c.JSON(200, echo.Map{
		"message": "Category Show",
		"data": echo.Map{
			"id":   resp.ID,
			"Name": resp.Name,
		},
	})
}
func (ch *CategoryHandler) CreateCategory(c echo.Context) error {
	var req dto.CategoryCreate
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := ch.categoryService.Create(req.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Category Create",
		"data": echo.Map{
			"ID":   resp.ID,
			"Name": resp.Name,
		},
	})
}
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
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update category",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Category updated successfully",
		"data": echo.Map{
			"id":   category.ID,
			"name": category.Name,
		},
	})
}

func (ch *CategoryHandler) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := ch.categoryService.Delete(uint(id))
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Category deleted successfully",
		"data":    category,
	})
}
