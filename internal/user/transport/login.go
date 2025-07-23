package transport

import (
	"net/http"
	"restaurant/common"
	"restaurant/component/appctx"
	"restaurant/component/hasher"
	"restaurant/component/tokenprovider/jwt"
	"restaurant/internal/user/biz"
	"restaurant/internal/user/model"
	"restaurant/internal/user/storage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginUserData model.UserLogin

		if err := ctx.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewJWTProvider("restaurant_test_project")

		store := storage.NewUserStore(db)
		md5 := hasher.NewMd5Hash()

		biz := biz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(ctx.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(account))

	}

}
