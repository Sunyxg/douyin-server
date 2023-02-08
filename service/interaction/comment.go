package interactionimport

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"context"

	. "ByteTech-7355608/douyin-server/pkg/configs"
)

const (
	KAddType    = 1
	KDeleteType = 2
)

func (s *Service) CommentAction(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	resp = interaction.NewDouyinCommentActionResponse()
	switch req.ActionType {
	case KAddType:
		comment, err := s.dao.User.AddComment(ctx, req)
		if err != nil {
			Log.Errorf("add comment err: %v", err)
			return nil, err
		}
		resp.Comment = &comment
	case KDeleteType:
		err = s.dao.User.DeleteComment(ctx, req)
		if err != nil {
			Log.Errorf("delete comment err: %v", err)
			return nil, err
		}
	}

	return
}
