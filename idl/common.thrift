struct User {
    1: i64 ID (api.body="id")
    2: string Name (api.body="name")
    3: i64 FollowCount (api.body="follow_count")
    4: i64 FollowerCount (api.body="follower_count")
    5: bool Is_follow (api.body="is_follow") 
}

struct Video {
    1: i64 ID (api.body="id")
    2: User Author (api.body="author")
    3: string PlayUrl (api.body="play_url")
    4: string CoverUrl (api.body="cover_url ")
    5: i64 FavoriteCount (api.body="favorite_count")
    6: i64 CommentCount (api.body="comment_count")
    7: bool IsFavorite (api.body="is_favorite")
    8: string Title (api.body="title")
}
