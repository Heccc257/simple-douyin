package database

import (
	"fmt"
	"simple_douyin/biz/model/common"
	"simple_douyin/biz/model/extra/first"
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

func FavoriteAction(uid int64, vid int64, action_type int32) error {
	var ut UserThumb
	flag := false
	if action_type == 1{
		flag = false
	}
	if result := DB.FirstOrCreate(&ut, UserThumb{Uid : uid, Vid : vid}); result.Error != nil{
		fmt.Println("点赞查询出错")
		return result.Error
	} 
	ut.IsFavorite = flag
	if result := DB.Save(&ut); result.Error != nil{
		fmt.Println("点赞更新查询出错")
		return result.Error
	}
	return nil
}

func GetFavoriteList(uid int64) ([]*common.Video, error) {
	var utLis []UserThumb
	var ret []VideoEntry
	if result := DB.Where("uid = ?", uid).Find(&utLis); result.Error != nil{
		return nil, result.Error
	} else {
		var t []int64
		for _, value := range utLis{
			t = append(t, value.Vid)
		}
		DB.Where("vid IN ?", t).Find(&ret)
		return VideoEntries2Videos(ret), nil
	}
	return nil, nil
}

func CommentActionPush(uid int64, vid int64, cid int64, comment_text string) (string, error) {
	var CE CommentEntry
	CE = CommentEntry{
		Uid : uid,
		Vid : vid,
		Cid : cid,
		CommentText : comment_text,
	}
	if result := DB.Where("uid = ? AND  vid = ? ", uid, vid).Find(&CE); result.Error != nil{
		return "", result.Error
	} else {
		// 返回日期格式 mm-dd
		return CE.Model.CreatedAt.String()[5:10], result.Error
	}
	return "", nil
}

func CommentActionDel(uid int64, vid int64) (*CommentEntry, error){
	var CE CommentEntry

	if result := DB.Where("uid = ? AND vid = ? ", uid, vid).First(&CE); result.Error != nil{
		return nil, result.Error
	} else {
		if e := DB.Delete(&CE); e.Error != nil{
			return nil, e.Error
		}
		return &CE, nil
	}

	return nil, nil
}

func GetCommentList(vid int64) ([]*first.Comment, error){
	var ret 	[]*first.Comment
	var TempC 	first.Comment
	var CELis		[]CommentEntry
	var ue 		UserEntry

	if result := DB.Order("created_at").Where("vid = ?", vid).Find(CELis); result.Error != nil{
		return nil, result.Error
	}
	if len(CELis) == 0{
		return nil, nil
	}
	for _, ce := range CELis{
		uid := ce.Uid
		if result2 := DB.First(&ue, UserEntry{UID : uid}); result2.Error != nil{
			return nil, result2.Error
		}
		TempC = first.Comment{
			ID 	: 	ce.Cid,
			User 	: 	UserEntry2User(&ue),
			Content : 	ce.CommentText,
			CreateDate	:	ce.Model.CreatedAt.String()[5:10],
		}

		ret = append(ret, &TempC)
	}

	return ret, nil
}