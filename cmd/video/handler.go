package main

import (
	"context"
	"errors"
	"fmt"
	"mini-tiktok-hanyongyan/cmd/video/cos"
	videoservice "mini-tiktok-hanyongyan/cmd/video/kitex_gen/videoService"
	"mini-tiktok-hanyongyan/pkg/dal/model"
	"mini-tiktok-hanyongyan/pkg/dal/query"
	"mini-tiktok-hanyongyan/pkg/utils"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videoservice.DouyinPublishActionRequest) (resp *videoservice.DouyinPublishActionResponse, err error) {
	video := query.Q.TVideo
	claims, flag := utils.CheckToken(req.Token)
	if !flag {
		return nil, errors.New("token is expired")
	}
	userId := claims.UserId
	// 将 int8 数组转为 byte 数组
	data := []byte(string(req.Data))
	// 将视频上传到 cos，返回上传路径及文件名
	flag, videoPath, photoPath := cos.SaveUploadedFile(ctx, data)
	if !flag {
		return nil, errors.New("file processing failed")
	}
	fmt.Println(".................")
	err = video.WithContext(ctx).Create(&model.TVideo{
		AuthorID:      userId,
		PlayURL:       videoPath,
		CoverURL:      photoPath,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         req.Title,
	})
	if err != nil {
		return nil, err
	}
	resp = &videoservice.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "视频上传成功",
	}
	return resp, nil
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videoservice.DouyinFeedRequest) (resp *videoservice.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *videoservice.DouyinPublishListRequest) (resp *videoservice.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *videoservice.DouyinFavoriteActionRequest) (resp *videoservice.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteList(ctx context.Context, req *videoservice.DouyinFavoriteListRequest) (resp *videoservice.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentAction(ctx context.Context, req *videoservice.DouyinCommentActionRequest) (resp *videoservice.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentList(ctx context.Context, req *videoservice.DouyinCommentListRequest) (resp *videoservice.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
