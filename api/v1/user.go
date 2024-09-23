package v1

import (
	// "errors"
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	// "serein/consts"
	// "serein/pkg/e"
	// "serein/pkg/utils/ctl"
	"serein/pkg/utils/log"
	"serein/service"
	"serein/types"
	// "github.com/swaggo/swag/e"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9,_%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

//	@Summary	新用户注册
//	@title		后台接口
//	@Tags		用户服务
//	@Router		/api/v1/user/register [post]
//	@param		param	body	types.UserRegisterReq	true	"用户注册请求参数"
//
// @Success 200 {object} types.UserTokenData "成功"
func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, "error:bind")
			return
		}

		// 邮箱格式校验
		if !CheckEmailFormat(req.Email) {
			err := errors.New("邮箱格式不正确")
			ctx.JSON(http.StatusOK, "error"+err.Error())
			return
		}

		l := service.GetUserSrv()
		_, err := l.UserRegister(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, "注册失败")
			return
		}
		ctx.JSON(http.StatusOK, "注册成功")
	}
}

//	@Summary	用户登录
//	@title		后台接口
//	@Tags		用户服务
//	@Router		/api/v1/user/login [post]
//	@param		param	body	types.UserLoginReq	true	"用户登录请求参数"
//
// @Success 200 {object} types.UserTokenData "成功"
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserLoginReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, "error: args")
			return
		}

		l := service.GetUserSrv()
		_, err := l.UserLogin(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, "error:login")
			return
		}
		ctx.JSON(http.StatusOK, "登录成功")
	}
}

//	@Summary	用户信息更新
//	@title		后台接口
//	@Tags		用户服务
//	@Router		/api/v1/user/update [post]
//	@param		param	body	types.UserLoginReq	true	"用户登录请求参数"
//
// @Success 200 {object} types.UserTokenData "成功"
func UserUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoUpdateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, "error: args")
			return
		}

		l := service.GetUserSrv()
		resp, err := l.UserInfoUpdate(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, "error: userinfo update")
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

// CheckEmailFormat 检查邮箱地址格式是否有效
func CheckEmailFormat(email string) bool {
	return emailRegex.MatchString(email)
}
