package ulrtransport

import (
	"net/http"
	"restaurant/common"
	"restaurant/component/appctx"
	rstorage "restaurant/internal/restaurant/storage"
	"restaurant/internal/restaurantlike/biz"
	"restaurant/internal/restaurantlike/storage"

	"github.com/gin-gonic/gin"
)

func UserUnLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurentUser).(common.Requester)

		store := storage.NewSqlStore(appCtx.GetMainDBConnection())
		deStore := rstorage.NewRestaurantStore(appCtx.GetMainDBConnection())
		biz := biz.NewUserUnLikeRestaurantBiz(store, deStore)

		if err := biz.UnlikeRestaurant(ctx, requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
