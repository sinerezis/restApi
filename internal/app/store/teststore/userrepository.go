package teststore

import "resApi/internal/app/model"

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func Create(u *model.User)
