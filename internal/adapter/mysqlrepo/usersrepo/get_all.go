package usersrepo

import (
	"errors"
	"go-project/internal/models/db"
)

func (r *UsersRepo) GetAll() ([]db.User, error) {
	var users []db.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}
	return users, nil
}
