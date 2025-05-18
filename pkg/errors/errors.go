package apperror

import "gorm.io/gorm"

func HandleNotFoundError(err error, entity string, id uint) *AppError {
	if err == gorm.ErrRecordNotFound {
		return NotFound(entity, id)
	}
	return Internal(err)
}
