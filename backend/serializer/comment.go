package serializer

import (
	"comment/model"
	"time"
)

type CommentSerializer struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	ArticleID uint              `json:"Article_id"`
	Content   string            `json:"content"`
	CreatedAt string            `json:"created_at"`
	Replys    []ReplySerializer `json:"replys"`
}

type ReplySerializer struct {
	ID        uint           `json:"id"`
	User      UserSerializer `json:"user"`
	ReplyTo   UserSerializer `json:"reply_to"`
	ParentID  uint           `json:"parent_id"`
	RootID    uint           `json:"root_id"`
	Content   string         `json:"content"`
	CreatedAt string         `json:"created_at"`
}

// 序列化回复
func BuildReply(item model.Comment) ReplySerializer {
	return ReplySerializer{
		ID:        item.ID,
		User:      BuildUser(item.User),
		ReplyTo:   BuildUser(item.ReplyTo),
		ParentID:  item.ParentID,
		RootID:    item.RootID,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Format(time.Now().Format("2006-01-02 15:04:05")),
	}
}

// 序列化评论
func BuildComment(item model.Comment) CommentSerializer {
	var replys []ReplySerializer
	if len(item.Replys) != 0 {
		for _, reply := range item.Replys {
			replys = append(replys, BuildReply(reply))
		}
	}
	return CommentSerializer{
		ID:        item.ID,
		User:      BuildUser(item.User),
		ArticleID: item.ArticleID,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Format(time.Now().Format("2006-01-02 15:04:05")),
		Replys:    replys,
	}
}

func BuildComments(items []model.Comment) []CommentSerializer {
	var comments []CommentSerializer
	for _, item := range items {
		comments = append(comments, BuildComment(item))
	}
	return comments
}
