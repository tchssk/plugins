// Code generated by goa v3.0.0, DO NOT EDIT.
//
// fetcher go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/fetcher/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/fetcher

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	goahttp "goa.design/goa/v3/http"
	"goa.design/plugins/v3/goakit/examples/fetcher/fetcher/gen/http/fetcher/server"
)

// EncodeFetchResponse returns a go-kit EncodeResponseFunc suitable for
// encoding fetcher fetch responses.
func EncodeFetchResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeFetchResponse(encoder)
}

// DecodeFetchRequest returns a go-kit DecodeRequestFunc suitable for decoding
// fetcher fetch requests.
func DecodeFetchRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeFetchRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeFetchError returns a go-kit EncodeResponseFunc suitable for encoding
// errors returned by the fetcher fetch endpoint.
func EncodeFetchError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.ErrorEncoder {
	enc := server.EncodeFetchError(encoder)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}
