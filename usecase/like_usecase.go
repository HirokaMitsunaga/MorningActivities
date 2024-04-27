package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ILikeUsecase interface {
	CreateLike(like model.Like) (model.LikeResponse, error)
	DeleteLike(like model.Like) error
	ToggleLike(like *model.Like) error
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

func (lu *likeUsecase) DeleteLike(like model.Like) error {
	if err := lu.lr.DeleteLike(like); err != nil {
		return err
	}
	return nil
}

func (lu *likeUsecase) ToggleLike(like *model.Like) error {
	return lu.lr.ToggleLike(like)
}
