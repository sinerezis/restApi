package model

import validation "github.com/go-ozzo/ozzo-validation"

// Кастомная валидация  Нужна для того, чтобы
// в случае перезаписи пользователя
// он смог пройти валидацию по полю password
// тк в данном случае оно остается пустым
func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
