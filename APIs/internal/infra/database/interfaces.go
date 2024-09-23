package database

import "github.com/andremarquezz/Go_Expert/APIS/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}
