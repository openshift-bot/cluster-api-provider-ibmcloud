// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2VolumescloneExecutePostReader is a Reader for the PcloudV2VolumescloneExecutePost structure.
type PcloudV2VolumescloneExecutePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneExecutePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudV2VolumescloneExecutePostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2VolumescloneExecutePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2VolumescloneExecutePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV2VolumescloneExecutePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2VolumescloneExecutePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2VolumescloneExecutePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudV2VolumescloneExecutePostAccepted creates a PcloudV2VolumescloneExecutePostAccepted with default headers values
func NewPcloudV2VolumescloneExecutePostAccepted() *PcloudV2VolumescloneExecutePostAccepted {
	return &PcloudV2VolumescloneExecutePostAccepted{}
}

/* PcloudV2VolumescloneExecutePostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudV2VolumescloneExecutePostAccepted struct {
	Payload *models.VolumesClone
}

func (o *PcloudV2VolumescloneExecutePostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/execute][%d] pcloudV2VolumescloneExecutePostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudV2VolumescloneExecutePostAccepted) GetPayload() *models.VolumesClone {
	return o.Payload
}

func (o *PcloudV2VolumescloneExecutePostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesClone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneExecutePostBadRequest creates a PcloudV2VolumescloneExecutePostBadRequest with default headers values
func NewPcloudV2VolumescloneExecutePostBadRequest() *PcloudV2VolumescloneExecutePostBadRequest {
	return &PcloudV2VolumescloneExecutePostBadRequest{}
}

/* PcloudV2VolumescloneExecutePostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2VolumescloneExecutePostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneExecutePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/execute][%d] pcloudV2VolumescloneExecutePostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudV2VolumescloneExecutePostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneExecutePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneExecutePostUnauthorized creates a PcloudV2VolumescloneExecutePostUnauthorized with default headers values
func NewPcloudV2VolumescloneExecutePostUnauthorized() *PcloudV2VolumescloneExecutePostUnauthorized {
	return &PcloudV2VolumescloneExecutePostUnauthorized{}
}

/* PcloudV2VolumescloneExecutePostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2VolumescloneExecutePostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneExecutePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/execute][%d] pcloudV2VolumescloneExecutePostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudV2VolumescloneExecutePostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneExecutePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneExecutePostForbidden creates a PcloudV2VolumescloneExecutePostForbidden with default headers values
func NewPcloudV2VolumescloneExecutePostForbidden() *PcloudV2VolumescloneExecutePostForbidden {
	return &PcloudV2VolumescloneExecutePostForbidden{}
}

/* PcloudV2VolumescloneExecutePostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV2VolumescloneExecutePostForbidden struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneExecutePostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/execute][%d] pcloudV2VolumescloneExecutePostForbidden  %+v", 403, o.Payload)
}
func (o *PcloudV2VolumescloneExecutePostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneExecutePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneExecutePostNotFound creates a PcloudV2VolumescloneExecutePostNotFound with default headers values
func NewPcloudV2VolumescloneExecutePostNotFound() *PcloudV2VolumescloneExecutePostNotFound {
	return &PcloudV2VolumescloneExecutePostNotFound{}
}

/* PcloudV2VolumescloneExecutePostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudV2VolumescloneExecutePostNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneExecutePostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/execute][%d] pcloudV2VolumescloneExecutePostNotFound  %+v", 404, o.Payload)
}
func (o *PcloudV2VolumescloneExecutePostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneExecutePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneExecutePostInternalServerError creates a PcloudV2VolumescloneExecutePostInternalServerError with default headers values
func NewPcloudV2VolumescloneExecutePostInternalServerError() *PcloudV2VolumescloneExecutePostInternalServerError {
	return &PcloudV2VolumescloneExecutePostInternalServerError{}
}

/* PcloudV2VolumescloneExecutePostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneExecutePostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneExecutePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/execute][%d] pcloudV2VolumescloneExecutePostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudV2VolumescloneExecutePostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneExecutePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}