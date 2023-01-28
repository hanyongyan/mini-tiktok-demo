package main

import (
	"context"
	"errors"
	"fmt"
	"mini-tiktok-hanyongyan/cmd/user/kitex_gen/userservice"
	"mini-tiktok-hanyongyan/pkg/dal/model"
	"mini-tiktok-hanyongyan/pkg/dal/query"
	"mini-tiktok-hanyongyan/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *userservice.DouyinUserLoginRequest) (resp *userservice.DouyinUserLoginResponse, err error) {
	q := query.Q.TUser

	// 查询当前用户是否存在
	user, err := q.WithContext(ctx).Where(q.Name.Eq(req.Username)).First()
	if err != nil {
		err = errors.New("用户不存在")
		return
	} else if user.Password != req.Password {
		// 说明密码错误
		err = errors.New("密码错误！")
		return
	}
	token, err := utils.CreateToken(user.ID)
	if err != nil {
		err = errors.New("token 生成错误")
		return
	}
	// 由于返回值是进行定义了一个指针，我们不能直接对其中的属性进行赋值，使用下面的用法会报错
	//resp.StatusCode = 0
	//resp.StatusMsg = "登陆成功"
	//resp.Token = token
	//resp.UserId = user.ID
	resp = &userservice.DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "登陆成功",
		UserId:     user.ID,
		Token:      token,
	}
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userservice.DouyinUserRegisterRequest) (resp *userservice.DouyinUserRegisterResponse, err error) {
	q := query.Q.TUser

	newUser := &model.TUser{Name: req.Username, Password: req.Password}

	// 查询数据库查看用户是否存在
	userIsExist, err := q.WithContext(ctx).Where(q.Name.Eq(newUser.Name)).First()
	if userIsExist != nil {
		err = fmt.Errorf("用户已存在，注册失败：%w", err)
		return
	}
	// 创建用户
	err = q.WithContext(ctx).Create(newUser)
	if err != nil {
		err = fmt.Errorf("注册失败：%w", err)
		return
	}
	//使用 jwt 生成 token
	token, err := utils.CreateToken(newUser.ID)
	if err != nil {
		err = fmt.Errorf("token生成失败：%w", err)
		return
	}
	resp = &userservice.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     newUser.ID,
		Token:      token,
	}
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *userservice.DouyinUserRequest) (resp *userservice.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// Action implements the UserServiceImpl interface.
func (s *UserServiceImpl) Action(ctx context.Context, req *userservice.DouyinRelationActionRequest) (resp *userservice.DouyinRelationActionResponse, err error) {
	// 关注操作
	q := query.Q.TFollow
	resp = &userservice.DouyinRelationActionResponse{}

	claims, flag := utils.CheckToken(req.Token)
	// 解析 token 失败
	if !flag {
		err = errors.New("token is expired")
		return
	}
	follow := &model.TFollow{
		UserID:     claims.UserId,
		FollowerID: req.ToUserId,
	}
	if req.ActionType == 1 {
		// 关注操作
		resultFollow, err := q.WithContext(ctx).Where(q.UserID.Eq(follow.UserID)).Where(q.FollowerID.Eq(follow.FollowerID)).First()
		// 说明还没有关注过
		if err != nil && err.Error() == "record not found" {
			err = q.WithContext(ctx).Create(follow)
			if err != nil {
				return nil, err
			}
			// 进行到此步说明 添加关注成功
			resp.StatusCode = 0
			resp.StatusMsg = "关注成功"
			return resp, nil
		}
		// 说明已经关注过
		if resultFollow != nil {
			err = errors.New("请勿重复关注！")
			return nil, err
		}

	} else {
		// 取消关注操作
		// 先进行是否存在这样一种关注关系
		_, err := q.WithContext(ctx).Where(q.UserID.Eq(follow.UserID)).Where(q.FollowerID.Eq(follow.FollowerID)).First()
		// 查询不到用户
		if err != nil && err.Error() == "record not found" {
			err = errors.New("请勿重复取消关注")
			return nil, err
		}

		// 进行删除数据库中的数据
		_, err = q.WithContext(ctx).Where(q.UserID.Eq(follow.UserID)).Where(q.FollowerID.Eq(follow.FollowerID)).Delete()
		if err != nil {
			return nil, err
		}
		resp.StatusMsg = "取消关注成功"
		resp.StatusCode = 0
		return resp, nil
	}
	return
}

// FollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowList(ctx context.Context, req *userservice.DouyinRelationFollowListRequest) (resp *userservice.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowerList(ctx context.Context, req *userservice.DouyinRelationFollowerListRequest) (resp *userservice.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FriendList(ctx context.Context, req *userservice.DouyinRelationFriendListRequest) (resp *userservice.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
