package main

import (
	"context"
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

// StringService provides operations on strings.
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

// stringService is a concrete implementation of StringService
type stringService struct {}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// Uppercase takes a string and uppercase's the whole thing
func(stringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// Count returns the length of a string
func(stringService) Count(_ context.Context, s string) int {
	return len(s)
}
/**
 * REQUEST AND RESPONSE
 * In Go kit, the primary messaging pattern is RPC.
 * So, every method in our interface will be modeled as a remote procedure call.
 * For each method, we define REQUEST and RESPONSE structs, capturing all of the input and output parameters.
 */
 type uppercaseRequest struct {
 	S string `json:"s"`
 }

 type uppercaseResponse struct {
 	V   string `json:"v"`
 	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
 }

 type countRequest struct {
 	S string `json:"s"`
 }

 type countResponse struct {
 	V int `json:"V"`
 }
/**
 * ENDPOINTS
 * Go Kit provides much of it's functionality through an abstraction called an endpoint.
 * An endpoint represents a single RPC.
 * That is, a single method in our service interface.
 * We'll write simple adapters to convert each of our service's methods into an endpoint.
 * Each adapter takes a StringService, and returns an endpoint that corresponds to one of the methods.
 */
 // makeUppsercaseEndpoint will create an rpc endpoint for Uppercase method for the StringService interface
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}

		return uppercaseResponse{v, ""}, nil
	}
}


// makeCountEndpoint will create an rpc endpoint for Count method for the StringService interface
func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(ctx, req.S)

		return countResponse{v}, nil
	}
}