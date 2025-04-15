package orm

import (
	"errors"
	"gorm.io/gorm"
)

func IsNotFound(seed error) bool {
	return seed != nil && errors.Is(seed, gorm.ErrRecordNotFound)
}

func IsFoundButHasErrors(seed error) bool {
	return seed != nil && !errors.Is(seed, gorm.ErrRecordNotFound)
}
