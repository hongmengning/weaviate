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
 package adapters




import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveAdaptersGetHandlerFunc turns a function with the right signature into a weave adapters get handler
type WeaveAdaptersGetHandlerFunc func(WeaveAdaptersGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveAdaptersGetHandlerFunc) Handle(params WeaveAdaptersGetParams) middleware.Responder {
	return fn(params)
}

// WeaveAdaptersGetHandler interface for that can handle valid weave adapters get params
type WeaveAdaptersGetHandler interface {
	Handle(WeaveAdaptersGetParams) middleware.Responder
}

// NewWeaveAdaptersGet creates a new http.Handler for the weave adapters get operation
func NewWeaveAdaptersGet(ctx *middleware.Context, handler WeaveAdaptersGetHandler) *WeaveAdaptersGet {
	return &WeaveAdaptersGet{Context: ctx, Handler: handler}
}

/*WeaveAdaptersGet swagger:route GET /adapters/{adapterId} adapters weaveAdaptersGet

Get an adapter.

*/
type WeaveAdaptersGet struct {
	Context *middleware.Context
	Handler WeaveAdaptersGetHandler
}

func (o *WeaveAdaptersGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveAdaptersGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
