package validator

import (
	"github.com/tigaweb/reversi-app/backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserCreateValidate(user model.User) error
	UserLoginValidate(user model.User) error
}

type userValidotor struct{}

func NewUserValidator() IUserValidator {
	return &userValidotor{}
}

func (uv *userValidotor) UserCreateValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.UserName,
			validation.Required.Error("user name is required"),
			validation.RuneLength(1, 8).Error("limited max 8 char"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 18).Error("limited min 6 max 18"),
		),
	)
}

func (uv *userValidotor) UserLoginValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 18).Error("limited min 6 max 18"),
		),
	)
}
