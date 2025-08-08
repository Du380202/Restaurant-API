package main

import (
	"restaurant/component/appctx"
	"restaurant/internal/user/transport"
	"restaurant/middleware"

	"github.com/gin-gonic/gin"
)

func setupAdminRouters(appCtx appctx.AppContext, routerGr *gin.RouterGroup) {
	admin := routerGr.Group("/admin",
		middleware.RequireAuth(appCtx),
		middleware.RoleRequire(appCtx, "admin", "mod"),
	)

	{
		admin.GET("/profile", transport.Profile(appCtx))
	}
}
