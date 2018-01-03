// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// fetcher go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/fetcher/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/http"
)

// MountFetchHandler configures the mux to serve the "fetcher" service "fetch"
// endpoint.
func MountFetchHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/fetch/{*url}", f)
}
