package v1

import (
	"net/http"
	"serein/consts"
	"serein/pkg/utils/ctl"
	"serein/pkg/utils/log"
	"serein/service"
	"serein/types"

	"github.com/gin-gonic/gin"
)

func ListProductsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductListReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Info(err)
			ctx.JSON(http.StatusOK, "ListProductsHandler 参数校验错误")
		}

		if req.PageSize == 0 {
			req.PageSize = consts.BaseProductPageSize
		}

		l := service.GetProductSrv()
		resp, err := l.ProductList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, "listproduct req 参数校验失败")
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
