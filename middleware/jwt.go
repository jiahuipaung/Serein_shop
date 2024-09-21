package middleware

import (
	"serein/consts"
	"serein/pkg/e"

	"serein/pkg/utils/ctl"
	util "serein/pkg/utils/jwt"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware token验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e.SUCCESS
		accessToken := ctx.GetHeader("accessToken")
		refreshToken := ctx.GetHeader("refreshToken")
		if accessToken == "" { //token为空
			code = e.InvalidParams
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token不能为空",
			})
			ctx.Abort()
			return
		}

		newAccessToken, newRefreshToken, err := util.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		}
		if code != e.SUCCESS {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "鉴权失败",
				"error":  err.Error(),
			})
			ctx.Abort()
			return
		}
		
		claims, err := util.ParseToken(newAccessToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   err.Error(),
			})
			ctx.Abort()
			return
		}

		SetToken(ctx, newAccessToken, newRefreshToken)
		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(), &ctl.UserInfo{Id: claims.ID}))
		ctl.InitUserInfo(ctx.Request.Context())
		ctx.Next()
	}
}

func SetToken(ctx *gin.Context, accessToken, refreshToken string) {
	secure := IsHttps(ctx)
	ctx.Header(consts.AccessTokenHeader, accessToken)
	ctx.Header(consts.RefreshTokenHeader, refreshToken)
	ctx.SetCookie(consts.AccessTokenHeader, accessToken, consts.MaxAge, "/", "", secure, true)
	ctx.SetCookie(consts.RefreshTokenHeader, refreshToken, consts.MaxAge, "/", "", secure, true)
}

func IsHttps(ctx *gin.Context) bool {
	if ctx.GetHeader(consts.HeaderForwardedProto) == "https" || ctx.Request.TLS != nil {
		return true
	}

	return false
}
