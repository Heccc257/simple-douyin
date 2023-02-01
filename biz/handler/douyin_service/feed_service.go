// Code generated by hertz generator.

package douyin_service

import (
	"context"
	"fmt"

	core "simple_douyin/biz/model/core"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Feed .
// @router douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("%+v\n", req)
	resp := new(core.DouyinFeedResponse)

	c.JSON(consts.StatusOK, resp)
}
