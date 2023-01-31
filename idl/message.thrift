include "common.thrift"

namespace go extra.second

struct Message {
    1: string Token (api.body="token")
    2: i64 ToUserID (api.body="to_user_id")
    3: i64 FromUserID (api.body="from_user_id")
    4: string Content (api.body="content")
    5: string CreateTime (api.body="create_time")
}

struct DouyinMessageChatRequest {
    1: string Token (api.query="token")
    2: i64 ToUserID (api.query="to_user_id")
}

struct DouyinMessageChatResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: list<Message> MessageList (api.body="message_list")
}

struct DouyinMessageActionRequest {
    1: string Token (api.query="token")
    2: i64 ToUserID (api.query="to_user_id")
    3: i32 ActionType (api.query="action_type")
    4: string Content (api.query="content")
}

struct DouyinMessageActionResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
}
