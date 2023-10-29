package controllers

import (
	"net/http"
	"vioxcd/dpl/middlewares"
	"vioxcd/dpl/models"
	"vioxcd/dpl/repo"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

const (
	ROLE_ADMIN = 1
)

func AddUser(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(400, models.BaseResponse{
			Status:  false,
			Message: "JWT token missing or invalid",
			Data:    nil,
		})
	}

	// not sure why `claims := user.Claims.(*jwtCustomClaims)` doesn't work
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed to cast claims as jwt.MapClaims",
			Data:    claims,
		})
	}

	// https://stackoverflow.com/a/70706082
	if uint(claims["roleId"].(float64)) != uint(ROLE_ADMIN) {
		return c.JSON(403, models.BaseResponse{
			Status:  false,
			Message: "Unauthorized to create user",
			Data:    nil,
		})
	}

	// pass privilege checks, creating user
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
	userResponse.Token = middlewares.GenerateToken(user.Id, user.Name, user.RoleId)
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Login successful",
		Data:    userResponse,
	})
}
