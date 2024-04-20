package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ILikeUsecase interface {
	CreateLike(like model.Like) (model.LikeResponse, error)
	DeleteLike(likeId uint) error
}

type likeUsecase struct {
	lr repository.ILikeRepository
}

func NewLikeUsecase(lr repository.ILikeRepository) ILikeUsecase {
	return &likeUsecase{lr}
}

func (lu *likeUsecase) CreateLike(like model.Like) (model.LikeResponse, error) {
	if err := lu.lr.CreateLike(&like); err != nil {
		return model.LikeResponse{}, err
	}
	resLike := model.LikeResponse{
		ID:         like.ID,
		TargetId:   like.TargetId,
		TargetType: like.TargetType,
		UserId:     like.UserId,
	}
	return resLike, nil
}

func (lu *likeUsecase) DeleteLike(likeId uint) error {
	if err := lu.lr.DeleteLike(likeId); err != nil {
		return err
	}
	return nil
}
