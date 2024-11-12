package queries

import (
	"github.com/shou-nian/EzCashier/app/models"
	"gorm.io/gorm"
)

type UserQueries struct {
	*gorm.DB
}

func (q *UserQueries) CreateUser(user *models.User) (*models.User, error) {
	result := q.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (q *UserQueries) UpdateUser(user *models.User) (*models.User, error) {
	result := q.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (q *UserQueries) DeleteUser(user *models.User) error {
	result := q.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (q *UserQueries) GetUsers() ([]*models.User, error) {
	var users []*models.User
	result := q.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (q *UserQueries) GetUser(id uint) (*models.User, error) {
	var user *models.User
	result := q.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
