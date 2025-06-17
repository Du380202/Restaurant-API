package ginrestaurant

import (
	"net/http"
	"restaurant/common"
	"restaurant/component/appctx"
	"restaurant/internal/restaurant/biz"
	rstorage "restaurant/internal/restaurant/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("id")

		id, err := strconv.Atoi(param)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db := appCtx.GetMainDBConnection()

		store := rstorage.NewRestaurantStore(db)
		biz := biz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))

	}
}
