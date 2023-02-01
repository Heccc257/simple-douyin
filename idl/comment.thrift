include "common.thrift"

namespace go extra.first

struct Comment {
    1: i64 id (api.body="id")
    2: common.User user (api.body="user")
    3: string content (api.body="content")
    4: string create_date (api.body="create_date")
}

struct DouyinCommentActionRequest {
    1: string token (api.query="token")
    2: i64 video_id (api.query="video_id")
    3: i32 action_type (api.query="action_type")
    4: string comment_text (api.query="comment_text")
    5: i64 comment_id (api.query="comment_id")
}

struct DouyinCommentActionResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: Comment comment (api.body="comment")
}

struct DouyinCommentListRequest {
    1: string token (api.query="token") // 用户鉴权token
    2: i64 video_id (api.query="video_id") // 视频id
}

struct DouyinCommentListResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: list<Comment> comment_list (api.body="comment_list")
}
