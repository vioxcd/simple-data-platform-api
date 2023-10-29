package controllers

import (
	"github.com/labstack/echo/v4"
	"vioxcd/dpl/models"
	"vioxcd/dpl/repo"
)


func GetRunsHistory(c echo.Context) error {
	var logs []models.UserLog

	err := repo.GetLogs(&logs)

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
		Data:    logs,
	})
}

func RunNewSnapshot(c echo.Context) error {
	var log models.UserLog
	c.Bind(&log)

	err := repo.TriggerNewRun(&log)

	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		})
	}

	// populate log with user and run data
	err = repo.GetLog(&log)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed to get data in database",
			Data:    nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    log,
	})
}
