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
	resp = &userservice.DouyinUserResponse{}
	queryUser := query.Q.TUser
	queryFollow := query.Q.TFollow

	// 进行查询当前用户信息
	user, err := queryUser.WithContext(ctx).Where(queryUser.ID.Eq(req.UserId)).First()
	if err != nil {
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询用户信息成功"
	// 用户信息进行赋值
	resp.User.Id = user.ID
	resp.User.Name = user.Name
	resp.User.FollowerCount = user.FollowerCount
	resp.User.FollowCount = user.FollowCount
	// 进行查询是否关注
	claims, flag := utils.CheckToken(req.Token)
	if !flag {
		return nil, errors.New("用户验证错误")
	}
	_, err = queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(claims.UserId)).Where(queryFollow.FollowerID.Eq(user.ID)).First()
	if err != nil {
		// 如果查询不到当前关注信息，说明未关注此用户
		if err.Error() == "record not found" {
			resp.User.IsFollow = false
			return resp, nil
		}
		return nil, err
	}
	resp.User.IsFollow = true
	return resp, nil
}

// Action implements the UserServiceImpl interface.
func (s *UserServiceImpl) Action(ctx context.Context, req *userservice.DouyinRelationActionRequest) (resp *userservice.DouyinRelationActionResponse, err error) {
	// 关注操作
	queryFollow := query.Q.TFollow
	queryFriend := query.Q.TFriend
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
		resultFollow, err := queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(follow.UserID)).
			Where(queryFollow.FollowerID.Eq(follow.FollowerID)).First()
		// 说明还没有关注过
		if err != nil && err.Error() == "record not found" {
			err = queryFollow.WithContext(ctx).Create(follow)
			if err != nil {
				return nil, err
			}
			// 进行到此步说明 添加关注成功
			whetherToCare, _ := queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(follow.FollowerID)).
				Where(queryFollow.FollowerID.Eq(follow.UserID)).First()
			// 所关注的用户关注了自己
			// 添加好友数据
			if whetherToCare != nil {
				_ = queryFriend.WithContext(ctx).Create(&model.TFriend{
					UserID:   follow.UserID,
					FriendID: follow.FollowerID,
				})
			}

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
		_, err := queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(follow.UserID)).Where(queryFollow.FollowerID.Eq(follow.FollowerID)).First()
		// 查询不到用户
		if err != nil && err.Error() == "record not found" {
			err = errors.New("请勿重复取消关注")
			return nil, err
		}

		// 进行删除数据库中的数据
		_, err = queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(follow.UserID)).Where(queryFollow.FollowerID.Eq(follow.FollowerID)).Delete()
		if err != nil {
			return nil, err
		}
		// 查看是否存好友关系，如果存在好友关系，将好友关系从数据库中进行删除
		isFriend, _ := queryFriend.WithContext(ctx).Where(queryFriend.FriendID.Eq(follow.FollowerID), queryFriend.UserID.Eq(follow.UserID)).
			Or(queryFriend.FriendID.Eq(follow.UserID), queryFriend.UserID.Eq(follow.FollowerID)).Find()
		// 说明存在好友关系
		if isFriend != nil {
			// 进行删除好友关系
			_, _ = queryFriend.WithContext(ctx).Where(queryFriend.FriendID.Eq(follow.FollowerID), queryFriend.UserID.Eq(follow.UserID)).
				Or(queryFriend.FriendID.Eq(follow.UserID), queryFriend.UserID.Eq(follow.FollowerID)).Delete()
		}
		resp.StatusMsg = "取消关注成功"
		resp.StatusCode = 0
		return resp, nil
	}
	return
}

// FollowList implements the UserServiceImpl interface. 关注列表操作
func (s *UserServiceImpl) FollowList(ctx context.Context, req *userservice.DouyinRelationFollowListRequest) (resp *userservice.DouyinRelationFollowListResponse, err error) {
	resp = &userservice.DouyinRelationFollowListResponse{}
	// 用于查询关注的用户 id
	queryFollow := query.Q.TFollow
	// 根据关注的用户id查询出所有的关注用户
	queryUser := query.Q.TUser
	// 只检索关注用户的id
	follows, err := queryFollow.WithContext(ctx).Select(queryFollow.FollowerID).Where(queryFollow.UserID.Eq(req.UserId)).Find()
	if err != nil {
		return
	}
	// 用来保存用户id
	followUserId := make([]int64, len(follows))
	for i, follow := range follows {
		followUserId[i] = follow.FollowerID
	}
	// 进行查询用户信息
	t_users, err := queryUser.WithContext(ctx).Where(queryUser.ID.In(followUserId...)).Find()
	if err != nil {
		return
	}
	//users := make([]userservice.User, len(t_users))

	for _, u := range t_users {
		var user userservice.User

		user.Id = u.ID
		user.FollowCount = u.FollowCount
		user.FollowerCount = u.FollowerCount
		user.Name = u.Name
		user.IsFollow = true
		resp.UserList = append(resp.UserList, &user)
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询关注成功"
	return
}

// FollowerList implements the UserServiceImpl interface. 粉丝列表
func (s *UserServiceImpl) FollowerList(ctx context.Context, req *userservice.DouyinRelationFollowerListRequest) (resp *userservice.DouyinRelationFollowerListResponse, err error) {
	resp = &userservice.DouyinRelationFollowerListResponse{}
	// 用于查询粉丝的用户 id
	queryFollow := query.Q.TFollow
	// 根据粉丝的id查询出所有的粉丝
	queryUser := query.Q.TUser
	// 检索粉丝的id
	followers, err := queryFollow.WithContext(ctx).Select(queryFollow.UserID).Where(queryFollow.FollowerID.Eq(req.UserId)).Find()
	if err != nil {
		return
	}
	// 用来绑定粉丝id
	followersId := make([]int64, len(followers))
	for i, follower := range followers {
		followersId[i] = follower.UserID
	}
	//进行查询粉丝用户信息
	t_users, err := queryUser.WithContext(ctx).Where(queryUser.ID.In(followersId...)).Find()
	if err != nil {
		return
	}
	for _, tUser := range t_users {
		var user userservice.User
		user.Id = tUser.ID
		user.FollowCount = tUser.FollowCount
		user.FollowerCount = tUser.FollowerCount
		user.Name = tUser.Name
		// 进行查询当前用户是否关注了此粉丝
		_, err := queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(req.UserId)).Where(queryFollow.FollowerID.Eq(user.Id)).First()
		if err != nil && err.Error() == "record not found" {
			user.IsFollow = false
		} else {
			user.IsFollow = true
		}
		resp.UserList = append(resp.UserList, &user)
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询粉丝成功"
	return
}

// FriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FriendList(ctx context.Context, req *userservice.DouyinRelationFriendListRequest) (resp *userservice.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
