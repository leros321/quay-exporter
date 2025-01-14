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

// RegenerateUserRobotTokenReader is a Reader for the RegenerateUserRobotToken structure.
type RegenerateUserRobotTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegenerateUserRobotTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewRegenerateUserRobotTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRegenerateUserRobotTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRegenerateUserRobotTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRegenerateUserRobotTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRegenerateUserRobotTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegenerateUserRobotTokenCreated creates a RegenerateUserRobotTokenCreated with default headers values
func NewRegenerateUserRobotTokenCreated() *RegenerateUserRobotTokenCreated {
	return &RegenerateUserRobotTokenCreated{}
}

/*RegenerateUserRobotTokenCreated handles this case with default header values.

Successful creation
*/
type RegenerateUserRobotTokenCreated struct {
}

func (o *RegenerateUserRobotTokenCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/robots/{robot_shortname}/regenerate][%d] regenerateUserRobotTokenCreated ", 201)
}

func (o *RegenerateUserRobotTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateUserRobotTokenBadRequest creates a RegenerateUserRobotTokenBadRequest with default headers values
func NewRegenerateUserRobotTokenBadRequest() *RegenerateUserRobotTokenBadRequest {
	return &RegenerateUserRobotTokenBadRequest{}
}

/*RegenerateUserRobotTokenBadRequest handles this case with default header values.

Bad Request
*/
type RegenerateUserRobotTokenBadRequest struct {
	Payload *models.APIError
}

func (o *RegenerateUserRobotTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/robots/{robot_shortname}/regenerate][%d] regenerateUserRobotTokenBadRequest  %+v", 400, o.Payload)
}

func (o *RegenerateUserRobotTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateUserRobotTokenUnauthorized creates a RegenerateUserRobotTokenUnauthorized with default headers values
func NewRegenerateUserRobotTokenUnauthorized() *RegenerateUserRobotTokenUnauthorized {
	return &RegenerateUserRobotTokenUnauthorized{}
}

/*RegenerateUserRobotTokenUnauthorized handles this case with default header values.

Session required
*/
type RegenerateUserRobotTokenUnauthorized struct {
	Payload *models.APIError
}

func (o *RegenerateUserRobotTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/robots/{robot_shortname}/regenerate][%d] regenerateUserRobotTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *RegenerateUserRobotTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateUserRobotTokenForbidden creates a RegenerateUserRobotTokenForbidden with default headers values
func NewRegenerateUserRobotTokenForbidden() *RegenerateUserRobotTokenForbidden {
	return &RegenerateUserRobotTokenForbidden{}
}

/*RegenerateUserRobotTokenForbidden handles this case with default header values.

Unauthorized access
*/
type RegenerateUserRobotTokenForbidden struct {
	Payload *models.APIError
}

func (o *RegenerateUserRobotTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/robots/{robot_shortname}/regenerate][%d] regenerateUserRobotTokenForbidden  %+v", 403, o.Payload)
}

func (o *RegenerateUserRobotTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateUserRobotTokenNotFound creates a RegenerateUserRobotTokenNotFound with default headers values
func NewRegenerateUserRobotTokenNotFound() *RegenerateUserRobotTokenNotFound {
	return &RegenerateUserRobotTokenNotFound{}
}

/*RegenerateUserRobotTokenNotFound handles this case with default header values.

Not found
*/
type RegenerateUserRobotTokenNotFound struct {
	Payload *models.APIError
}

func (o *RegenerateUserRobotTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/user/robots/{robot_shortname}/regenerate][%d] regenerateUserRobotTokenNotFound  %+v", 404, o.Payload)
}

func (o *RegenerateUserRobotTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
