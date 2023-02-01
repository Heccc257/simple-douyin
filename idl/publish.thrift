include "common.thrift"

namespace go core

struct DouyinPublishActionRequest {
    1: string token (api.form="token")
    2: list<byte> data (api.form="data")
    3: string title (api.form="title")
}

struct DouyinPublishActionResponse {
    1: i64 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
}

struct DouyinPublishListRequest {
    1: string token (api.query="token")
    2: i64 user_id (api.query="user_id")
}

struct DouyinPublishListResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: list<common.Video> video_list (api.body="video_list")
}