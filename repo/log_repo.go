package repo

import (
	"vioxcd/dpl/config"
	"vioxcd/dpl/models"

	"gorm.io/gorm"
)

func TriggerNewRun(log *models.UserLog) error {
	result := config.DB.Create(log)
	if result.Error != nil {
		return result.Error
	}

	return config.DB.Transaction(func(tx *gorm.DB) error {
	  if err := tx.Exec("TRUNCATE TABLE snapshots").Error; err != nil {
		return err
	  }

	  if err := tx.Exec(`
			INSERT INTO snapshots
			SELECT NULL, date, sum(amount)
			FROM transactions
			GROUP BY date
		`).Error; err != nil {
		return err
	  }

	  return nil
	})
}

func GetLog(log *models.UserLog) error {
	result := config.DB.Joins("Run").Joins("User").Find(log)
	return result.Error
}

func GetLogs(logs *[]models.UserLog) error {
	result := config.DB.Joins("Run").Joins("User").Find(logs)
	return result.Error
}
