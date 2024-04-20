package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ITimelineUsecase interface {
	GetAllTimelines() ([]model.TimelineResponse, error)
	GetTimelineById(timelineId uint) (model.TimelineResponse, error)
	CreateTimeline(timeline model.Timeline) (model.TimelineResponse, error)
	UpdateTimeline(timeline model.Timeline, timelineId uint) (model.TimelineResponse, error)
	DeleteTimeline(timelineId uint) error
}

type timelineUsecase struct {
	tlr repository.ITimelineRepository
}

func NewTimelineUsecase(tlr repository.ITimelineRepository) ITimelineUsecase {
	return &timelineUsecase{tlr}
}

func (tlu *timelineUsecase) GetAllTimelines() ([]model.TimelineResponse, error) {
	timelines := []model.Timeline{}
	if err := tlu.tlr.GetAllTimelines(&timelines); err != nil {
		return nil, err
	}
	resTimelines := []model.TimelineResponse{}
	for _, v := range timelines {
		t := model.TimelineResponse{
			ID:           v.ID,
			Sentence:     v.Sentence,
			LikeCount:    v.LikeCount,
			CommentCount: v.CommentCount,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			UserId:       v.UserId,
			//EmailはUser構造体から取得
			Email: v.User.Email,
		}

		resTimelines = append(resTimelines, t)
	}
	return resTimelines, nil
}

func (tlu *timelineUsecase) GetTimelineById(timelineId uint) (model.TimelineResponse, error) {
	timeline := model.Timeline{}
	if err := tlu.tlr.GetTimelineById(&timeline, timelineId); err != nil {
		return model.TimelineResponse{}, err
	}
	resTimeline := model.TimelineResponse{
		ID:           timeline.ID,
		Sentence:     timeline.Sentence,
		LikeCount:    timeline.LikeCount,
		CommentCount: timeline.CommentCount,
		CreatedAt:    timeline.CreatedAt,
		UpdatedAt:    timeline.UpdatedAt,
		UserId:       timeline.UserId,
		Email:        timeline.User.Email,
	}
	return resTimeline, nil
}

func (tlu *timelineUsecase) CreateTimeline(timeline model.Timeline) (model.TimelineResponse, error) {
	// if err := tlu.tlv.TaskValidate(task); err != nil {
	// 	return model.TaskResponse{}, err
	// }
	if err := tlu.tlr.CreateTimeline(&timeline); err != nil {
		return model.TimelineResponse{}, err
	}
	resTimeline := model.TimelineResponse{
		ID:           timeline.ID,
		Sentence:     timeline.Sentence,
		LikeCount:    timeline.LikeCount,
		CommentCount: timeline.CommentCount,
		CreatedAt:    timeline.CreatedAt,
		UpdatedAt:    timeline.UpdatedAt,
		UserId:       timeline.UserId,
		Email:        timeline.User.Email,
	}
	return resTimeline, nil
}

func (tlu *timelineUsecase) UpdateTimeline(timeline model.Timeline, timelineId uint) (model.TimelineResponse, error) {
	// if err := tlu.tv.TaskValidate(task); err != nil {
	// 	return model.TaskResponse{}, err
	// }
	if err := tlu.tlr.UpdateTimeline(&timeline, timelineId); err != nil {
		return model.TimelineResponse{}, err
	}
	resTimeline := model.TimelineResponse{
		ID:           timeline.ID,
		Sentence:     timeline.Sentence,
		CreatedAt:    timeline.CreatedAt,
		UpdatedAt:    timeline.UpdatedAt,
		LikeCount:    timeline.LikeCount,
		CommentCount: timeline.CommentCount,
		UserId:       timeline.UserId,
		Email:        timeline.User.Email,
	}
	return resTimeline, nil
}

func (tlu *timelineUsecase) DeleteTimeline(timelineId uint) error {
	if err := tlu.tlr.DeleteTimeline(timelineId); err != nil {
		return err
	}
	return nil
}
