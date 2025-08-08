package main

import (
	"restaurant/component/appctx"
	"restaurant/internal/restaurant/transport/ginrestaurant"
	ulrtransport "restaurant/internal/restaurantlike/transport"
	"restaurant/internal/user/transport"
	"restaurant/middleware"

	"github.com/gin-gonic/gin"
)

func setupRouters(appCtx appctx.AppContext, routerGr *gin.RouterGroup) {
	routerGr.POST("/register", transport.Register(appCtx))

	routerGr.POST("/login", transport.Login(appCtx))

	routerGr.GET("/profile", middleware.RequireAuth(appCtx), transport.Profile(appCtx))

	restaurant := routerGr.Group("/restaurant", middleware.RequireAuth(appCtx))

	restaurant.GET("/", ginrestaurant.ListRestaurant(appCtx))

	restaurant.POST("/", ginrestaurant.CreateRestaurant(appCtx))

	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	restaurant.POST("/:id/like", ulrtransport.UserLikeRestaurant(appCtx))

	restaurant.DELETE("/:id/unlike", ulrtransport.UserUnLikeRestaurant(appCtx))

	restaurant.GET("/:id/like-users", ulrtransport.ListUser(appCtx))

}
