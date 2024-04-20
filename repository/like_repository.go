package repository

import (
	"fmt"
	"go-api/model"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	CreateLike(like *model.Like) error
	DeleteLike(likeId uint) error
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) ILikeRepository {
	return &likeRepository{db}
}

func (lr *likeRepository) CreateLike(like *model.Like) error {
	if err := lr.db.Create(like).Error; err != nil {
		return err
	}
	// if err := lr.db.Preload("User").First(like, "id=?", like.ID).Error; err != nil {
	if err := lr.db.First(like, "id=?", like.ID).Error; err != nil {
		return err
	}
	return nil
}

func (lr *likeRepository) DeleteLike(likeId uint) error {
	result := lr.db.Where("id=?", likeId).Delete(&model.Like{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
