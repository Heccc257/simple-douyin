// Code generated by hertz generator.

package douyin_service

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"sync"

	"simple_douyin/biz/model/common"
	core "simple_douyin/biz/model/core"
	"simple_douyin/database"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishAction .
// @router douyin/publish/action/ [POST]
type DouyinPublishActionRequest struct {
	Token string                `thrift:"token,1" form:"token" json:"token"`
	Data  *multipart.FileHeader `thrift:"data,2" form:"data" json:"data"`
	Title string                `thrift:"title,3" form:"title" json:"title"`
}

// 总的VideoID，从1开始（与数据库表项的ID要区分）
var (
	globalVideoID int64 = 2
	videoIDLock   sync.Mutex
)

func assignVideoID() (userID int64) {
	userIDLock.Lock()
	globalUserID++
	userID = globalUserID
	userIDLock.Unlock()
	return
}

func PublishAction(ctx context.Context, c *app.RequestContext) {
	fmt.Println("publish Action")
	var err error
	// var req core.DouyinPublishActionRequest

	// Request的Data数据类型必须是*multipart.FileHeader
	// 所以自定义一个request类型
	var req DouyinPublishActionRequest
	err = c.BindAndValidate(&req)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	token := req.Token
	user, exist := userLoginInfo[token]
	// token未通过校验
	if !exist {
		c.JSON(consts.StatusOK, &core.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "unqualified",
		})
		return
	}

	file, err := req.Data.Open()
	if err != nil {
		c.JSON(consts.StatusOK, &core.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}
	defer file.Close()
	p := make([]byte, req.Data.Size)

	if _, err := file.Read(p); err != nil {
		log.Println(err.Error())
		c.JSON(consts.StatusOK, &core.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 保存到本地
	fileDir := "./public/" + user.Name + "/"
	fileName := req.Title
	os.MkdirAll(fileDir, 0777)
	os.Chmod(fileDir, 0777)

	saveFile, err := os.Create(fileDir + fileName)

	if err != nil {
		log.Println(err.Error())
	}

	defer saveFile.Close()
	if _, err := saveFile.Write(p); err != nil {
		c.JSON(consts.StatusOK, &core.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}
	log.Println("create file: ", fileDir+fileName)

	// 添加到数据库
	// 如果重复上传同一名称的视频可能会出错
	database.UpdateVideo(&common.Video{
		ID:       assignVideoID(),
		Author:   &user,
		PlayURL:  fileDir + fileName,
		CoverURL: "?",
		Title:    req.Title,
	})
	// var ves []database.VideoEntry
	// database.DB.Find(&ves)
	// fmt.Println(ves)

	// 正常
	c.JSON(consts.StatusOK, &core.DouyinPublishActionResponse{
		StatusCode: 0,
	})
}

// PublishList .
// @router douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	token := req.Token
	_, exist := userLoginInfo[token]
	if !exist {
		c.JSON(consts.StatusOK, &core.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "unqualified",
		})
		return
	}
	videos := database.FindVideosByUserID(req.UserID)
	c.JSON(consts.StatusOK, &core.DouyinPublishListResponse{
		StatusCode: 0,
		VideoList:  videos,
	})
}
