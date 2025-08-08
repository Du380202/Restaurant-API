package ulrtransport

import (
	"net/http"
	"restaurant/common"
	"restaurant/component/appctx"
	"restaurant/internal/restaurantlike/biz"
	"restaurant/internal/restaurantlike/model"
	"restaurant/internal/restaurantlike/storage"

	"github.com/gin-gonic/gin"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurentUser).(common.Requester)

		data := model.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := storage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := biz.NewUserLikeRestaurantBiz(store)

		if err := biz.LikeRestaurant(ctx, &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
