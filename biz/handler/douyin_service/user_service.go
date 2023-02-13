// Code generated by hertz generator.

package douyin_service

import (
	"context"
	"fmt"

	core "simple_douyin/biz/model/core"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("%+v\n", req)
	resp := new(core.DouyinUserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("%+v\n", req)

	resp := new(core.DouyinUserLoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("%+v\n", req)
	resp := new(core.DouyinUserResponse)

	c.JSON(consts.StatusOK, resp)
}
