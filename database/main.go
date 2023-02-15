package database

import (
	"fmt"
	"simple_douyin/biz/model/common"
	"time"
)

func FindUserEntryByName(user_name string) *UserEntry {
	var ue UserEntry
	// 保证名字唯一，所以可以按名字来查找
	if result := DB.Where("name = ?", user_name).First(&ue); result.Error != nil {
		// 未找到时默认返回unknown用户
		DB.First(&ue, 1)
	}

	return &ue
}

func FindUserEntryByUID(user_id int64) *UserEntry {
	var ue UserEntry
	if result := DB.Where("UID = ?", user_id).First(&ue); result.Error != nil {
		// 未找到时默认返回unknown用户
		DB.First(&ue, 1)
	}

	return &ue
}

func UpdateUser(u *common.User, password string) error {
	ue := User2UserEntry(u, password)
	if res := DB.Create(ue); res.Error != nil {
		return res.Error
	}
	return nil
}

func FindVideosBefore(latest_time int64) ([]*common.Video, int64, error) {
	var ves []VideoEntry
	if result := DB.Where("created_at < ?", time.Unix(latest_time, 0)).Order("created_at desc").Limit(30).Find(&ves); result.Error != nil {
		return nil, 0, result.Error
	}
	var videos []*common.Video
	for _, ve := range ves {
		videos = append(videos, VideoEntry2Video(&ve))
	}
	if len(ves) == 0 { // 未找到任何视频
		return nil, 0, fmt.Errorf("%s", "No videos")
	}
	return videos, ves[0].Model.CreatedAt.Unix(), nil
}

func UpdateVideo(v *common.Video) error {
	ve := Video2VideoEntry(v)
	fmt.Printf("%+v\n", ve)
	if res := DB.Create(ve); res.Error != nil {
		return res.Error
	}
	return nil
}

func FindVideosByUserID(user_id int64) []*common.Video {
	// 需要先根据UID求出数据库的ID
	ue := FindUserEntryByUID(user_id)
	fmt.Println("userentry", ue)
	if ue.Name == unknownUser.Name {
		return nil
	}
	var ves []VideoEntry
	DB.Where("UserEntryID = ?", ue.Model.ID).Find(&ves)
	return VideoEntries2Videos(ves)
}
