package v1

import (
	"net/http"
	"serein/pkg/utils/ctl"
	"serein/pkg/utils/log"
	"serein/service"

	"serein/types"

	"github.com/gin-gonic/gin"
)

func ListCategoryHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListCategoryReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, "category req 参数校验失败")
		}

		l := service.GetCategorySrv()
		resp, err := l.CategoryList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, "cart创建err")
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}