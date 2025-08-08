package ginrestaurant

import (
	"net/http"
	"restaurant/common"
	"restaurant/component/appctx"
	"restaurant/internal/restaurant/biz"
	rstorage "restaurant/internal/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// param := c.Param("id")

		requester := c.MustGet(common.CurentUser).(common.Requester)

		// id, err := strconv.Atoi(param)

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db := appCtx.GetMainDBConnection()

		store := rstorage.NewRestaurantStore(db)
		biz := biz.NewDeleteRestaurantBiz(store, requester)

		if err := biz.DeleteRestaurant(c, int(uid.GetLocalID())); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorNoPermission(err))
			return
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))

	}
}
