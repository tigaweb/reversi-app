package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type IMoveRepository interface {
	RegisterMove(move *model.Move) error
}

type moveRepository struct {
	db *gorm.DB
}

func NewMoveRepository(db *gorm.DB) IMoveRepository {
	return &moveRepository{db}
}

func (mr *moveRepository) RegisterMove(move *model.Move) error {
	if err := mr.db.Create(&move).Error; err != nil {
		return err
	}
	return nil
}
