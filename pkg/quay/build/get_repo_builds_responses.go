// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/leros321/quay-exporter/pkg/models"
)

// GetRepoBuildsReader is a Reader for the GetRepoBuilds structure.
type GetRepoBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepoBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRepoBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRepoBuildsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRepoBuildsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRepoBuildsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRepoBuildsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepoBuildsOK creates a GetRepoBuildsOK with default headers values
func NewGetRepoBuildsOK() *GetRepoBuildsOK {
	return &GetRepoBuildsOK{}
}

/*GetRepoBuildsOK handles this case with default header values.

Successful invocation
*/
type GetRepoBuildsOK struct {
}

func (o *GetRepoBuildsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/][%d] getRepoBuildsOK ", 200)
}

func (o *GetRepoBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoBuildsBadRequest creates a GetRepoBuildsBadRequest with default headers values
func NewGetRepoBuildsBadRequest() *GetRepoBuildsBadRequest {
	return &GetRepoBuildsBadRequest{}
}

/*GetRepoBuildsBadRequest handles this case with default header values.

Bad Request
*/
type GetRepoBuildsBadRequest struct {
	Payload *models.APIError
}

func (o *GetRepoBuildsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/][%d] getRepoBuildsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepoBuildsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoBuildsUnauthorized creates a GetRepoBuildsUnauthorized with default headers values
func NewGetRepoBuildsUnauthorized() *GetRepoBuildsUnauthorized {
	return &GetRepoBuildsUnauthorized{}
}

/*GetRepoBuildsUnauthorized handles this case with default header values.

Session required
*/
type GetRepoBuildsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetRepoBuildsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/][%d] getRepoBuildsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepoBuildsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoBuildsForbidden creates a GetRepoBuildsForbidden with default headers values
func NewGetRepoBuildsForbidden() *GetRepoBuildsForbidden {
	return &GetRepoBuildsForbidden{}
}

/*GetRepoBuildsForbidden handles this case with default header values.

Unauthorized access
*/
type GetRepoBuildsForbidden struct {
	Payload *models.APIError
}

func (o *GetRepoBuildsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/][%d] getRepoBuildsForbidden  %+v", 403, o.Payload)
}

func (o *GetRepoBuildsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoBuildsNotFound creates a GetRepoBuildsNotFound with default headers values
func NewGetRepoBuildsNotFound() *GetRepoBuildsNotFound {
	return &GetRepoBuildsNotFound{}
}

/*GetRepoBuildsNotFound handles this case with default header values.

Not found
*/
type GetRepoBuildsNotFound struct {
	Payload *models.APIError
}

func (o *GetRepoBuildsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/][%d] getRepoBuildsNotFound  %+v", 404, o.Payload)
}

func (o *GetRepoBuildsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
