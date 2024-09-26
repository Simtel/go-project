package usersrepo

import (
	"errors"
	"go-project/internal/models"
)

func (r *UsersRepo) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}
	return users, nil
}
