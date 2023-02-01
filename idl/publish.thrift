include "common.thrift"

namespace go core

struct DouyinPublishActionRequest {
    1: string Token (api.form="token")
    2: list<byte> Data (api.form="data")
    3: string Title (api.form="title")
}

struct DouyinPublishActionResponse {
    1: i64 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
}

struct DouyinPublishListRequest {
    1: string Token (api.query="token")
    2: i64 UserID (api.query="user_id")
}

struct DouyinPublishListResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: list<common.Video> VideoList (api.body="video_list")
}