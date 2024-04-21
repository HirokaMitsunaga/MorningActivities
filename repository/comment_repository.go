package repository

import (
	"fmt"
	"go-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ICommentRepository interface {
	GetAllComments(comments *[]model.Comment, userId uint) error
	GetCommentById(comment *model.Comment, commentId uint) error
	GetCommentsByTimelineId(comments *[]model.Comment, timelineId uint) error
	CreateComment(comment *model.Comment) error
	UpdateComment(comment *model.Comment, userId uint, commentId uint) error
	DeleteComment(userId uint, commentId uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) GetAllComments(comments *[]model.Comment, userId uint) error {
	if err := cr.db.
		Joins("JOIN users ON comments.user_id = users.id").
		Where("comments.user_id = ?", userId).
		Order("comments.created_at").
		Find(comments).Error; err != nil {
		return err
	}
	return nil
}
func (cr *commentRepository) GetCommentById(comment *model.Comment, commentId uint) error {
	if err := cr.db.
		Where("comments.id = ?", commentId).
		Find(comment).Error; err != nil {
		return err
	}
	return nil
	// 下記のSQLと同様の動作をしている
	// SELECT comments.*
	// FROM comments
	// JOIN Timeline ON comments.timeline_id = Timeline.id
	// WHERE comments.timeline_id = ? AND comments.id = ?
	// LIMIT 1;
}

// 特定のタイムラインに紐づくコメントを全て取り出す
func (cr *commentRepository) GetCommentsByTimelineId(comments *[]model.Comment, timelineId uint) error {
	if err := cr.db.
		Joins("JOIN timelines ON comments.timeline_id = timelines.id").
		Where("comments.timeline_id = ?", timelineId).
		Find(comments).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) CreateComment(comment *model.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) UpdateComment(comment *model.Comment, userId uint, commentId uint) error {
	//Clausesの設定のところは、更新した値をcommentの先に書き込んでくれる
	result := cr.db.Debug().Model(comment).Clauses(clause.Returning{}).Where("id=? AND user_id=?", commentId, userId).Updates(map[string]interface{}{"comment": comment.Comment})
	if result.Error != nil {
		return result.Error
	}
	//更新の場合オブジェクトが存在しなかった場合エラーにならない仕様のため下記のif文を追記
	//RowsAffectedは実際に更新したレコードの数を表している
	//RowsAffectedが1より小さい場合
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (cr *commentRepository) DeleteComment(userId uint, commentId uint) error {
	result := cr.db.Where("id=? AND user_id=?", commentId, userId).Delete(&model.Comment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
