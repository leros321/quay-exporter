// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/leros321/quay-exporter/pkg/models"
)

// DeleteUserRobotReader is a Reader for the DeleteUserRobot structure.
type DeleteUserRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserRobotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteUserRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteUserRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserRobotNoContent creates a DeleteUserRobotNoContent with default headers values
func NewDeleteUserRobotNoContent() *DeleteUserRobotNoContent {
	return &DeleteUserRobotNoContent{}
}

/*DeleteUserRobotNoContent handles this case with default header values.

Deleted
*/
type DeleteUserRobotNoContent struct {
}

func (o *DeleteUserRobotNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/user/robots/{robot_shortname}][%d] deleteUserRobotNoContent ", 204)
}

func (o *DeleteUserRobotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserRobotBadRequest creates a DeleteUserRobotBadRequest with default headers values
func NewDeleteUserRobotBadRequest() *DeleteUserRobotBadRequest {
	return &DeleteUserRobotBadRequest{}
}

/*DeleteUserRobotBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserRobotBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteUserRobotBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/user/robots/{robot_shortname}][%d] deleteUserRobotBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRobotUnauthorized creates a DeleteUserRobotUnauthorized with default headers values
func NewDeleteUserRobotUnauthorized() *DeleteUserRobotUnauthorized {
	return &DeleteUserRobotUnauthorized{}
}

/*DeleteUserRobotUnauthorized handles this case with default header values.

Session required
*/
type DeleteUserRobotUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteUserRobotUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/user/robots/{robot_shortname}][%d] deleteUserRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRobotForbidden creates a DeleteUserRobotForbidden with default headers values
func NewDeleteUserRobotForbidden() *DeleteUserRobotForbidden {
	return &DeleteUserRobotForbidden{}
}

/*DeleteUserRobotForbidden handles this case with default header values.

Unauthorized access
*/
type DeleteUserRobotForbidden struct {
	Payload *models.APIError
}

func (o *DeleteUserRobotForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/user/robots/{robot_shortname}][%d] deleteUserRobotForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRobotNotFound creates a DeleteUserRobotNotFound with default headers values
func NewDeleteUserRobotNotFound() *DeleteUserRobotNotFound {
	return &DeleteUserRobotNotFound{}
}

/*DeleteUserRobotNotFound handles this case with default header values.

Not found
*/
type DeleteUserRobotNotFound struct {
	Payload *models.APIError
}

func (o *DeleteUserRobotNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/user/robots/{robot_shortname}][%d] deleteUserRobotNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
