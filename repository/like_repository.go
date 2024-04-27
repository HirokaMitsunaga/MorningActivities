package repository

import (
	"fmt"
	"go-api/model"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	CreateLike(like *model.Like) error
	DeleteLike(like model.Like) error
	ToggleLike(like *model.Like) error
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
	// タイムラインのLikeCountを更新
	//likeを集計してlikeCountへ代入する
	var likeCount int64
	if err := lr.db.Model(&model.Like{}).Where("target_id = ? AND target_type = ?", like.TargetId, like.TargetType).Count(&likeCount).Error; err != nil {
		return err
	}
	//集計したlikeCountをtimeline.LikeCountに反映
	if err := lr.db.Model(&model.Timeline{}).Where("id = ?", like.TargetId).Update("like_count", likeCount).Error; err != nil {
		return err
	}

	if err := lr.db.First(like, "id=?", like.ID).Error; err != nil {
		return err
	}
	return nil
}

func (lr *likeRepository) DeleteLike(like model.Like) error {
	result := lr.db.Where("id=?", like.ID).Delete(&model.Like{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	// タイムラインのLikeCountを更新
	//likeを集計してlikeCountへ代入する
	var likeCount int64
	if err := lr.db.Model(&model.Like{}).Where("target_id = ? AND target_type = ?", like.TargetId, like.TargetType).Count(&likeCount).Error; err != nil {
		return err
	}
	//集計したlikeCountをtimeline.LikeCountに反映
	if err := lr.db.Model(&model.Timeline{}).Where("id = ?", like.TargetId).Update("like_count", likeCount).Error; err != nil {
		return err
	}
	return nil
}

func (lr *likeRepository) ToggleLike(like *model.Like) error {
	var existingLike model.Like
	if err := lr.db.Where("user_id = ? AND target_id = ? AND target_type = ?", like.UserId, like.TargetId, like.TargetType).First(&existingLike).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return lr.CreateLike(like)
		}
		return err
	}
	return lr.DeleteLike(existingLike)
}
