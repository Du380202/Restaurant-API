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

func CreateRestaurant(appContext appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appContext.GetMainDBConnection()

		requester := c.MustGet(common.CurentUser).(common.Requester)

		var data rmodel.Restaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))
			return
		}

		data.UserId = requester.GetUserId()

		store := rstorage.NewRestaurantStore(db)
		biz := biz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data))

	}
}
