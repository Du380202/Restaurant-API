package middleware

import (
	"errors"
	"restaurant/common"
	"restaurant/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequire(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurentUser).(common.Requester)

		check := false

		for _, item := range allowRoles {
			if u.GetRole() == item {
				check = true
				break
			}

		}
		if !check {
			panic(common.ErrNoPermission(errors.New("invalid role user")))
		}
		c.Next()
	}
}
