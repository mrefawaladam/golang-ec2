package controllers

import (
	"curdusers/configs"
	"curdusers/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog
	if err := configs.DB.Find(&blogs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"blogs":   blogs,
	})

}

func GetBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid blog id",
		})
	}
	var blog models.Blog
	if err := configs.DB.First(&blog, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Blog not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"blog":    blog,
	})
}

func CreateBlogController(c echo.Context) error {
	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := configs.DB.Create(blog).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

func UpdateBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid blog id",
		})
	}
	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := configs.DB.First(&models.Blog{}, id).Updates(blog).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updated",
		"blog":    blog,
	})
}

func DeleteBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid blog id",
		})
	}
	if err := configs.DB.Delete(&models.Blog{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleted blog",
	})
}
