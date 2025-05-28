package handler

import (
	"dot/app/models/dto"
	"dot/app/services"
	"dot/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookServices services.BookServices
}

func NewBookHandler(bookServices services.BookServices) *BookHandler {
	return &BookHandler{bookServices: bookServices}
}

func (bh *BookHandler) GetAllBook(c echo.Context) error {
	resp, err := bh.bookServices.GetAllBooks()
	if err != nil {
		return c.JSON(404, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"message": "Select All Book",
		"data":    resp,
	})
}
func (bh *BookHandler) ShowBook(c echo.Context) error {
	id := c.Param("id")
	iduint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "id must number",
		})
	}
	resp, error := bh.bookServices.Show(uint(iduint))
	if error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "cannot get this data",
		})
	}
	return c.JSON(200, echo.Map{
		"message": "Book Show",
		"data": echo.Map{
			"id":          resp.ID,
			"title":       resp.Title,
			"description": resp.Description,
			"category":    resp.Category.Name,
		},
	})
}
func (bh *BookHandler) CreateBook(c echo.Context) error {
	var req dto.BookRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := bh.bookServices.Create(req.Title, req.Writer, req.Description, req.CategoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Book Create",
		"data": echo.Map{
			"ID":   resp.ID,
			"Name": resp.Title,
		},
	})
}
func (bh *BookHandler) UpdateBook(c echo.Context) error {
	var req dto.BookRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid ID format",
		})
	}
	book, err := bh.bookServices.UpdateBook(uint(id), req.Title, req.Writer, req.Description, req.CategoryID)
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		if err.Error() == "book not found" {
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
		"message": "Book updated successfully",
		"data": echo.Map{
			"title":       book.Title,
			"writer":      book.Writer,
			"description": book.Description,
			"category":    book.Category,
		},
	})
}
func (bh *BookHandler) DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := bh.bookServices.DeleteBook(uint(id))
	if err != nil {
		if err.Error() == "book not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete Book",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Book deleted successfully",
		"data":    book,
	})
}
func (bh *BookHandler) GetBooksByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid category ID format",
		})
	}

	books, err := bh.bookServices.GetBooksBycategory(uint(categoryID))
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to fetch books by category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Books fetched successfully",
		"data":    books,
	})
}

func (bh *BookHandler) CountBooksByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid category ID format",
		})
	}
	count, err := bh.bookServices.CountBooksByCategory(uint(categoryID))
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to count books by category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Books count fetched successfully",
		"data":    count,
	})
}
