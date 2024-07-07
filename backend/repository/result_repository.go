package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type IResultRepository interface {
	RegisterResult(result model.GameResult) error
}

type resultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) IResultRepository {
	return &resultRepository{db}
}

func (rr *resultRepository) RegisterResult(result model.GameResult) error {
	if err := rr.db.Create(&result).Error; err != nil {
		return nil
	}
	return nil
}
