package controllers

import (
	"curdusers/configs"
	"curdusers/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book
	if err := configs.DB.Preload("Author").Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid book id")
	}

	var book models.Book
	if err := configs.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "book not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	if err := c.Bind(&book); err != nil {
		log.Errorf("Failed to bind request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := configs.DB.Create(&book).Error; err != nil {
		log.Errorf("Failed to create book: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid book id")
	}

	var book models.Book
	if err := configs.DB.First(&book, bookId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "book not found")
	}

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := configs.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updated book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	// get book id from url param
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid book id")
	}

	// get book by id
	var book models.Book
	if err := configs.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "book not found")
	}

	// delete book from database
	if err := configs.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete book")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleted",
	})
}
