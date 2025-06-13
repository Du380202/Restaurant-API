package ginrestaurant

import (
	"net/http"
	"restaurant/internal/restaurant/biz"
	rmodel "restaurant/internal/restaurant/model"
	rstorage "restaurant/internal/restaurant/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data rmodel.Restaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := rstorage.NewRestaurantStore(db)
		biz := biz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"data": data,
		})

	}
}
