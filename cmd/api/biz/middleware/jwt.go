package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	utils2 "mini-tiktok-hanyongyan/pkg/utils"
	"time"
)

// JwtMiddleware jwt校验中间件
func JwtMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取 token 串
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
			if token == "" {
				c.JSON(consts.StatusOK, utils.H{"status_code": 1, "msg": "用户不存在"})
				c.Abort()
				return
			}
		}
		claims, flag := utils2.CheckToken(token)
		if !flag || time.Now().Unix() > claims.ExpiresAt {
			c.JSON(consts.StatusOK, utils.H{"status_code": 1, "msg": "登陆过期"})
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
