package database

import (
	"simple_douyin/biz/model/common"

	"gorm.io/gorm"
)

type UserEntry struct {
	gorm.Model
	UID           int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
	PassWord      string
}

func User2UserEntry(u *common.User, password string) *UserEntry {
	return &UserEntry{
		UID:           u.ID,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
		PassWord:      password,
	}
}

func UserEntry2User(ue *UserEntry) *common.User {
	// 不会包含密码信息
	return &common.User{
		ID:            ue.UID,
		Name:          ue.Name,
		FollowCount:   ue.FollowCount,
		FollowerCount: ue.FollowerCount,
		IsFollow:      ue.IsFollow,
	}
}

func UserExist(user_name string) bool {
	if ue := FindUserEntryByName(user_name); ue.Model.ID == 1 {
		// 不存在范围unknown用户
		return false
	} else {
		return true
	}
}

type VideoEntry struct {
	gorm.Model
	VID           int64
	Author        UserEntry `gorm:"foreignKey:UserEntryID"`
	PlayURL       string
	CoverURL      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
	UserEntryID   int64 `gorm:"default:1;column:UserEntryID"`
}

func Video2VideoEntry(v *common.Video) *VideoEntry {
	return &VideoEntry{
		VID:           v.ID,
		PlayURL:       v.PlayURL,
		CoverURL:      v.CoverURL,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
		UserEntryID:   int64(FindUserEntryByName(v.Author.Name).Model.ID),
	}
}

func VideoEntry2Video(ve *VideoEntry) *common.Video {
	var ue UserEntry
	if res := DB.First(&ue, ve.UserEntryID); res.Error != nil {
		DB.First(&ue, 1)
	}
	return &common.Video{
		ID:            ve.VID,
		Author:        UserEntry2User(&ue),
		PlayURL:       ve.PlayURL,
		CoverURL:      ve.CoverURL,
		FavoriteCount: ve.FavoriteCount,
		CommentCount:  ve.CommentCount,
		IsFavorite:    ve.IsFavorite,
		Title:         ve.Title,
	}
}

func VideoEntries2Videos(ves []VideoEntry) (videos []*common.Video) {
	for _, ve := range ves {
		videos = append(videos, VideoEntry2Video(&ve))
	}
	return videos
}

type UserThumb struct{
	gorm.Model
	UID 		int64
	VID			int64
	IsFavorite	bool
}
