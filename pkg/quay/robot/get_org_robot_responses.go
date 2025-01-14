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

// GetOrgRobotReader is a Reader for the GetOrgRobot structure.
type GetOrgRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrgRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrgRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetOrgRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrgRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrgRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrgRobotOK creates a GetOrgRobotOK with default headers values
func NewGetOrgRobotOK() *GetOrgRobotOK {
	return &GetOrgRobotOK{}
}

/*GetOrgRobotOK handles this case with default header values.

Successful invocation
*/
type GetOrgRobotOK struct {
}

func (o *GetOrgRobotOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] getOrgRobotOK ", 200)
}

func (o *GetOrgRobotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrgRobotBadRequest creates a GetOrgRobotBadRequest with default headers values
func NewGetOrgRobotBadRequest() *GetOrgRobotBadRequest {
	return &GetOrgRobotBadRequest{}
}

/*GetOrgRobotBadRequest handles this case with default header values.

Bad Request
*/
type GetOrgRobotBadRequest struct {
	Payload *models.APIError
}

func (o *GetOrgRobotBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] getOrgRobotBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotUnauthorized creates a GetOrgRobotUnauthorized with default headers values
func NewGetOrgRobotUnauthorized() *GetOrgRobotUnauthorized {
	return &GetOrgRobotUnauthorized{}
}

/*GetOrgRobotUnauthorized handles this case with default header values.

Session required
*/
type GetOrgRobotUnauthorized struct {
	Payload *models.APIError
}

func (o *GetOrgRobotUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] getOrgRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotForbidden creates a GetOrgRobotForbidden with default headers values
func NewGetOrgRobotForbidden() *GetOrgRobotForbidden {
	return &GetOrgRobotForbidden{}
}

/*GetOrgRobotForbidden handles this case with default header values.

Unauthorized access
*/
type GetOrgRobotForbidden struct {
	Payload *models.APIError
}

func (o *GetOrgRobotForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] getOrgRobotForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotNotFound creates a GetOrgRobotNotFound with default headers values
func NewGetOrgRobotNotFound() *GetOrgRobotNotFound {
	return &GetOrgRobotNotFound{}
}

/*GetOrgRobotNotFound handles this case with default header values.

Not found
*/
type GetOrgRobotNotFound struct {
	Payload *models.APIError
}

func (o *GetOrgRobotNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] getOrgRobotNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
