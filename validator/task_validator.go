package validator

import (
	"go-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 20).Error("limited max 20 char"),
		),
		validation.Field(
			&task.ScheduledMinutes,
			validation.Required.Error("scheduled minutes are required"),
			validation.Min(1).Error("scheduled minutes must be at least 1"),
			validation.Max(1440).Error("scheduled minutes must be less than 1440"), // 1日の分数
		),
		validation.Field(
			&task.ActualMinutes,
			validation.Min(0).Error("actual minutes must be at least 0"),
			validation.Max(1440).Error("actual minutes must be less than 1440"), // 1日の分数
		),
	)
}
