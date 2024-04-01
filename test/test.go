package main

import (
	"go-api/db"
	"go-api/model"
	"time"
)

// 下記を実行すると、timelineテーブルへテストデータが挿入される
func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	// テストデータの作成
	testTimelines := []model.Timeline{
		{
			Sentence:     "This is the first test timeline",
			LikeCount:    5,
			CommentCount: 2,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			UserId:       1, // 既存のユーザIDを指定する
		},
		{
			Sentence:     "This is the second test timeline",
			LikeCount:    3,
			CommentCount: 1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			UserId:       1, // 既存のユーザIDを指定する
		},
	}

	// テストデータの挿入
	for _, timeline := range testTimelines {
		dbConn.Create(&timeline)
	}
}
