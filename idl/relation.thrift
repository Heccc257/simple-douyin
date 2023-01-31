include "common.thrift"

namespace go extra.second

struct DouyinRelationActionRequest {
    1: string Token (api.query="token")
    2: i64 ToUserID (api.query="to_user_id")
    3: i32 ActionType (api.query="action_type")
}

struct DouyinRelationActionResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
}

struct DouyinRelationFollowListRequest {
    1: i64 UserID (api.query="user_id")
    2: string Token (api.query="token")
}

struct DouyinRelationFollowListResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: list<common.User> UserList (api.body="user_list")
}

struct DouyinRelationFollowerListRequest {
    1: i64 UserID (api.query="user_id")
    2: string Token (api.query="token")
}

struct DouyinRelationFollowerListResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: list<common.User> UserList (api.body="user_list")
}

struct DouyinRelationFriendListRequest {
    1: i64 UserID (api.query="user_id")
    2: string Token (api.query="token")
}

struct DouyinRelationFriendListResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: list<common.User> UserList (api.body="user_list")
}