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

func FavoriteAction(uid int64, VID int64, action_type int32) error {
	var ut UserThumb
	flag := false
	if action_type == 1{
		flag = false
	}
	if result := DB.Where("UID = ? and VID = ?", uid, VID).First(&ut); result.Error != nil{
		ut.UID = uid
		ut.VID = VID
		ut.IsFavorite = flag
		if result := DB.Create(&ut); result.Error != nil{
			fmt.Println("出错")
			fmt.Println(result.Error)
		}else{
			fmt.Println("点赞/撤销成功")
		}

		// fmt.Println("点赞，查询不到，成功创建")
		// DB.Create(&ut)
		return nil
	}
	result := DB.Save(&ut);
	ut.IsFavorite = flag
	return result.Error
}

func GetFavoriteList(uid int64) ([]*common.Video, error) {
	var utLis []UserThumb
	var ret []VideoEntry
	if result := DB.Where("UID = ?", uid).Find(&utLis); result.Error != nil{
		return nil, result.Error
	} else {
		var t []int64
		for index := 0; index < len(utLis); index++ {
			t = append(t, utLis[index].VID)
		}
		DB.Where("VID IN ?", t).Find(&ret)
		return VideoEntries2Videos(ret), nil
	}
	return nil, nil
}