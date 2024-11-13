package migrations

import (
	"github.com/shou-nian/EzCashier/app/models"
	"github.com/shou-nian/EzCashier/repository/database"
)

func AutoMigration() error {
	db, err := database.MysqlConnection()
	if err != nil {
		return err
	}

	return db.AutoMigrate(&models.User{})
}
