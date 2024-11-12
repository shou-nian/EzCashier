package queries

import (
	"github.com/shou-nian/EzCashier/app/models"
	"gorm.io/gorm"
)

type UserQueries struct {
	*gorm.DB
}

func (q *UserQueries) CreateUser(user *models.User) error {
	return nil
}
