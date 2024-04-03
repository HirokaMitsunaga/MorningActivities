package model

import "time"

type LikeComment struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	User             User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId           uint      `json:"user_id" gorm:"not null"`
	IsLike           bool      `json:"is_like" gorm:"not null"`
	Comment          string    `json:"comment" gorm:"not null"`
	LikeCount        int       `json:"like_count" gorm:"not null"`
	CommentLikeCount int       `json:"comment_like_count" gorm:"not null"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type LikeCommentResponse struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Sentence     string    `json:"sentence" gorm:"not null"`
	LikeCount    int       `json:"like_count" gorm:"not null"`
	CommentCount int       `json:"comment_count" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserId       uint      `json:"user_id"`
}
