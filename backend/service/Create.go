package service

import (
	"comment/e"
	"comment/model"
	"comment/serializer"
)

type CreateCommentService struct {
    ReplyToID uint   `json:"reply_to_id" form:"reply_to_id"`
    ArticleID uint   `json:"Article_id" form:"Article_id" binding:"required"`
    ParentID  uint   `json:"parent_id" form:"parent_id"`
    RootID    uint   `json:"root_id" form:"root_id"`
    Content   string `json:"content" form:"content" binding:"required,min=3"`
}

func (service *CreateCommentService) Create(user *model.User) *serializer.Response {
    comment := model.Comment{
        UserID:    user.ID,
        ReplyToID: service.ReplyToID,
        ArticleID: service.ArticleID,
        ParentID:  service.ParentID,
        RootID:    service.RootID,
        Content:   service.Content,
    }
    if err := model.DB.Create(&comment).Error; err != nil {
        return &serializer.Response{
            Code:  e.CREATE_ERROR,
            Msg: e.GetMsg(e.CREATE_ERROR),
            Error:   err.Error(),
        }
    }

    // 返回当前评论的用户和   replyTo用户(如果是回复的话)
    comment.User = *user
    if comment.ReplyToID != 0 {
        var replyToUser model.User
        if err := model.DB.Find(&replyToUser, service.ReplyToID).Error; err != nil {
            return &serializer.Response{
                Code:e.SELECT_ERROR,
                Msg:e.GetMsg(e.SELECT_ERROR),
                Error:err.Error(),
            }
        }
        comment.ReplyTo = replyToUser
        return &serializer.Response{
            Code:  e.SUCCESS,
            Msg: e.GetMsg(e.SUCCESS),
            Data:    serializer.BuildReply(comment),
        }
    }

    return &serializer.Response{
        Code:  e.SUCCESS,
        Msg: e.GetMsg(e.SUCCESS),
        Data:    serializer.BuildComment(comment),
    }
}