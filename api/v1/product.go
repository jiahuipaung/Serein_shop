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

//	@Summary	根据目录信息获取整个list的所有商品信息
//	@title		后台接口
//	@Tags		商品服务
//	@Router		/api/v1/product/list [get]
//	@param		param	body	types.ProductListReq	true	"商品列表请求请求参数"
//
// @Success 200 {object} types.DataListResp "成功"
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

//	@Summary	获取商品详情
//	@title		后台接口
//	@Tags		商品服务
//	@Router		/api/v1/product/show [get]
//	@param		param	body	types.ProductListReq	true	"商品列表请求请求参数"
//
// @Success 200 {object} types.DataListResp "成功"
func ShowProductsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductShowReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, "err: args")
			return
		}

		l := service.GetProductSrv()
		resp, err := l.ProductShow(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, "err: show product")
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
