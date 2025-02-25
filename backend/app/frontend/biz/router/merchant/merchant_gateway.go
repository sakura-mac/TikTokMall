// Code generated by hertz generator. DO NOT EDIT.

package merchant

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	merchant "github.com/tiktokmall/backend/app/frontend/biz/handler/merchant"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_merchant := root.Group("/merchant", _merchantMw()...)
		_merchant.GET("/auth", append(_merchantauthMw(), merchant.MerchantAuth)...)
		_merchant.GET("/ping", append(_merchantpingMw(), merchant.MerchantPing)...)
		_merchant.POST("/register", append(_merchantregisterMw(), merchant.MerchantRegister)...)
		{
			_product := _merchant.Group("/product", _productMw()...)
			_product.GET("/:id", append(_merchantgetproductdetailMw(), merchant.MerchantGetProductDetail)...)
			_product.POST("/add", append(_merchantaddproductMw(), merchant.MerchantAddProduct)...)
			_product.POST("/delete", append(_merchantdeleteproductMw(), merchant.MerchantDeleteProduct)...)
			_product.POST("/list", append(_merchantgetproductlistMw(), merchant.MerchantGetProductList)...)
			_product.POST("/update", append(_merchantupdateproductMw(), merchant.MerchantUpdateProduct)...)
		}
	}
}
