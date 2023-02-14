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
}

func User2UserEntry(u *common.User) *UserEntry {
	return &UserEntry{}
}

func UserEntry2User(ue *UserEntry) *common.User {
	return &common.User{
		ID:            ue.UID,
		Name:          ue.Name,
		FollowCount:   ue.FollowCount,
		FollowerCount: ue.FollowerCount,
		IsFollow:      ue.IsFollow,
	}
}

type VideoEntry struct {
	gorm.Model
	VID           int64
	Author        UserEntry `gorm:"ForeignKey:UserEntryID"`
	PlayURL       string
	CoverURL      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
	UserEntryID   int64 `gorm:"default:1"`
}

func Video2VideoEntry(v *common.Video) *VideoEntry {
	return &VideoEntry{}
}

func VideoEntry2Video(ve *VideoEntry) *common.Video {
	return &common.Video{
		ID:            ve.VID,
		Author:        UserEntry2User(&ve.Author),
		PlayURL:       ve.PlayURL,
		CoverURL:      ve.CoverURL,
		FavoriteCount: ve.FavoriteCount,
		CommentCount:  ve.CommentCount,
		IsFavorite:    ve.IsFavorite,
		Title:         ve.Title,
	}
}
