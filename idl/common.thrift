struct User {
    1: i64 id (api.body="id")
    2: string name (api.body="name")
    3: i64 follow_count (api.body="follow_count")
    4: i64 follower_count (api.body="follower_count")
    5: bool is_follow (api.body="is_follow") 
}

struct Video {
    1: i64 id (api.body="id")
    2: User author (api.body="author")
    3: string play_url (api.body="play_url")
    4: string cover_url (api.body="cover_url")
    5: i64 favorite_count (api.body="favorite_count")
    6: i64 comment_count (api.body="comment_count")
    7: bool is_favorite (api.body="is_favorite")
    8: string title (api.body="title")
}
