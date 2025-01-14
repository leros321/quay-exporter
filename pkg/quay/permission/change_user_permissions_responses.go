// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/leros321/quay-exporter/pkg/models"
)

// ChangeUserPermissionsReader is a Reader for the ChangeUserPermissions structure.
type ChangeUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChangeUserPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewChangeUserPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewChangeUserPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChangeUserPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeUserPermissionsOK creates a ChangeUserPermissionsOK with default headers values
func NewChangeUserPermissionsOK() *ChangeUserPermissionsOK {
	return &ChangeUserPermissionsOK{}
}

/*ChangeUserPermissionsOK handles this case with default header values.

Successful invocation
*/
type ChangeUserPermissionsOK struct {
}

func (o *ChangeUserPermissionsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsOK ", 200)
}

func (o *ChangeUserPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeUserPermissionsBadRequest creates a ChangeUserPermissionsBadRequest with default headers values
func NewChangeUserPermissionsBadRequest() *ChangeUserPermissionsBadRequest {
	return &ChangeUserPermissionsBadRequest{}
}

/*ChangeUserPermissionsBadRequest handles this case with default header values.

Bad Request
*/
type ChangeUserPermissionsBadRequest struct {
	Payload *models.APIError
}

func (o *ChangeUserPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeUserPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeUserPermissionsUnauthorized creates a ChangeUserPermissionsUnauthorized with default headers values
func NewChangeUserPermissionsUnauthorized() *ChangeUserPermissionsUnauthorized {
	return &ChangeUserPermissionsUnauthorized{}
}

/*ChangeUserPermissionsUnauthorized handles this case with default header values.

Session required
*/
type ChangeUserPermissionsUnauthorized struct {
	Payload *models.APIError
}

func (o *ChangeUserPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ChangeUserPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeUserPermissionsForbidden creates a ChangeUserPermissionsForbidden with default headers values
func NewChangeUserPermissionsForbidden() *ChangeUserPermissionsForbidden {
	return &ChangeUserPermissionsForbidden{}
}

/*ChangeUserPermissionsForbidden handles this case with default header values.

Unauthorized access
*/
type ChangeUserPermissionsForbidden struct {
	Payload *models.APIError
}

func (o *ChangeUserPermissionsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *ChangeUserPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeUserPermissionsNotFound creates a ChangeUserPermissionsNotFound with default headers values
func NewChangeUserPermissionsNotFound() *ChangeUserPermissionsNotFound {
	return &ChangeUserPermissionsNotFound{}
}

/*ChangeUserPermissionsNotFound handles this case with default header values.

Not found
*/
type ChangeUserPermissionsNotFound struct {
	Payload *models.APIError
}

func (o *ChangeUserPermissionsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/permissions/user/{username}][%d] changeUserPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *ChangeUserPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
