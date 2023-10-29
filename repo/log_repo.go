package repo

import (
	"fmt"

	"vioxcd/dpl/config"
	"vioxcd/dpl/models"
)

func TriggerNewRun(log *models.UserLog) error {
	result := config.DB.Create(log)
	return result.Error
}

func GetLog(log *models.UserLog) error {
	result := config.DB.Joins("Run").Joins("User").Find(log)
	fmt.Println(log)
	return result.Error
}

func GetLogs(logs *[]models.UserLog) error {
	result := config.DB.Joins("Run").Joins("User").Find(logs)
	fmt.Println(logs)
	return result.Error
}
