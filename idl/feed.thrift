include "common.thrift"

namespace go core

struct DouyinFeedRequest {
    1: i64 latest_time (api.query="latest_time")
    2: string token (api.query="token")
}

struct DouyinFeedResponse {
    1: i64 status_code (api.body="status_code") // 状态码，0-成功，其他值-失败
    2: string status_msg (api.body="status_msg") 
    3: list<common.Video> video_list (api.body="video_list")
    4: i64 next_time (api.body="next_time")
}
