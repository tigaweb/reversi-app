package repository

import "github.com/tigaweb/reversi-app/backend/model"

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
