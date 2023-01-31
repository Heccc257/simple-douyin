include "common.thrift"

namespace go extra.first

struct Comment {
    1: i64 ID (api.body="id")
    2: common.User User (api.body="user")
    3: string Content (api.body="content")
    4: string CreateDate (api.body="create_date")
}

struct DouyinCommentActionRequest {
    1: string Token (api.query="token")
    2: i64 VideoID (api.query="video_id")
    3: i32 ActionType (api.query="action_type")
    4: string CommentText (api.query="comment_text")
    5: i64 CommentID (api.query="comment_id")
}

struct DouyinCommentActionResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: Comment Comment (api.body="comment")
}

struct DouyinCommentListRequest {
    1: string Token (api.query="token") // 用户鉴权token
    2: i64 VideoID (api.query="video_id") // 视频id
}

struct DouyinCommentListResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: list<Comment> CommentList (api.body="comment_list")
}
