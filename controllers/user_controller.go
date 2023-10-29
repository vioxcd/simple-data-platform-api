package controllers

import (
	"net/http"
	"vioxcd/dpl/middlewares"
	"vioxcd/dpl/models"
	"vioxcd/dpl/repo"

	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	err := repo.AddUser(&user)

	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		})
	}

	err = repo.GetUser(&user)

	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed to get data from the database",
			Data:    nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    user,
	})
}

func Login(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	err := repo.Login(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Username dan Password tidak cocok",
		})
	}

	var userResponse models.UserResponse
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Token = middlewares.GenerateToken(user.Id, user.Name)
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Login successful",
		Data: userResponse,
	})
}

