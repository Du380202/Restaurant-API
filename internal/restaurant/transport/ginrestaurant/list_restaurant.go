package ginrestaurant

import (
	"net/http"
	"restaurant/common"
	"restaurant/component/appctx"
	"restaurant/internal/restaurant/biz"
	rmodel "restaurant/internal/restaurant/model"
	rstorage "restaurant/internal/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var filter rmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := rstorage.NewRestaurantStore(db)
		// likeStore := storage.NewSqlStore(db)
		biz := biz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c, &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}
}
