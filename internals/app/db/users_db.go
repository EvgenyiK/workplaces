package db

import (
	"context"
	"workplaces/internals/app/models"

	"github.com/georgysavva/scany/pgxscan"
	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UsersStorage struct {
	databasePool *pgxpool.Pool
}

func NewUsersStorage(pool *pgxpool.Pool) *UsersStorage {
	storage := new(UsersStorage)
	storage.databasePool = pool
	return storage
}

func (storage *UsersStorage) CreateUser(user models.User) error {
	query := "INSERT INTO users(pc_name, mac_adress, user_name) VALUES ($1, $2, $3)"

	_, err := storage.databasePool.Exec(context.Background(), query, user.PCName, user.MacAdress, user.UserName)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

func (storage *UsersStorage) UpdateUser(id int64, user models.User) error {
	query := "UPDATE users SET pc_name=$2, mac_adress=$3, user_name=$4 WHERE id=$1"

	_, err := storage.databasePool.Exec(context.Background(), query, id, user.PCName, user.MacAdress, user.UserName)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

func (storage *UsersStorage) DeleteUser(id int64) error {
	query := "DELETE FROM users WHERE id=$1"

	_, err := storage.databasePool.Exec(context.Background(), query, id)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

func (storage *UsersStorage) GetUserById(id int64) models.User {
	query := "SELECT id, pc_name, mac_adress, user_name FROM users WHERE id = $1"

	var result models.User

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, id)

	if err != nil {
		log.Errorln(err)
	}

	return result
}
