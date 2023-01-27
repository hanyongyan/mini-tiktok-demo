namespace go videoService

service VideoService {
    douyin_publish_action_response  PublishAction(1: douyin_publish_action_request req)
    douyin_feed_response  Feed(1: douyin_feed_request req)
    douyin_publish_list_response  PublishList(1: douyin_publish_list_request req)
    douyin_favorite_action_response  FavoriteAction(1: douyin_favorite_action_request req)
    douyin_favorite_list_response  FavoriteList(1: douyin_favorite_list_request req)
    douyin_comment_action_response  CommentAction(1: douyin_comment_action_request req)
    douyin_comment_list_response  CommentList(1: douyin_comment_list_request req)
}

struct douyin_publish_action_request {
    1: string token  // 用户鉴权token​
    2: byte data  // 视频数据​
    3: string title  // 视频标题​
}

struct douyin_publish_action_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
}

struct douyin_feed_request {
    1: i64 latest_time  // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间​
    2: string token  // 可选参数，登录用户设置​
}

struct douyin_feed_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: list<Video> video_list  // 视频列表​
    4: i64 next_time  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time​
}

struct Video {
    1: i64 id  // 视频唯一标识​
    2: User author  // 视频作者信息​
    3: string play_url  // 视频播放地址​
    4: string cover_url  // 视频封面地址​
    5: i64 favorite_count  // 视频的点赞总数​
    6: i64 comment_count  // 视频的评论总数​
    7: bool is_favorite  // true-已点赞，false-未点赞​
    8: string title  // 视频标题​
}

struct User {
    1: i64 id  // 用户id​
    2: string name  // 用户名称​
    3: i64 follow_count  // 关注总数​
    4: i64 follower_count  // 粉丝总数​
    5: bool is_follow  // true-已关注，false-未关注​
}
struct douyin_publish_list_request {
    1: i64 user_id  // 用户id​
    2: string token  // 用户鉴权token​
}

struct douyin_publish_list_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: list<Video> video_list  // 用户发布的视频列表​
}
struct douyin_favorite_action_request {
    1: string token  // 用户鉴权token​
    2: i64 video_id  // 视频id​
    3: i32 action_type  // 1-点赞，2-取消点赞​
}

struct douyin_favorite_action_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
}
struct douyin_favorite_list_request {
    1: i64 user_id  // 用户id​
    2: string token  // 用户鉴权token​
}

struct douyin_favorite_list_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: list<Video> video_list  // 用户点赞视频列表​
}
struct douyin_comment_action_request {
    1: string token  // 用户鉴权token​
    2: i64 video_id  // 视频id​
    3: i32 action_type  // 1-发布评论，2-删除评论​
    4: string comment_text  // 用户填写的评论内容，在action_type候使用​
    5: i64 comment_id  // 要删除的评论id，在action_type候使用​
}

struct douyin_comment_action_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: Comment comment  // 评论成功返回评论内容，不需要重新拉取整个列表​
}

struct Comment {
    1: i64 id  // 视频评论id​
    2: User user // 评论用户信息​
    3: string content  // 评论内容​
    4: string create_date  // 评论发布日期，格式 mm-dd​
}
struct douyin_comment_list_request {
    1: string token  // 用户鉴权token​
    2: i64 video_id  // 视频id​
}

struct douyin_comment_list_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败​
    2: string status_msg  // 返回状态描述​
    3: list<Comment> comment_list  // 评论列表​
}
