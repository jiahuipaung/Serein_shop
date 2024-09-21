package v1

// import (
// 	"net/http"
// 	"serein/pkg/utils/log"
// 	"serein/service"

// 	"serein/types"

// 	"github.com/gin-gonic/gin"
// )

// func ListCategoryHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var req types.ListCategoryReq
// 		if err := ctx.ShouldBind(&req); err != nil {
// 			log.LogrusObj.Infoln(err)
// 			ctx.JSON(http.StatusOK, "category req 参数校验失败")
// 		}

// 		l := service.ge
// 	}
// }