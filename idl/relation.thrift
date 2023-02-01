include "common.thrift"

namespace go extra.second

struct DouyinRelationActionRequest {
    1: string token (api.query="token")
    2: i64 to_user_id (api.query="to_user_id")
    3: i32 action_type (api.query="action_type")
}

struct DouyinRelationActionResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
}

struct DouyinRelationFollowListRequest {
    1: i64 user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct DouyinRelationFollowListResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: list<common.User> user_list (api.body="user_list")
}

struct DouyinRelationFollowerListRequest {
    1: i64 user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct DouyinRelationFollowerListResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: list<common.User> user_list (api.body="user_list")
}

struct DouyinRelationFriendListRequest {
    1: i64 user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct DouyinRelationFriendListResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: list<common.User> user_list (api.body="user_list")
}