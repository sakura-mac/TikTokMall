package service

import (
	"context"

	category "bytetech/course/app/frontend/hertz_gen/frontend/category"
	"bytetech/course/app/frontend/infra/rpc"
	"bytetech/course/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"title": "Category",
		"items": p.Products,
	}, nil
}
