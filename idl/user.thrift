include "common.thrift"

namespace go core


struct DouyinUserRegisterRequest {
    1: string UserName (api.query="username")
    2: string PassWord (api.query="password")
}

struct DouyinUserRegisterResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: i64 UserID (api.body="user_id")
    4: string Token (api.body="token")
}

struct DouyinUserLoginRequest {
    1: string UserName (api.query="username")
    2: string PassWord (api.query="password")
}

struct DouyinUserLoginResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: i64 UserID (api.body="user_id")
    4: string Token (api.body="token")
}

struct DouyinUserRequest {
    1: string UserID (api.query="user_id")
    2: string Token (api.query="token")
}

struct DouyinUserResponse {
    1: i32 StatusCode (api.body="status_code")
    2: string StatusMsg (api.body="status_msg")
    3: common.User User (api.body="user")
}

