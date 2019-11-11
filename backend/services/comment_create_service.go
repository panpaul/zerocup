package services

import (
	"comment/e"
	"comment/models"
	"comment/serializers"
)

type CreateCommentService struct {
	ReplyToID uint   `json:"reply_to_id" form:"reply_to_id"`
	ArticleID uint   `json:"article_id" form:"article_id" binding:"required"`
	ParentID  uint   `json:"parent_id" form:"parent_id"`
	RootID    uint   `json:"root_id" form:"root_id"`
	Content   string `json:"content" form:"content" binding:"required,min=3"`
}

func (service *CreateCommentService) Create(user *models.User) *serializers.Response {
	comment := models.Comment{
		UserID:    user.ID,
		ReplyToID: service.ReplyToID,
		ArticleID: service.ArticleID,
		ParentID:  service.ParentID,
		RootID:    service.RootID,
		Content:   service.Content,
	}
	if err := models.DB.Create(&comment).Error; err != nil {
		return &serializers.Response{
			Status:  e.CREATE_ERROR,
			Message: e.GetMsg(e.CREATE_ERROR),
			Error:   err.Error(),
		}
	}

	// 返回当前评论的用户和   replyTo用户(如果是回复的话)
	comment.User = *user
	if comment.ReplyToID != 0 {
		var replyToUser models.User
		if err := models.DB.Find(&replyToUser, service.ReplyToID).Error; err != nil {
			return &serializers.Response{
				Status:e.SELECT_ERROR,
				Message:e.GetMsg(e.SELECT_ERROR),
				Error:err.Error(),
			}
		}
		comment.ReplyTo = replyToUser
		return &serializers.Response{
			Status:  e.SUCCESS,
			Message: e.GetMsg(e.SUCCESS),
			Data:    serializers.BuildReply(comment),
		}
	}

	return &serializers.Response{
		Status:  e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data:    serializers.BuildComment(comment),
	}
}
