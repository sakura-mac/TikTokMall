package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/redis"
	"github.com/tiktokmall/backend/app/merchant/biz/model"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *merchant.UpdateProductReq) (resp *merchant.UpdateProductResp, err error) {
	// Finish your business logic.
	// 1. 检查参数
	if err = checkUpdateProductReq(req); err != nil {
		return nil, err
	}
	// 2. 检查商户, TODO: 去掉也行
	_, err = model.NewMerchantQuery(s.ctx, mysql.DB).GetById(int(req.GetMerchantId()))
	if err != nil {
		return nil, fmt.Errorf("merchant [%v] not found, err:%w", req.GetMerchantId(), err)
	}
	// 3. 检查类别, 目前只有 req.Categories 中的 id 有用。即使类别修改了，也会忽略
	// TODO: 增加类别的修改
	reqCategories := req.GetProduct().GetCategory()
	var categories []model.Category
	if reqCategories == nil {
		categories = nil
	} else {
		reqCsIds := []int64{}
		for _, c := range reqCategories {
			reqCsIds = append(reqCsIds, c.GetId())
		}
		categories, err = model.NewCategoryQuery(s.ctx, mysql.DB).GetManyById(reqCsIds)
		if err != nil {
			return nil, fmt.Errorf("cannot find all categories [%v], err:%w", reqCsIds, err)
		}
	}
	// 4. 插入商品 与 类别关系
	newProduct := &model.Product{
		Base:        model.Base{ID: int(req.GetProduct().GetId())},
		Name:        req.GetProduct().GetName(),
		Description: req.GetProduct().GetDescription(),
		Picture:     req.GetProduct().GetImgUrl(),
		Price:       req.GetProduct().GetPrice(),
		SliderImgs:  strings.Join(req.GetProduct().GetSliderImgs(), "\n"),
		Stock:       int32(req.GetProduct().GetStock()),
		MerchantID:  int(req.GetMerchantId()),
		Categories:  categories,
	}
	err = model.NewProductQuery(s.ctx, mysql.DB).UpdateOne(*newProduct)

	// TODO: 删除 product_list 的缓存
	model.ClearProductListCachedKey(s.ctx, redis.RedisClient)
	return
}

func checkUpdateProductReq(req *merchant.UpdateProductReq) error {
	if req.GetMerchantId() <= 0 {
		return kerrors.NewBizStatusError(2004001, "merchant id must be > 0")
	}
	reqProduct := req.GetProduct()
	if reqProduct == nil {
		return kerrors.NewBizStatusError(2004001, "product is required")
	}
	if reqProduct.GetId() <= 0 {
		return kerrors.NewBizStatusError(2004001, "product id must be > 0")
	}
	if reqProduct.GetName() == "" {
		return kerrors.NewBizStatusError(2004001, "product name is required")
	}
	if reqProduct.GetPrice() < 0 {
		return kerrors.NewBizStatusError(2004001, "product price must be >= 0")
	}
	if reqProduct.GetStock() < 0 {
		return kerrors.NewBizStatusError(2004001, "product stock must be >= 0")
	}
	// reqCategory := reqProduct.GetCategory()
	// if reqCategory == nil {
	// 	return kerrors.NewBizStatusError(2004001, "product category is required")
	// }
	return nil
}
