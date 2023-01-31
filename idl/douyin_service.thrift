include "feed.thrift"
include "user.thrift"
include "publish.thrift"
include "favorite.thrift"
include "comment.thrift"
include "relation.thrift"
include "message.thrift"

service FeedService {
    feed.DouyinFeedResponse Feed(1: feed.DouyinFeedRequest req) (api.get="douyin/feed/")
}

service UserService {
    user.DouyinUserRegisterResponse Register(1: user.DouyinUserRegisterRequest req) (api.post="douyin/user/register/")
    user.DouyinUserLoginResponse Login(1: user.DouyinUserLoginRequest req) (api.post="douyin/user/login/")
    user.DouyinUserResponse UserInfo(1: user.DouyinUserRequest req) (api.get="douyin/user/")
}

service PublishService {
    publish.DouyinPublishActionResponse PublishAction(1: publish.DouyinPublishActionRequest req) (api.post="douyin/publish/action/")
    publish.DouyinPublishListResponse PublishList(1: publish.DouyinPublishListRequest req) (api.get="douyin/publish/list/")
}

service FavoriteService {
    favorite.DouyinFavoriteActionResponse FavoriteAction(1: favorite.DouyinFavoriteActionRequest req) (api.post="douyin/favorite/action/")
    favorite.DouyinFavoriteListResponse FavoriteList(1: favorite.DouyinFavoriteListRequest req) (api.get="douyin/favorite/list/")
}

service CommentService {
    comment.DouyinCommentActionResponse CommentAction(1: comment.DouyinCommentActionRequest req) (api.post="douyin/comment/action/")
    comment.DouyinCommentListResponse CommentList(1: comment.DouyinCommentListRequest req) (api.get="douyin/comment/list/")
}

service RelationService {
    relation.DouyinRelationActionResponse RelationAction(1: relation.DouyinRelationActionRequest req) (api.post="douyin/relation/action/")
    relation.DouyinRelationFollowListResponse RelationFollowList(1: relation.DouyinRelationFollowListRequest req) (api.get="douyin/relation/follow/list/")
    relation.DouyinRelationFollowerListResponse RelationFollowerList(1: relation.DouyinRelationFollowerListRequest req) (api.get="douyin/relation/follower/list/")
    relation.DouyinRelationFriendListResponse RelationFriendList(1: relation.DouyinRelationFriendListRequest req) (api.get="douyin/friend/follower/list/")
}

service MessageService {
    message.DouyinMessageChatResponse MessageChat(1: message.DouyinMessageChatResponse req) (api.get="douyin/message/chat/")
    message.DouyinMessageActionResponse MessageAction(1: message.DouyinMessageActionRequest req) (api.post="douyin/message/action/")

}