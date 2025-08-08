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

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		filter := model.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		paging.Fullfill()

		store := storage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := biz.NewListUserLikeRestaurant(store)

		result, err := biz.ListUser(ctx, &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, filter, paging))
	}
}
