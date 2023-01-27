package model

//
//import "context"
//
//type UserService interface {
//	Login(ctx context.Context, req *DouyinUserLoginRequest) (r *DouyinUserLoginResponse, err error)
//
//	Register(ctx context.Context, req *DouyinUserRegisterRequest) (r *DouyinUserRegisterResponse, err error)
//
//	Info(ctx context.Context, req *DouyinUserRequest) (r *DouyinUserResponse, err error)
//
//	Action(ctx context.Context, req *DouyinRelationActionRequest) (r *DouyinRelationActionResponse, err error)
//
//	FollowList(ctx context.Context, req *DouyinRelationFollowListRequest) (r *DouyinRelationFollowListResponse, err error)
//
//	FollowerList(ctx context.Context, req *DouyinRelationFollowerListRequest) (r *DouyinRelationFollowerListResponse, err error)
//
//	FriendList(ctx context.Context, req *DouyinRelationFriendListRequest) (r *DouyinRelationFriendListResponse, err error)
//}
//
//type DouyinRelationFriendListRequest struct {
//	UserId int64  `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
//	Token  string `thrift:"token,2" frugal:"2,default,string" json:"token"`
//}
//
//func NewDouyinRelationFriendListRequest() *DouyinRelationFriendListRequest {
//	return &DouyinRelationFriendListRequest{}
//}
//
//func (p *DouyinRelationFriendListRequest) InitDefault() {
//	*p = DouyinRelationFriendListRequest{}
//}
//
//func (p *DouyinRelationFriendListRequest) GetUserId() (v int64) {
//	return p.UserId
//}
//
//func (p *DouyinRelationFriendListRequest) GetToken() (v string) {
//	return p.Token
//}
//func (p *DouyinRelationFriendListRequest) SetUserId(val int64) {
//	p.UserId = val
//}
//func (p *DouyinRelationFriendListRequest) SetToken(val string) {
//	p.Token = val
//}
//
//type DouyinRelationFriendListResponse struct {
//	StatusCode int32   `frugal:"1,default,i32" json:"status_code"`
//	StatusMsg  string  `frugal:"2,default,string" json:"status_msg"`
//	UserList   []*User `frugal:"3,default,list<User>" json:"user_list"`
//}
//
//func NewDouyinRelationFriendListResponse() *DouyinRelationFriendListResponse {
//	return &DouyinRelationFriendListResponse{}
//}
//
//func (p *DouyinRelationFriendListResponse) InitDefault() {
//	*p = DouyinRelationFriendListResponse{}
//}
//
//func (p *DouyinRelationFriendListResponse) GetStatusCode() (v int32) {
//	return p.StatusCode
//}
//
//func (p *DouyinRelationFriendListResponse) GetStatusMsg() (v string) {
//	return p.StatusMsg
//}
//
//func (p *DouyinRelationFriendListResponse) GetUserList() (v []*User) {
//	return p.UserList
//}
//func (p *DouyinRelationFriendListResponse) SetStatusCode(val int32) {
//	p.StatusCode = val
//}
//func (p *DouyinRelationFriendListResponse) SetStatusMsg(val string) {
//	p.StatusMsg = val
//}
//func (p *DouyinRelationFriendListResponse) SetUserList(val []*User) {
//	p.UserList = val
//}
//
//type DouyinRelationFollowerListRequest struct {
//	UserId int64  `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
//	Token  string `thrift:"token,2" frugal:"2,default,string" json:"token"`
//}
//
//func NewDouyinRelationFollowerListRequest() *DouyinRelationFollowerListRequest {
//	return &DouyinRelationFollowerListRequest{}
//}
//
//func (p *DouyinRelationFollowerListRequest) InitDefault() {
//	*p = DouyinRelationFollowerListRequest{}
//}
//
//func (p *DouyinRelationFollowerListRequest) GetUserId() (v int64) {
//	return p.UserId
//}
//
//func (p *DouyinRelationFollowerListRequest) GetToken() (v string) {
//	return p.Token
//}
//func (p *DouyinRelationFollowerListRequest) SetUserId(val int64) {
//	p.UserId = val
//}
//func (p *DouyinRelationFollowerListRequest) SetToken(val string) {
//	p.Token = val
//}
//
//type DouyinRelationFollowerListResponse struct {
//	StatusCode int32   `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
//	StatusMsg  string  `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
//	UserList   []*User `thrift:"user_list,3" frugal:"3,default,list<User>" json:"user_list"`
//}
//
//func NewDouyinRelationFollowerListResponse() *DouyinRelationFollowerListResponse {
//	return &DouyinRelationFollowerListResponse{}
//}
//
//func (p *DouyinRelationFollowerListResponse) InitDefault() {
//	*p = DouyinRelationFollowerListResponse{}
//}
//
//func (p *DouyinRelationFollowerListResponse) GetStatusCode() (v int32) {
//	return p.StatusCode
//}
//
//func (p *DouyinRelationFollowerListResponse) GetStatusMsg() (v string) {
//	return p.StatusMsg
//}
//
//func (p *DouyinRelationFollowerListResponse) GetUserList() (v []*User) {
//	return p.UserList
//}
//func (p *DouyinRelationFollowerListResponse) SetStatusCode(val int32) {
//	p.StatusCode = val
//}
//func (p *DouyinRelationFollowerListResponse) SetStatusMsg(val string) {
//	p.StatusMsg = val
//}
//func (p *DouyinRelationFollowerListResponse) SetUserList(val []*User) {
//	p.UserList = val
//}
//
//type DouyinRelationFollowListRequest struct {
//	UserId int64  `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
//	Token  string `thrift:"token,2" frugal:"2,default,string" json:"token"`
//}
//
//func NewDouyinRelationFollowListRequest() *DouyinRelationFollowListRequest {
//	return &DouyinRelationFollowListRequest{}
//}
//
//func (p *DouyinRelationFollowListRequest) InitDefault() {
//	*p = DouyinRelationFollowListRequest{}
//}
//
//func (p *DouyinRelationFollowListRequest) GetUserId() (v int64) {
//	return p.UserId
//}
//
//func (p *DouyinRelationFollowListRequest) GetToken() (v string) {
//	return p.Token
//}
//func (p *DouyinRelationFollowListRequest) SetUserId(val int64) {
//	p.UserId = val
//}
//func (p *DouyinRelationFollowListRequest) SetToken(val string) {
//	p.Token = val
//}
//
//type DouyinRelationFollowListResponse struct {
//	StatusCode int32   `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
//	StatusMsg  string  `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
//	UserList   []*User `thrift:"user_list,3" frugal:"3,default,list<User>" json:"user_list"`
//}
//
//func NewDouyinRelationFollowListResponse() *DouyinRelationFollowListResponse {
//	return &DouyinRelationFollowListResponse{}
//}
//
//func (p *DouyinRelationFollowListResponse) InitDefault() {
//	*p = DouyinRelationFollowListResponse{}
//}
//
//func (p *DouyinRelationFollowListResponse) GetStatusCode() (v int32) {
//	return p.StatusCode
//}
//
//func (p *DouyinRelationFollowListResponse) GetStatusMsg() (v string) {
//	return p.StatusMsg
//}
//
//func (p *DouyinRelationFollowListResponse) GetUserList() (v []*User) {
//	return p.UserList
//}
//func (p *DouyinRelationFollowListResponse) SetStatusCode(val int32) {
//	p.StatusCode = val
//}
//func (p *DouyinRelationFollowListResponse) SetStatusMsg(val string) {
//	p.StatusMsg = val
//}
//func (p *DouyinRelationFollowListResponse) SetUserList(val []*User) {
//	p.UserList = val
//}
//
//type DouyinRelationActionRequest struct {
//	Token      string `thrift:"token,1" frugal:"1,default,string" json:"token"`
//	ToUserId   int64  `thrift:"to_user_id,2" frugal:"2,default,i64" json:"to_user_id"`
//	ActionType int32  `thrift:"action_type,3" frugal:"3,default,i32" json:"action_type"`
//}
//
//func NewDouyinRelationActionRequest() *DouyinRelationActionRequest {
//	return &DouyinRelationActionRequest{}
//}
//
//func (p *DouyinRelationActionRequest) InitDefault() {
//	*p = DouyinRelationActionRequest{}
//}
//
//func (p *DouyinRelationActionRequest) GetToken() (v string) {
//	return p.Token
//}
//
//func (p *DouyinRelationActionRequest) GetToUserId() (v int64) {
//	return p.ToUserId
//}
//
//func (p *DouyinRelationActionRequest) GetActionType() (v int32) {
//	return p.ActionType
//}
//func (p *DouyinRelationActionRequest) SetToken(val string) {
//	p.Token = val
//}
//func (p *DouyinRelationActionRequest) SetToUserId(val int64) {
//	p.ToUserId = val
//}
//func (p *DouyinRelationActionRequest) SetActionType(val int32) {
//	p.ActionType = val
//}
//
//type DouyinRelationActionResponse struct {
//	StatusCode int32  `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
//	StatusMsg  string `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
//}
//
//func NewDouyinRelationActionResponse() *DouyinRelationActionResponse {
//	return &DouyinRelationActionResponse{}
//}
//
//func (p *DouyinRelationActionResponse) InitDefault() {
//	*p = DouyinRelationActionResponse{}
//}
//
//func (p *DouyinRelationActionResponse) GetStatusCode() (v int32) {
//	return p.StatusCode
//}
//
//func (p *DouyinRelationActionResponse) GetStatusMsg() (v string) {
//	return p.StatusMsg
//}
//func (p *DouyinRelationActionResponse) SetStatusCode(val int32) {
//	p.StatusCode = val
//}
//func (p *DouyinRelationActionResponse) SetStatusMsg(val string) {
//	p.StatusMsg = val
//}
//
//type DouyinUserRegisterRequest struct {
//	Username string `thrift:"username,1" frugal:"1,default,string" json:"username"`
//	Password string `thrift:"password,2" frugal:"2,default,string" json:"password"`
//}
//
//func NewDouyinUserRegisterRequest() *DouyinUserRegisterRequest {
//	return &DouyinUserRegisterRequest{}
//}
//
//func (p *DouyinUserRegisterRequest) InitDefault() {
//	*p = DouyinUserRegisterRequest{}
//}
//
//func (p *DouyinUserRegisterRequest) GetUsername() (v string) {
//	return p.Username
//}
//
//func (p *DouyinUserRegisterRequest) GetPassword() (v string) {
//	return p.Password
//}
//func (p *DouyinUserRegisterRequest) SetUsername(val string) {
//	p.Username = val
//}
//func (p *DouyinUserRegisterRequest) SetPassword(val string) {
//	p.Password = val
//}
//
//type DouyinUserRegisterResponse struct {
//	StatusCode int32  `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
//	StatusMsg  string `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
//	UserId     int64  `thrift:"user_id,3" frugal:"3,default,i64" json:"user_id"`
//	Token      string `thrift:"token,4" frugal:"4,default,string" json:"token"`
//}
//
//func NewDouyinUserRegisterResponse() *DouyinUserRegisterResponse {
//	return &DouyinUserRegisterResponse{}
//}
//
//func (p *DouyinUserRegisterResponse) InitDefault() {
//	*p = DouyinUserRegisterResponse{}
//}
//
//func (p *DouyinUserRegisterResponse) GetStatusCode() (v int32) {
//	return p.StatusCode
//}
//
//func (p *DouyinUserRegisterResponse) GetStatusMsg() (v string) {
//	return p.StatusMsg
//}
//
//func (p *DouyinUserRegisterResponse) GetUserId() (v int64) {
//	return p.UserId
//}
//
//func (p *DouyinUserRegisterResponse) GetToken() (v string) {
//	return p.Token
//}
//func (p *DouyinUserRegisterResponse) SetStatusCode(val int32) {
//	p.StatusCode = val
//}
//func (p *DouyinUserRegisterResponse) SetStatusMsg(val string) {
//	p.StatusMsg = val
//}
//func (p *DouyinUserRegisterResponse) SetUserId(val int64) {
//	p.UserId = val
//}
//func (p *DouyinUserRegisterResponse) SetToken(val string) {
//	p.Token = val
//}
//
//type DouyinUserLoginRequest struct {
//	Username string `thrift:"username,1" frugal:"1,default,string" json:"username"`
//	Password string `thrift:"password,2" frugal:"2,default,string" json:"password"`
//}
//
//func NewDouyinUserLoginRequest() *DouyinUserLoginRequest {
//	return &DouyinUserLoginRequest{}
//}
//
//func (p *DouyinUserLoginRequest) InitDefault() {
//	*p = DouyinUserLoginRequest{}
//}
//
//func (p *DouyinUserLoginRequest) GetUsername() (v string) {
//	return p.Username
//}
//
//func (p *DouyinUserLoginRequest) GetPassword() (v string) {
//	return p.Password
//}
//func (p *DouyinUserLoginRequest) SetUsername(val string) {
//	p.Username = val
//}
//func (p *DouyinUserLoginRequest) SetPassword(val string) {
//	p.Password = val
//}
//
//type DouyinUserLoginResponse struct {
//	StatusCode int32  `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
//	StatusMsg  string `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
//	UserId     int64  `thrift:"user_id,3" frugal:"3,default,i64" json:"user_id"`
//	Token      string `thrift:"token,4" frugal:"4,default,string" json:"token"`
//}
//
//func NewDouyinUserLoginResponse() *DouyinUserLoginResponse {
//	return &DouyinUserLoginResponse{}
//}
//
//func (p *DouyinUserLoginResponse) InitDefault() {
//	*p = DouyinUserLoginResponse{}
//}
//
//func (p *DouyinUserLoginResponse) GetStatusCode() (v int32) {
//	return p.StatusCode
//}
//
//func (p *DouyinUserLoginResponse) GetStatusMsg() (v string) {
//	return p.StatusMsg
//}
//
//func (p *DouyinUserLoginResponse) GetUserId() (v int64) {
//	return p.UserId
//}
//
//func (p *DouyinUserLoginResponse) GetToken() (v string) {
//	return p.Token
//}
//func (p *DouyinUserLoginResponse) SetStatusCode(val int32) {
//	p.StatusCode = val
//}
//func (p *DouyinUserLoginResponse) SetStatusMsg(val string) {
//	p.StatusMsg = val
//}
//func (p *DouyinUserLoginResponse) SetUserId(val int64) {
//	p.UserId = val
//}
//func (p *DouyinUserLoginResponse) SetToken(val string) {
//	p.Token = val
//}
//
//type DouyinUserRequest struct {
//	UserId int64  `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
//	Token  string `thrift:"token,2" frugal:"2,default,string" json:"token"`
//}
//
//func NewDouyinUserRequest() *DouyinUserRequest {
//	return &DouyinUserRequest{}
//}
//
//func (p *DouyinUserRequest) InitDefault() {
//	*p = DouyinUserRequest{}
//}
//
//func (p *DouyinUserRequest) GetUserId() (v int64) {
//	return p.UserId
//}
//
//func (p *DouyinUserRequest) GetToken() (v string) {
//	return p.Token
//}
//func (p *DouyinUserRequest) SetUserId(val int64) {
//	p.UserId = val
//}
//func (p *DouyinUserRequest) SetToken(val string) {
//	p.Token = val
//}
//
//type DouyinUserResponse struct {
//	StatusCode int32  `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
//	StatusMsg  string `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
//	User       *User  `thrift:"user,3" frugal:"3,default,User" json:"user"`
//}
//
//func NewDouyinUserResponse() *DouyinUserResponse {
//	return &DouyinUserResponse{}
//}
//
//func (p *DouyinUserResponse) InitDefault() {
//	*p = DouyinUserResponse{}
//}
//
//func (p *DouyinUserResponse) GetStatusCode() (v int32) {
//	return p.StatusCode
//}
//
//func (p *DouyinUserResponse) GetStatusMsg() (v string) {
//	return p.StatusMsg
//}
//
//func (p *DouyinUserResponse) GetUser() (v *User) {
//	return p.User
//}
//func (p *DouyinUserResponse) SetStatusCode(val int32) {
//	p.StatusCode = val
//}
//func (p *DouyinUserResponse) SetStatusMsg(val string) {
//	p.StatusMsg = val
//}
//func (p *DouyinUserResponse) SetUser(val *User) {
//	p.User = val
//}
//
//type User struct {
//	Id            int64  `thrift:"id,1" frugal:"1,default,i64" json:"id"`
//	Name          string `thrift:"name,2" frugal:"2,default,string" json:"name"`
//	FollowCount   int64  `thrift:"follow_count,3" frugal:"3,default,i64" json:"follow_count"`
//	FollowerCount int64  `thrift:"follower_count,4" frugal:"4,default,i64" json:"follower_count"`
//	IsFollow      bool   `thrift:"is_follow,5" frugal:"5,default,bool" json:"is_follow"`
//}
//
//func NewUser() *User {
//	return &User{}
//}
//
//func (p *User) InitDefault() {
//	*p = User{}
//}
//
//func (p *User) GetId() (v int64) {
//	return p.Id
//}
//
//func (p *User) GetName() (v string) {
//	return p.Name
//}
//
//func (p *User) GetFollowCount() (v int64) {
//	return p.FollowCount
//}
//
//func (p *User) GetFollowerCount() (v int64) {
//	return p.FollowerCount
//}
//
//func (p *User) GetIsFollow() (v bool) {
//	return p.IsFollow
//}
//func (p *User) SetId(val int64) {
//	p.Id = val
//}
//func (p *User) SetName(val string) {
//	p.Name = val
//}
//func (p *User) SetFollowCount(val int64) {
//	p.FollowCount = val
//}
//func (p *User) SetFollowerCount(val int64) {
//	p.FollowerCount = val
//}
//func (p *User) SetIsFollow(val bool) {
//	p.IsFollow = val
//}
