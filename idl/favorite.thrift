include "common.thrift"

namespace go extra.first

struct DouyinFavoriteActionRequest {
    1: string token (api.query="token")
    2: i64 video_id (api.query="video_id")
    3: i32 action_type (api.query="action_type")
}

struct DouyinFavoriteActionResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
}

struct DouyinFavoriteListRequest {
    1: i64 user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct DouyinFavoriteListResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: list<common.Video> video_list (api.body="video_list")
}