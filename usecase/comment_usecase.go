package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ICommentUsecase interface {
	GetAllComments(userId uint) ([]model.CommentResponse, error)
	GetCommentById(commentId uint) (model.CommentResponse, error)
	GetCommentsByTimelineId(timelineId uint) ([]model.CommentResponse, error)
	CreateComment(comment model.Comment) (model.CommentResponse, error)
	UpdateComment(comment model.Comment, userId uint, commentId uint) (model.CommentResponse, error)
	DeleteComment(userId uint, commentId uint) error
}

type commentUsecase struct {
	cr repository.ICommentRepository
}

func NewCommentUsecase(cr repository.ICommentRepository) ICommentUsecase {
	return &commentUsecase{cr}
}

func (cu *commentUsecase) GetAllComments(userId uint) ([]model.CommentResponse, error) {
	comments := []model.Comment{}
	if err := cu.cr.GetAllComments(&comments, userId); err != nil {
		return nil, err
	}
	resComments := []model.CommentResponse{}
	for _, v := range comments {
		t := model.CommentResponse{
			ID:         v.ID,
			TimelineId: v.TimelineId,
			UserId:     v.UserId,
			Comment:    v.Comment,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}
		resComments = append(resComments, t)
	}
	return resComments, nil
}

func (cu *commentUsecase) GetCommentById(commentId uint) (model.CommentResponse, error) {
	comment := model.Comment{}
	if err := cu.cr.GetCommentById(&comment, commentId); err != nil {
		return model.CommentResponse{}, err
	}
	resComment := model.CommentResponse{
		ID:         comment.ID,
		TimelineId: comment.TimelineId,
		UserId:     comment.UserId,
		Comment:    comment.Comment,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
	}
	return resComment, nil
}

func (cu *commentUsecase) GetCommentsByTimelineId(timelineId uint) ([]model.CommentResponse, error) {
	comments := []model.Comment{}
	if err := cu.cr.GetCommentsByTimelineId(&comments, timelineId); err != nil {
		return nil, err
	}
	resComments := []model.CommentResponse{}
	for _, v := range comments {
		t := model.CommentResponse{
			ID:         v.ID,
			TimelineId: v.TimelineId,
			UserId:     v.UserId,
			Comment:    v.Comment,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}
		resComments = append(resComments, t)
	}
	return resComments, nil
}

func (cu *commentUsecase) CreateComment(comment model.Comment) (model.CommentResponse, error) {
	if err := cu.cr.CreateComment(&comment); err != nil {
		return model.CommentResponse{}, err
	}
	resComment := model.CommentResponse{
		ID:         comment.ID,
		TimelineId: comment.TimelineId,
		UserId:     comment.UserId,
		Comment:    comment.Comment,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
	}
	return resComment, nil
}

func (cu *commentUsecase) UpdateComment(comment model.Comment, userId uint, commentId uint) (model.CommentResponse, error) {
	if err := cu.cr.UpdateComment(&comment, userId, commentId); err != nil {
		return model.CommentResponse{}, err
	}
	resComment := model.CommentResponse{
		ID:         comment.ID,
		TimelineId: comment.TimelineId,
		UserId:     comment.UserId,
		Comment:    comment.Comment,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
	}
	return resComment, nil
}

func (cu *commentUsecase) DeleteComment(userId uint, commentId uint) error {
	if err := cu.cr.DeleteComment(userId, commentId); err != nil {
		return err
	}
	return nil
}
