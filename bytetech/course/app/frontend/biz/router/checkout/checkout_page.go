// Code generated by hertz generator. DO NOT EDIT.

package checkout

import (
	checkout "bytetech/course/app/frontend/biz/handler/checkout"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/checkout", append(_checkout0Mw(), checkout.Checkout)...)
	_checkout := root.Group("/checkout", _checkoutMw()...)
	_checkout.GET("/result", append(_checkoutresultMw(), checkout.CheckoutResult)...)
	_checkout.POST("/waiting", append(_checkoutwaitingMw(), checkout.CheckoutWaiting)...)
}
