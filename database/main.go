package database

import (
	"fmt"
	"simple_douyin/biz/model/common"
	"time"
)

func FindUserEntry(u *common.User) *UserEntry {
	var ue UserEntry
	// 保证名字唯一，所以可以按名字来查找
	if result := db.Where("name = ?", u.Name).First(&ue); result.Error != nil {
		// 未找到时默认返回unknown用户
		db.First(&ue, 1)
	}

	return &ue
}

func FindVideosBefore(latest_time int64) ([]*common.Video, int64, error) {
	var ves []VideoEntry
	if result := db.Where("created_at < ?", time.Unix(latest_time, 0)).Order("created_at desc").Limit(30).Find(&ves); result.Error != nil {
		return nil, 0, result.Error
	}
	var videos []*common.Video
	for _, ve := range ves {
		fmt.Printf("%+v\n", ve)
		videos = append(videos, VideoEntry2Video(&ve))
	}
	if len(ves) == 0 { // 未找到任何视频
		return nil, 0, fmt.Errorf("%s", "No videos")
	}
	return videos, ves[0].Model.CreatedAt.Unix(), nil
}
