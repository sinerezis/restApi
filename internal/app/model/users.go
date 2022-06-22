package model

// Структура - модель данных для нашей таблицы Users
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
}
