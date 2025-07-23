package transport

import (
	"net/http"
	"restaurant/common"
	"restaurant/component/appctx"
	"restaurant/component/hasher"
	"restaurant/internal/user/biz"
	"restaurant/internal/user/model"
	"restaurant/internal/user/storage"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data model.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := storage.NewUserStore(db)
		md5 := hasher.NewMd5Hash()
		biz := biz.NewRegisterBussiness(store, md5)

		if err := biz.Register(ctx, &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
