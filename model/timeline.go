package model

import "time"

type Timeline struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Sentence     string    `json:"sentence" gorm:"not null"`
	LikeCount    int       `json:"like_count" gorm:"not null"`
	CommentCount int       `json:"comment_count" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId       uint      `json:"user_id" gorm:"not null"`
	Email        string    `json:"email" gorm:"not null"`
}

type TimelineResponse struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Sentence     string    `json:"sentence" gorm:"not null"`
	LikeCount    int       `json:"like_count" gorm:"not null"`
	CommentCount int       `json:"comment_count" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserId       uint      `json:"user_id"`
	Email        string    `json:"email" gorm:"not null"`
}
