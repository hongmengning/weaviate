/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package devices




import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveDevicesAddNicknameHandlerFunc turns a function with the right signature into a weave devices add nickname handler
type WeaveDevicesAddNicknameHandlerFunc func(WeaveDevicesAddNicknameParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveDevicesAddNicknameHandlerFunc) Handle(params WeaveDevicesAddNicknameParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveDevicesAddNicknameHandler interface for that can handle valid weave devices add nickname params
type WeaveDevicesAddNicknameHandler interface {
	Handle(WeaveDevicesAddNicknameParams, interface{}) middleware.Responder
}

// NewWeaveDevicesAddNickname creates a new http.Handler for the weave devices add nickname operation
func NewWeaveDevicesAddNickname(ctx *middleware.Context, handler WeaveDevicesAddNicknameHandler) *WeaveDevicesAddNickname {
	return &WeaveDevicesAddNickname{Context: ctx, Handler: handler}
}

/*WeaveDevicesAddNickname swagger:route POST /devices/{deviceId}/addNickname devices weaveDevicesAddNickname

Adds a nickname to the device.

*/
type WeaveDevicesAddNickname struct {
	Context *middleware.Context
	Handler WeaveDevicesAddNicknameHandler
}

func (o *WeaveDevicesAddNickname) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveDevicesAddNicknameParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
