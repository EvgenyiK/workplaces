package processors

import (
	"errors"
	"workplaces/server/internals/app/db"
	"workplaces/server/internals/app/models"
)

type UsersProcessor struct {
	storage *db.UsersStorage
}

func NewUsersProcessor(storage *db.UsersStorage) *UsersProcessor {
	processor := new(UsersProcessor)
	processor.storage = storage
	return processor
}

func (processor *UsersProcessor) CreateUser(user models.User) error {
	if user.PCName == "" {
		return errors.New("pc_name should not be empty")
	}
	if user.MacAdress == "" {
		return errors.New("mac_adress should not be empty")
	}
	if user.UserName == "" {
		return errors.New("name should not be empty")
	}

	return processor.storage.CreateUser(user)
}

func (processor *UsersProcessor) UpdateUser(id int64, user models.User) error {
	us := processor.storage.GetUserById(id)
	if us.Id != id {
		return errors.New("user not found")
	}

	return processor.storage.UpdateUser(id, user)
}

func (processor *UsersProcessor) DeleteUser(id int64) error {
	us := processor.storage.GetUserById(id)
	if us.Id != id {
		return errors.New("user not found")
	}

	return processor.storage.DeleteUser(id)
}
