package middleware

import (
	"restaurant/common"
	"restaurant/component/appctx"

	"github.com/gin-gonic/gin"
)

func Recover(app appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// panic(err)
				return
			}
		}()

		ctx.Next()
	}
}
