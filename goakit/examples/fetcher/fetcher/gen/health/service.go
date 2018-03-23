// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// health service
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/fetcher/design

package health

import (
	"context"
)

// Service is the health service interface.
type Service interface {
	// Health check endpoint
	Show(context.Context) (string, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "health"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"show"}
