include "feed.thrift"
include "user.thrift"
include "publish.thrift"

service FeedService {
    feed.DouyinFeedResponse Feed(1: feed.DouyinFeedRequest req) (api.get="douyin/feed")
}

service UserService {
    user.DouyinUserRegisterResponse Register(1: user.DouyinUserRegisterRequest req) (api.get="douyin/user/register")
    user.DouyinUserLoginResponse Login(1: user.DouyinUserLoginRequest req) (api.post="douyin/user/login")
    user.DouyinUserResponse UserInfo(1: user.DouyinUserRequest req) (api.get="douyin/user/")
}

service PublishService {
    publish.DouyinPublishActionResponse Action(1: publish.DouyinPublishActionRequest req) (api.post="douyin/publish/action/")
    publish.DouyinPublishListResponse List(1: publish.DouyinPublishListRequest req) (api.post="douyin/publish/list/")
}