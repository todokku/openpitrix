// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// ValidateRuntimeCredentialReader is a Reader for the ValidateRuntimeCredential structure.
type ValidateRuntimeCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateRuntimeCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewValidateRuntimeCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateRuntimeCredentialOK creates a ValidateRuntimeCredentialOK with default headers values
func NewValidateRuntimeCredentialOK() *ValidateRuntimeCredentialOK {
	return &ValidateRuntimeCredentialOK{}
}

/*ValidateRuntimeCredentialOK handles this case with default header values.

A successful response.
*/
type ValidateRuntimeCredentialOK struct {
	Payload *models.OpenpitrixValidateRuntimeCredentialResponse
}

func (o *ValidateRuntimeCredentialOK) Error() string {
	return fmt.Sprintf("[POST /v1/runtimes/credentials:validate][%d] validateRuntimeCredentialOK  %+v", 200, o.Payload)
}

func (o *ValidateRuntimeCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixValidateRuntimeCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}