package service

import (
	"mini-tiktok/dao"
	"time"
)

type CommentServiceImpl struct {
	Usi UserServiceImpl
}

// GetCommentListByVideoID 返回对应视频下的评论
func (csi *CommentServiceImpl) GetCommentListByVideoID(videoID int64, currUserID int64) ([]*Comment, error) {
	daoComments, err := dao.GetCommentByVideoID(videoID)
	if err != nil {
		return nil, err
	}

	var res []*Comment
	for _, c := range daoComments {
		to := ToComment(c)
		to.User, err = csi.Usi.GetUserInfoById(c.ID, currUserID)
		if err != nil {
			return nil, err
		}
		res = append(res, to)
	}
	return res, nil
}

// CreateComment 创建评论, 返回评论 id
func (csi *CommentServiceImpl) CreateComment(userID int64, videoID int64, content string) (int64, error) {
	return dao.InsertComment(&dao.Comment{
		UserID:     userID,
		VideoID:    videoID,
		Content:    content,
		CommitTime: time.Now().Unix(),
	})
}

// DeleteCommentByID 删除评论
func (csi *CommentServiceImpl) DeleteCommentByID(id int64) error {
	return dao.DeleteCommentByID(id)
}
