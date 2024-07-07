package usecase

import (
	"github.com/tigaweb/reversi-app/backend/model"
	"github.com/tigaweb/reversi-app/backend/repository"
)

type IMoveUsecase interface {
	RegisterMove(move *model.Move) error
}

type moveUsecase struct {
	mr repository.IMoveRepository
}

func NewMoveUsecase(mr repository.IMoveRepository) IMoveUsecase {
	return &moveUsecase{mr}
}
func (mu *moveUsecase) RegisterMove(move *model.Move) error {
	if err := mu.mr.RegisterMove(move); err != nil {
		return err
	}
	return nil
}
