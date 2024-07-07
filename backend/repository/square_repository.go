package repository

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type ISquareRepository interface {
	CreateSquares(turn_id uint, board model.Board) error
	// TODO 盤面を取得する
	GetSquaresByTurnId(turn_id uint) ([]model.Square, error)
	FindWinnerDiscByTurnId(turn_id uint) (model.Disc, error)
}

type squareRepository struct {
	db *gorm.DB
}

func NewSquareRepository(db *gorm.DB) ISquareRepository {
	return &squareRepository{db}
}

func (sr *squareRepository) CreateSquares(turn_id uint, board model.Board) error {
	for y, row := range board.Discs {
		for x, disc := range row {
			square := &model.Square{
				TurnId: turn_id,
				X:      x,
				Y:      y,
				Disc:   int(disc),
			}
			if err := sr.db.Create(&square).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (sr *squareRepository) GetSquaresByTurnId(turn_id uint) ([]model.Square, error) {
	var squares []model.Square
	if err := sr.db.Where("turn_id = ?", turn_id).Find(&squares).Error; err != nil {
		return nil, err
	}
	return squares, nil
}

func (sr *squareRepository) FindWinnerDiscByTurnId(turn_id uint) (model.Disc, error) {
	var result struct {
		Disc  int
		Count int
	}

	err := sr.db.Table("squares").
		Select("disc, COUNT(*) as count").
		Where("turn_id = ?", turn_id).
		Group("disc").
		Order("count DESC").
		First(&result).Error

	if err != nil {
		return 0, err
	}

	return model.Disc(result.Disc), nil
}
