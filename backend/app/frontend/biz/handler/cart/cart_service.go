package cart

import (
	"context"

	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	cart "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/cart"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"

	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		// c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	_, err = service.NewAddCartItemService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		// c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}
	c.Redirect(consts.StatusFound, []byte("/cart"))
}

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, map[string]any{"warning": err}))
		// c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, map[string]any{"warning": err}))
		return
	}
	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, map[string]any{"warning": err}))
		// c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, map[string]any{"warning": err}))
		return
	}
	c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, resp))
	// c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}
