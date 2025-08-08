package transport

import (
	"net/http"
	"restaurant/common"
	"restaurant/component/appctx"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := ctx.MustGet(common.CurentUser).(common.Requester)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
