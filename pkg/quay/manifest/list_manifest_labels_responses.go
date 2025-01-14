// Code generated by go-swagger; DO NOT EDIT.

package manifest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/leros321/quay-exporter/pkg/models"
)

// ListManifestLabelsReader is a Reader for the ListManifestLabels structure.
type ListManifestLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListManifestLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListManifestLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListManifestLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListManifestLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListManifestLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListManifestLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListManifestLabelsOK creates a ListManifestLabelsOK with default headers values
func NewListManifestLabelsOK() *ListManifestLabelsOK {
	return &ListManifestLabelsOK{}
}

/*ListManifestLabelsOK handles this case with default header values.

Successful invocation
*/
type ListManifestLabelsOK struct {
}

func (o *ListManifestLabelsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/manifest/{manifestref}/labels][%d] listManifestLabelsOK ", 200)
}

func (o *ListManifestLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManifestLabelsBadRequest creates a ListManifestLabelsBadRequest with default headers values
func NewListManifestLabelsBadRequest() *ListManifestLabelsBadRequest {
	return &ListManifestLabelsBadRequest{}
}

/*ListManifestLabelsBadRequest handles this case with default header values.

Bad Request
*/
type ListManifestLabelsBadRequest struct {
	Payload *models.APIError
}

func (o *ListManifestLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/manifest/{manifestref}/labels][%d] listManifestLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *ListManifestLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManifestLabelsUnauthorized creates a ListManifestLabelsUnauthorized with default headers values
func NewListManifestLabelsUnauthorized() *ListManifestLabelsUnauthorized {
	return &ListManifestLabelsUnauthorized{}
}

/*ListManifestLabelsUnauthorized handles this case with default header values.

Session required
*/
type ListManifestLabelsUnauthorized struct {
	Payload *models.APIError
}

func (o *ListManifestLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/manifest/{manifestref}/labels][%d] listManifestLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListManifestLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManifestLabelsForbidden creates a ListManifestLabelsForbidden with default headers values
func NewListManifestLabelsForbidden() *ListManifestLabelsForbidden {
	return &ListManifestLabelsForbidden{}
}

/*ListManifestLabelsForbidden handles this case with default header values.

Unauthorized access
*/
type ListManifestLabelsForbidden struct {
	Payload *models.APIError
}

func (o *ListManifestLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/manifest/{manifestref}/labels][%d] listManifestLabelsForbidden  %+v", 403, o.Payload)
}

func (o *ListManifestLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManifestLabelsNotFound creates a ListManifestLabelsNotFound with default headers values
func NewListManifestLabelsNotFound() *ListManifestLabelsNotFound {
	return &ListManifestLabelsNotFound{}
}

/*ListManifestLabelsNotFound handles this case with default header values.

Not found
*/
type ListManifestLabelsNotFound struct {
	Payload *models.APIError
}

func (o *ListManifestLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/manifest/{manifestref}/labels][%d] listManifestLabelsNotFound  %+v", 404, o.Payload)
}

func (o *ListManifestLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
