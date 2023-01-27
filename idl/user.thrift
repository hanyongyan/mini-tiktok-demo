namespace go userservice

service UserService {
    douyin_user_login_response Login(1: douyin_user_login_request req);
    douyin_user_register_response Register(1: douyin_user_register_request req);
    douyin_user_response Info(1: douyin_user_request req);
    douyin_relation_action_response Action(1: douyin_relation_action_request req);
    douyin_relation_follow_list_response FollowList(1: douyin_relation_follow_list_request req);
    douyin_relation_follower_list_response FollowerList(1: douyin_relation_follower_list_request req);
    douyin_relation_friend_list_response FriendList(1: douyin_relation_friend_list_request req);
}

struct douyin_relation_friend_list_request {
    1: i64 user_id // 用户id​
    2: string token // 用户鉴权token​
}

struct douyin_relation_friend_list_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: list<User> user_list  // 用户列表​
}
struct douyin_relation_follower_list_request {
    1: i64 user_id  // 用户id​
    2: string token  // 用户鉴权token​
}

struct douyin_relation_follower_list_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: list<User> user_list  // 用户列表​
}
struct douyin_relation_follow_list_request {
    1: i64 user_id  // 用户id​
    2:string token  // 用户鉴权token​
}

struct douyin_relation_follow_list_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: list<User> user_list  // 用户信息列表​
}
struct douyin_relation_action_request {
    1: string token  // 用户鉴权token​
    2: i64 to_user_id  // 对方用户id​
    3: i32 action_type  // 1-关注，2-取消关注​
}

struct douyin_relation_action_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
}
struct douyin_user_register_request {
    1: string username  // 注册用户名，最长32个字符​
    2: string password  // 密码，最长32个字符​
}

struct douyin_user_register_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: i64 user_id  // 用户id​
    4: string token  // 用户鉴权token​
}
struct douyin_user_login_request {
    1: string username  // 登录用户名​
    2: string password  // 登录密码​
}

struct douyin_user_login_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: i64 user_id  // 用户id​
    4: string token  // 用户鉴权token​
}
struct douyin_user_request {
    1: i64 user_id  // 用户id​
    2: string token  // 用户鉴权token​
}

struct douyin_user_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: User user  // 用户信息​
}

struct User {
    1: i64 id  // 用户id​
    2: string name  // 用户名称​
    3: i64 follow_count  // 关注总数​
    4: i64 follower_count  // 粉丝总数​
    5: bool is_follow  // true-已关注，false-未关注​
}