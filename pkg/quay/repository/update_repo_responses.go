// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/leros321/quay-exporter/pkg/models"
)

// UpdateRepoReader is a Reader for the UpdateRepo structure.
type UpdateRepoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateRepoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateRepoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateRepoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateRepoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateRepoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRepoOK creates a UpdateRepoOK with default headers values
func NewUpdateRepoOK() *UpdateRepoOK {
	return &UpdateRepoOK{}
}

/*UpdateRepoOK handles this case with default header values.

Successful invocation
*/
type UpdateRepoOK struct {
}

func (o *UpdateRepoOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}][%d] updateRepoOK ", 200)
}

func (o *UpdateRepoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepoBadRequest creates a UpdateRepoBadRequest with default headers values
func NewUpdateRepoBadRequest() *UpdateRepoBadRequest {
	return &UpdateRepoBadRequest{}
}

/*UpdateRepoBadRequest handles this case with default header values.

Bad Request
*/
type UpdateRepoBadRequest struct {
	Payload *models.APIError
}

func (o *UpdateRepoBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}][%d] updateRepoBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRepoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepoUnauthorized creates a UpdateRepoUnauthorized with default headers values
func NewUpdateRepoUnauthorized() *UpdateRepoUnauthorized {
	return &UpdateRepoUnauthorized{}
}

/*UpdateRepoUnauthorized handles this case with default header values.

Session required
*/
type UpdateRepoUnauthorized struct {
	Payload *models.APIError
}

func (o *UpdateRepoUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}][%d] updateRepoUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRepoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepoForbidden creates a UpdateRepoForbidden with default headers values
func NewUpdateRepoForbidden() *UpdateRepoForbidden {
	return &UpdateRepoForbidden{}
}

/*UpdateRepoForbidden handles this case with default header values.

Unauthorized access
*/
type UpdateRepoForbidden struct {
	Payload *models.APIError
}

func (o *UpdateRepoForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}][%d] updateRepoForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRepoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepoNotFound creates a UpdateRepoNotFound with default headers values
func NewUpdateRepoNotFound() *UpdateRepoNotFound {
	return &UpdateRepoNotFound{}
}

/*UpdateRepoNotFound handles this case with default header values.

Not found
*/
type UpdateRepoNotFound struct {
	Payload *models.APIError
}

func (o *UpdateRepoNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}][%d] updateRepoNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRepoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
