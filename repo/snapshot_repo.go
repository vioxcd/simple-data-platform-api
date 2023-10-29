package repo

import (
	"vioxcd/dpl/config"
	"vioxcd/dpl/models"
)

func GetSnapshots(snapshots *[]models.Snapshot) error {
	result := config.DB.Find(snapshots)
	return result.Error
}
