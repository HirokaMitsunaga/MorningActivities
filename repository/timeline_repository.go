package repository

import (
	"fmt"
	"go-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITimelineRepository interface {
	GetAllTimelines(timelines *[]model.Timeline) error
	GetTimelineById(timeline *model.Timeline, timelineId uint) error
	CreateTimeline(timeline *model.Timeline) error
	UpdateTimeline(timeline *model.Timeline, timelineId uint) error
	DeleteTimeline(timelineId uint) error
}

type timelineRepository struct {
	db *gorm.DB
}

func NewTimelineRepository(db *gorm.DB) ITimelineRepository {
	return &timelineRepository{db}
}

func (tlr *timelineRepository) GetAllTimelines(timelines *[]model.Timeline) error {
	if err := tlr.db.Preload("User").Find(timelines).Error; err != nil {
		return err
	}
	//いいね数を数える
	for index, timeline := range *timelines {
		var likeCount int64
		if err := tlr.db.Model(&model.Like{}).Where("target_id = ? AND target_type = ?", timeline.ID, "timeline").Count(&likeCount).Error; err != nil {
			return err
		}
		(*timelines)[index].LikeCount = int(likeCount)
	}
	//コメント数を数える
	for index, timeline := range *timelines {
		var commentCount int64
		if err := tlr.db.Model(&model.Comment{}).Where("timeline_id = ?", timeline.ID).Count(&commentCount).Error; err != nil {
			return err
		}
		(*timelines)[index].CommentCount = int(commentCount)
	}
	return nil
}

func (tlr *timelineRepository) GetTimelineById(timeline *model.Timeline, timelineId uint) error {
	//email情報を取得するためPreload("User")は必要
	if err := tlr.db.Preload("User").First(timeline, "id=?", timelineId).Error; err != nil {
		return err
	}
	//いいね数を数える
	var likeCount int64
	if err := tlr.db.Model(&model.Like{}).Where("target_id = ? AND target_type = ?", timeline.ID, "timeline").Count(&likeCount).Error; err != nil {
		return err
	}
	timeline.LikeCount = int(likeCount)

	//コメント数を数える
	var commentCount int64
	if err := tlr.db.Model(&model.Comment{}).Where("timeline_id = ?", timeline.ID).Count(&commentCount).Error; err != nil {
		return err
	}
	timeline.CommentCount = int(commentCount)

	return nil
}

func (tlr *timelineRepository) CreateTimeline(timeline *model.Timeline) error {
	if err := tlr.db.Create(timeline).Error; err != nil {
		return err
	}
	if err := tlr.db.Preload("User").First(timeline, "id=?", timeline.ID).Error; err != nil {
		return err
	}
	return nil
}

func (tlr *timelineRepository) UpdateTimeline(timeline *model.Timeline, timelineId uint) error {
	//Clausesの設定のところは、更新した値をtimelineの先に書き込んでくれる
	result := tlr.db.Model(timeline).Clauses(clause.Returning{}).Where("id=?", timelineId).Updates(map[string]interface{}{"sentence": timeline.Sentence})
	if result.Error != nil {
		return result.Error
	}
	//更新の場合オブジェクトが存在しなかった場合エラーにならない仕様のため下記のif文を追記
	//RowsAffectedは実際に更新したレコードの数を表している
	//RowsAffectedが1より小さい場合
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	err := tlr.db.Preload("User").First(timeline, "id=?", timeline.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (tlr *timelineRepository) DeleteTimeline(timelineId uint) error {
	result := tlr.db.Where("id=?", timelineId).Delete(&model.Timeline{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
