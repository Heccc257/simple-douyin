include "common.thrift"

namespace go core


struct DouyinUserRegisterRequest {
    1: string username (api.query="username")
    2: string password (api.query="password")
}

struct DouyinUserRegisterResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: i64 user_id (api.body="user_id")
    4: string token (api.body="token")
}

struct DouyinUserLoginRequest {
    1: string username (api.query="username")
    2: string password (api.query="password")
}

struct DouyinUserLoginResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: i64 user_id (api.body="user_id")
    4: string token (api.body="token")
}

struct DouyinUserRequest {
    1: string user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct DouyinUserResponse {
    1: i32 status_code (api.body="status_code")
    2: string status_msg (api.body="status_msg")
    3: common.User user (api.body="user")
}

