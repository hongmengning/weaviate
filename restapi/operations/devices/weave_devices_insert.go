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

// WeaveDevicesInsertHandlerFunc turns a function with the right signature into a weave devices insert handler
type WeaveDevicesInsertHandlerFunc func(WeaveDevicesInsertParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveDevicesInsertHandlerFunc) Handle(params WeaveDevicesInsertParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveDevicesInsertHandler interface for that can handle valid weave devices insert params
type WeaveDevicesInsertHandler interface {
	Handle(WeaveDevicesInsertParams, interface{}) middleware.Responder
}

// NewWeaveDevicesInsert creates a new http.Handler for the weave devices insert operation
func NewWeaveDevicesInsert(ctx *middleware.Context, handler WeaveDevicesInsertHandler) *WeaveDevicesInsert {
	return &WeaveDevicesInsert{Context: ctx, Handler: handler}
}

/*WeaveDevicesInsert swagger:route POST /devices devices weaveDevicesInsert

Registers a new device. This method may be used only by aggregator devices or adapters.

*/
type WeaveDevicesInsert struct {
	Context *middleware.Context
	Handler WeaveDevicesInsertHandler
}

func (o *WeaveDevicesInsert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveDevicesInsertParams()

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
