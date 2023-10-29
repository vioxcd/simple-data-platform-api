package controllers

import (
	"github.com/labstack/echo/v4"
	"vioxcd/dpl/models"
	"vioxcd/dpl/repo"
)

func GetSnapshots(c echo.Context) error {
	var snapshots []models.Snapshot

	err := repo.GetSnapshots(&snapshots)

	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully get data",
		Data:    snapshots,
	})
}
