include "common.thrift"

namespace go extra.first

struct DouyinFavoriteActionRequest {
    1: string Token (api.query="token")
    2: i64 VideoID (api.query="video_id")
    3: i32 ActionType (api.query="action_type")
}

struct DouyinFavoriteActionResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
}

struct DouyinFavoriteListRequest {
    1: i64 UserID (api.query="user_id")
    2: string Token (api.query="token")
}

struct DouyinFavoriteListResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: list<common.Video> VideoList (api.body="video_list")
}