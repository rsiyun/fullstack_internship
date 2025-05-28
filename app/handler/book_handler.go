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

// GetAllBook godoc
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} dto.MultipleBookResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Security BearerAuth
// @Router /admin/book [get]
func (bh *BookHandler) GetAllBook(c echo.Context) error {
	resp, err := bh.bookServices.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.BookErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.MultipleBookResponse{
		Message: "Select All Book",
		Data:    resp,
	})
}

// ShowBook godoc
// @Summary Get a book by ID
// @Description Retrieve a single book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.SingleBookResponse
// @Failure 400 {object} dto.BookErrorResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Security BearerAuth
// @Router /admin/book/{id} [get]
func (bh *BookHandler) ShowBook(c echo.Context) error {
	id := c.Param("id")
	iduint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BookErrorResponse{
			Message: "id must number",
		})
	}
	resp, error := bh.bookServices.Show(uint(iduint))
	if error != nil {
		return c.JSON(http.StatusBadRequest, dto.BookErrorResponse{
			Message: "cannot get this data",
		})
	}
	return c.JSON(http.StatusOK, dto.SingleBookResponse{
		Message: "Book Show",
		Data: dto.BookResponse{
			ID:          resp.ID,
			Title:       resp.Title,
			Description: resp.Description,
			CategoryId:  resp.CategoryID,
			Category: dto.CategoryResponse{
				ID:   resp.CategoryID,
				Name: resp.Category.Name,
			},
		},
	})
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the database
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.BookRequest true "Book data"
// @Success 200 {object} dto.SingleBookResponse
// @Failure 400 {object} dto.BookErrorResponse
// @Security BearerAuth
// @Router /admin/book [post]
func (bh *BookHandler) CreateBook(c echo.Context) error {
	var req dto.BookRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := bh.bookServices.Create(req.Title, req.Writer, req.Description, req.CategoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BookErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SingleBookResponse{
		Message: "Create Book",
		Data: dto.BookResponse{
			ID:          resp.ID,
			Title:       resp.Title,
			Description: resp.Description,
			CategoryId:  resp.CategoryID,
			Category: dto.CategoryResponse{
				ID:   resp.CategoryID,
				Name: resp.Category.Name,
			},
		},
	})
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update the details of an existing book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body dto.BookRequest true "Updated book data"
// @Success 200 {object} dto.SingleBookResponse
// @Failure 400 {object} dto.BookErrorResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Failure 500 {object} dto.BookErrorResponse
// @Security BearerAuth
// @Router /admin/book/{id} [put]
func (bh *BookHandler) UpdateBook(c echo.Context) error {
	var req dto.BookRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BookErrorResponse{
			Message: "Invalid ID format",
		})
	}
	book, err := bh.bookServices.UpdateBook(uint(id), req.Title, req.Writer, req.Description, req.CategoryID)
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, dto.BookErrorResponse{
				Message: err.Error(),
			})
		}
		if err.Error() == "book not found" {
			return c.JSON(http.StatusNotFound, dto.BookErrorResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.BookErrorResponse{
			Message: "Failed to update category",
		})
	}
	return c.JSON(http.StatusOK, dto.SingleBookResponse{
		Message: "Update book",
		Data: dto.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			CategoryId:  book.CategoryID,
			Category: dto.CategoryResponse{
				ID:   book.CategoryID,
				Name: book.Category.Name,
			},
		},
	})
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Remove a book from the database by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.SingleBookResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Failure 500 {object} dto.BookErrorResponse
// @Security BearerAuth
// @Router /admin/book/{id} [delete]
func (bh *BookHandler) DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := bh.bookServices.DeleteBook(uint(id))
	if err != nil {
		if err.Error() == "book not found" {
			return c.JSON(http.StatusNotFound, dto.BookErrorResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.BookErrorResponse{
			Message: "Failed to delete Book",
		})
	}

	return c.JSON(http.StatusOK, dto.SingleBookResponse{
		Message: "Update book",
		Data: dto.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			CategoryId:  book.CategoryID,
			Category: dto.CategoryResponse{
				ID:   book.CategoryID,
				Name: book.Category.Name,
			},
		},
	})
}

// GetBooksByCategory godoc
// @Summary Get books by category
// @Description Retrieve all books that belong to a specific category
// @Tags books
// @Accept json
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {object} dto.MultipleBookResponse
// @Failure 400 {object} dto.BookErrorResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Security BearerAuth
// @Router /admin/book/category/{category_id} [get]
func (bh *BookHandler) GetBooksByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BookErrorResponse{
			Message: "Invalid category ID format",
		})
	}

	books, err := bh.bookServices.GetBooksBycategory(uint(categoryID))
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, dto.BookErrorResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.BookErrorResponse{
			Message: "Failed to get Book",
		})
	}

	return c.JSON(http.StatusOK, dto.MultipleBookResponse{
		Message: "Books Fetched Successfully",
		Data:    books,
	})
}

// CountBooksByCategory godoc
// @Summary Count books by category
// @Description Count the number of books in a specific category
// @Tags books
// @Accept json
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {object} dto.CountBooksByCategory
// @Failure 400 {object} dto.BookErrorResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Security BearerAuth
// @Router /admin/book/category/{category_id}/count [get]
func (bh *BookHandler) CountBooksByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BookErrorResponse{
			Message: "Invalid category ID format",
		})
	}
	count, err := bh.bookServices.CountBooksByCategory(uint(categoryID))
	if err != nil {
		if err.Error() == "category not found" {
			return c.JSON(http.StatusNotFound, dto.BookErrorResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dto.BookErrorResponse{
			Message: "Failed to count books by category",
		})
	}

	return c.JSON(http.StatusOK, dto.CountBooksByCategory{
		Message: "Books count fetched successfully",
		Count:   count,
	})
}
