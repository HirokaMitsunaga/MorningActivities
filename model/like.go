package model

import "time"

type Like struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	TargetId   uint      `json:"target_id" gorm:"not null"`   // タイムラインまたはコメントのID
	TargetType string    `json:"target_type" gorm:"not null"` // 'timeline' または 'comment'
	UserId     uint      `json:"user_id" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
}

type LikeResponse struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	TargetId   uint      `json:"target_id" gorm:"not null"`   // タイムラインまたはコメントのID
	TargetType string    `json:"target_type" gorm:"not null"` // 'timeline' または 'comment'
	UserId     uint      `json:"user_id" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
}
