// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new products API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for products API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	Create(params *CreateParams) (*CreateOK, error)

	DeleteProduct(params *DeleteProductParams) (*DeleteProductCreated, error)

	ListProducts(params *ListProductsParams) (*ListProductsOK, error)

	UpdateProduct(params *UpdateProductParams) (*UpdateProductCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Create Create a new product
*/
func (a *Client) Create(params *CreateParams) (*CreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create",
		Method:             "POST",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProduct Returns a list of products
*/
func (a *Client) DeleteProduct(params *DeleteProductParams) (*DeleteProductCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProductParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProduct",
		Method:             "DELETE",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProductCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListProducts Returns a list of products
*/
func (a *Client) ListProducts(params *ListProductsParams) (*ListProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProductsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProducts",
		Method:             "GET",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListProductsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProductsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listProducts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateProduct Update a products details
*/
func (a *Client) UpdateProduct(params *UpdateProductParams) (*UpdateProductCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProductParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProduct",
		Method:             "PUT",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProductCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
