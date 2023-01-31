include "common.thrift"

namespace go core

struct DouyinFeedRequest {
    1: i64 LatestTime (api.query="latest_time")
    2: string Token (api.query="token")
}

struct DouyinFeedResponse {
    1: i64 StatusCode (api.body="status_code") // 状态码，0-成功，其他值-失败
    2: string StatusMsg (api.body="status_msg") 
    3: list<common.Video> VideoList (api.body="video_list")
    4: i64 NextTime (api.body="next_time")
}
