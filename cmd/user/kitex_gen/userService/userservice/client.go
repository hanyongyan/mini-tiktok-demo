// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	userService "mini-tiktok-hanyongyan/cmd/user/kitex_gen/userService"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, Req *userService.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *userService.DouyinUserLoginResponse, err error)
	Register(ctx context.Context, Req *userService.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *userService.DouyinUserRegisterResponse, err error)
	Info(ctx context.Context, Req *userService.DouyinUserRequest, callOptions ...callopt.Option) (r *userService.DouyinUserResponse, err error)
	Action(ctx context.Context, Req *userService.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *userService.DouyinRelationActionResponse, err error)
	FollowList(ctx context.Context, Req *userService.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *userService.DouyinRelationFollowListResponse, err error)
	FollowerList(ctx context.Context, Req *userService.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *userService.DouyinRelationFollowerListResponse, err error)
	FriendList(ctx context.Context, Req *userService.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *userService.DouyinRelationFriendListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Login(ctx context.Context, Req *userService.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *userService.DouyinUserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}

func (p *kUserServiceClient) Register(ctx context.Context, Req *userService.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *userService.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kUserServiceClient) Info(ctx context.Context, Req *userService.DouyinUserRequest, callOptions ...callopt.Option) (r *userService.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Info(ctx, Req)
}

func (p *kUserServiceClient) Action(ctx context.Context, Req *userService.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *userService.DouyinRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Action(ctx, Req)
}

func (p *kUserServiceClient) FollowList(ctx context.Context, Req *userService.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *userService.DouyinRelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowList(ctx, Req)
}

func (p *kUserServiceClient) FollowerList(ctx context.Context, Req *userService.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *userService.DouyinRelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowerList(ctx, Req)
}

func (p *kUserServiceClient) FriendList(ctx context.Context, Req *userService.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *userService.DouyinRelationFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendList(ctx, Req)
}
