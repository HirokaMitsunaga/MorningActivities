package model

import "time"

type Task struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	Title            string    `json:"title" gorm:"not null"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	User             User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId           uint      `json:"user_id" gorm:"not null"`
	ScheduledMinutes int       `json:"scheduled_minutes"`
	ActualMinutes    int       `json:"actual_minutes"`
}

type TaskResponse struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	Title            string    `json:"title" gorm:"not null"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ScheduledMinutes int       `json:"scheduled_minutes"`
	ActualMinutes    int       `json:"actual_minutes"`
}
