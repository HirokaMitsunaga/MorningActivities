package repository

import (
	"fmt"
	"go-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
	// 下記のSQLと同様の動作をしている
	// SELECT tasks.*
	// FROM tasks
	// JOIN Users ON tasks.user_id = Users.id
	// WHERE tasks.user_id = ?
	// ORDER BY tasks.created_at;
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
	// 下記のSQLと同様の動作をしている
	// SELECT tasks.*
	// FROM tasks
	// JOIN Users ON tasks.user_id = Users.id
	// WHERE tasks.user_id = ? AND tasks.id = ?
	// LIMIT 1;
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	//Clausesの設定のところは、更新した値をtaskの先に書き込んでくれる
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Updates(map[string]interface{}{"title": task.Title, "scheduled_minutes": task.ScheduledMinutes, "actual_minutes": task.ActualMinutes})
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

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
