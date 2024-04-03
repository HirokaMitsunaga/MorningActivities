package model

import "time"

type TimelineLikeAssociation struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	TimelineId    uint      `json:"timeline_id" gorm:"not null;index:idx_timeline_like_association,unique"`
	LikeCommentId uint      `json:"like_comment_id" gorm:"not null;index:idx_timeline_like_association,unique"`
	CreatedAt     time.Time `json:"created_at"`
}
