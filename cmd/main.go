package main

import (
	"log"
	"net/http"
	"restaurant/component/appctx"
	"restaurant/config"
	"restaurant/internal/restaurant/transport/ginrestaurant"
	"restaurant/internal/user/transport"
	"restaurant/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectPostgres()

	db := config.GetDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic("Không thể lấy sql.DB từ GORM: " + err.Error())
	}
	defer sqlDB.Close()

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	routerGr := r.Group("/api/v1")

	routerGr.POST("/register", transport.Register(appCtx))

	routerGr.POST("/login", transport.Login(appCtx))

	routerGr.GET("/profile", middleware.RequireAuth(appCtx), transport.Profile(appCtx))

	routerGr.GET("/restaurant", ginrestaurant.ListRestaurant(appCtx))

	routerGr.POST("restaurant", ginrestaurant.CreateRestaurant(appCtx))

	port := config.AppConfig.Server.Port
	server := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Fatal(server.ListenAndServe())
}
